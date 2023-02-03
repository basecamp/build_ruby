package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"github.com/google/uuid"
	"github.com/urfave/cli"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"text/template"
)

var (
	distros = map[string]string{
		"ubuntu:18.04":        "ubuntu:18.04",
		"ubuntu:18.04:libssl": "ubuntu:18.04:libssl", // drop this variant once we move beyond bionic
		"ubuntu:18.04:yjit":   "ubuntu:18.04:yjit",   // used for building 3.2.0 with yjit support
		"ubuntu:18.04:libssl-yjit":   "ubuntu:18.04:libssl-yjit",   // used for building 3.2.0 with yjit support and new libssl
	}

	docker_client   *docker.Client
	docker_endpoint string
	red             func(string) string
	green           func(string) string
	light_green     func(string) string
)

const image_tag string = "ruby_build"

func init() {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}

	docker_client = client
}

func main() {

	app := cli.NewApp()
	app.Name = "build_ruby"
	app.Usage = "Build ruby debs from source for Ubuntu"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "ruby, r",
			Value: "",
			Usage: "Required. The version to build, eg. 2.1.0 (for recent versions with no patch release) or 2.0.0-p451",
		},
		cli.StringFlag{
			Name:  "distro, d",
			Value: "ubuntu:12.04",
			Usage: "Which distro to use for the build",
		},
		cli.StringFlag{
			Name:  "arch, a",
			Value: "amd64",
			Usage: "Arch to use in package filename, eg: 'none', 'all', 'amd64' etc.",
		},
		cli.StringFlag{
			Name:  "iteration, i",
			Value: "",
			Usage: "eg: 37s~precise",
		},
		cli.IntFlag{
			Name:  "cpus, c",
			Usage: "The number of CPUs to use for the make process, defaults to the number in the local machine. Change if you're running on a remote docker host",
		},
	}
	app.Action = buildRuby
	app.Run(os.Args)
}

func buildRuby(c *cli.Context) error {
	if c.String("ruby") == "" {
		fmt.Fprintln(os.Stderr, "@{r!}You didn't specify a Ruby version to build!")
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

	if distros[c.String("distro")] == "" {
		fmt.Fprintln(os.Stderr, "@{r!}You specified a distro that I don't know how to build for")
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

	var parallel_make_tasks int
	if c.Int("cpus") != 0 {
		parallel_make_tasks = c.Int("cpus")
	} else {
		parallel_make_tasks = runtime.NumCPU()
	}

	var patch_file_full_paths []string = patchFilePathsFromRubyVersion(c.String("ruby"))
	var gemfile_path, gemfile_lock_path = gemfilesFromDistro(distros[c.String("distro")])

	var dockerfile *bytes.Buffer = dockerFileFromTemplate(distros[c.String("distro")], c.String("ruby"), c.String("arch"), c.String("iteration"), fileBasePaths(patch_file_full_paths), parallel_make_tasks)
	fmt.Println("@{g!}Using Dockerfile:")
	fmt.Printf("@{gc}%s\n", dockerfile)

	var build_tarfile *bytes.Buffer = createTarFileFromDockerfile(dockerfile, patch_file_full_paths, gemfile_path, gemfile_lock_path)

	image_uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	image_name := fmt.Sprintf("ruby_build_%s_image", image_uuid)
	opts := docker.BuildImageOptions{
		Name:                image_name,
		NoCache:             true,
		SuppressOutput:      false,
		RmTmpContainer:      true,
		ForceRmTmpContainer: true,
		InputStream:         build_tarfile,
		OutputStream:        os.Stdout,
	}
	if err := docker_client.BuildImage(opts); err != nil {
		panic(err)
	}
	fmt.Printf("@{g!}Created image with name %s\n", image_name)

	image, err := docker_client.InspectImage(image_name)
	if err != nil {
		panic(err)
	}

	/*
		Create a container with the created image id

		This seems like a hack. We need a "container" to enable us to copy the ruby
		package out, but I can't see how to do this without needed to run a command
		or just use specify an image ID directly, hence the noop 'date' command.

	*/

	fmt.Printf("@{g!}Creating container with from image id %s\n", image.ID)
	config := docker.Config{AttachStdout: false, AttachStdin: false, Image: image.ID, Cmd: []string{"date"}}
	container_uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	container_name := fmt.Sprintf("ruby_build_%s_container", container_uuid)
	create_container_opts := docker.CreateContainerOptions{Name: container_name, Config: &config}
	container, err := docker_client.CreateContainer(create_container_opts)
	if err != nil {
		panic(err)
	}

	// See https://github.com/wjessop/build_ruby/issues/2
	if err := docker_client.StopContainer(container.ID, 1); err != nil {
		// panic(err)
		fmt.Printf("@{r!}Failed to stop container %d, error was: %s\n", container.ID, err.Error())
	}

	copyPackageFromContainerToLocalFs(container, rubyPackageFileName(c.String("ruby"), c.String("iteration"), c.String("arch"), c.String("distro")))

	fmt.Println("@{g!}Removing container:", container.ID)
	if err := docker_client.RemoveContainer(docker.RemoveContainerOptions{ID: container.ID, RemoveVolumes: true, Force: false}); err != nil {
		panic(err)
	}

	return nil
}

func patchFilePathsFromRubyVersion(version string) []string {
	var patch_files []string
	for _, name := range AssetNames() {
		if strings.Contains(name, fmt.Sprintf("/%s/", version)) {
			patch_files = append(patch_files, name)
		}
	}

	sort.Strings(patch_files)
	fmt.Printf("@{g}Found patch files for current Ruby version: %v\n", patch_files)
	return patch_files
}

func createTarFileFromDockerfile(dockerfile *bytes.Buffer, patch_file_paths []string, gemfile_path string, gemfile_lock_path string) *bytes.Buffer {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new tar archive.
	tw := tar.NewWriter(buf)

	// Add the Dockerfile
	hdr := &tar.Header{
		Name: "Dockerfile",
		Size: int64(dockerfile.Len()),
	}

	if err := tw.WriteHeader(hdr); err != nil {
		panic(err)
	}

	if _, err := tw.Write(dockerfile.Bytes()); err != nil {
		panic(err)
	}

	for _, path := range patch_file_paths {
		fmt.Printf("@{g}Adding patch file to the tar: %s (at path %s)\n", patch_file_paths, filepath.Base(path))

		asset_bytes, err := Asset(path)
		if err != nil {
			panic(err)
		}

		// We store the patch files flat in the root dir, hence the Base call
		hdr := &tar.Header{
			Name: filepath.Base(path),
			Size: int64(len(asset_bytes)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			panic(err)
		}

		if _, err := tw.Write(asset_bytes); err != nil {
			panic(err)
		}
	}

	fmt.Printf("@{g}Adding %s to the tar as Gemfile\n", gemfile_path)

	asset_bytes, err := Asset(gemfile_path)
	if err != nil {
		panic(err)
	}

	hdr = &tar.Header{
		Name: "Gemfile",
		Size: int64(len(asset_bytes)),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		panic(err)
	}

	if _, err := tw.Write(asset_bytes); err != nil {
		panic(err)
	}

	fmt.Printf("@{g}Adding %s to the tar as Gemfile.lock\n", gemfile_lock_path)

	asset_bytes, err = Asset(gemfile_path)
	if err != nil {
		panic(err)
	}

	hdr = &tar.Header{
		Name: "Gemfile.lock",
		Size: int64(len(asset_bytes)),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		panic(err)
	}

	if _, err := tw.Write(asset_bytes); err != nil {
		panic(err)
	}

	// Make sure to check the error on Close.
	if err := tw.Close(); err != nil {
		panic(err)
	}

	return buf
}

func copyPackageFromContainerToLocalFs(container *docker.Container, filename string) {
	fmt.Println("@{g!}Copying package out of the container")

	var buf bytes.Buffer
	if err := docker_client.DownloadFromContainer(container.ID, docker.DownloadFromContainerOptions{
		Path:         filename,
		OutputStream: &buf,
	}); err != nil {
		panic(err)
	}

	buffer := bytes.NewReader(buf.Bytes())

	var tar_out *tar.Reader = tar.NewReader(buffer)
	tar_header, err := tar_out.Next()
	if err != nil {
		panic(err)
	}

	fmt.Printf("@{g!}Extracting package file %s (%d bytes)\n", tar_header.Name, tar_header.Size)

	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	io.Copy(out, tar_out)
}

func rubyPackageFileName(version, iteration, arch string, distro string) string {
	var formatted_iteration = ""
	if iteration != "" {
		formatted_iteration = "_" + iteration
	}

	var formatted_arch = ""
	if arch != "none" {
		formatted_arch = "_" + arch
	}
	return "ruby-" + version + formatted_iteration + formatted_arch + packageFormat(distro)
}

func packageFormat(distro string) string {
	if strings.Contains(distro, "centos") || strings.Contains(distro, "rhel") {
		return ".rpm"
	} else {
		return ".deb"
	}
}

func dockerFileFromTemplate(distro, ruby_version, arch, iteration string, patches []string, parallel_make_jobs int) *bytes.Buffer {
	type buildVars struct {
		Distro      string
		RubyVersion string
		Arch        string
		Iteration   string
		DownloadUrl string
		FileName    string
		Patches     []string
		NumCPU      int
	}

	var formatted_iteration = ""
	if iteration != "" {
		formatted_iteration = fmt.Sprintf("--iteration %s \\", iteration)
	}

	download_url := rubyDownloadUrl(ruby_version)
	dockerfile_vars := buildVars{distro, ruby_version, arch, formatted_iteration, download_url, rubyPackageFileName(ruby_version, iteration, arch, distro), patches, parallel_make_jobs}

	// This would be way better as a look up table, or with a more formal lookup process
	var template_location string
	switch distro {
	case "ubuntu:18.04":
		template_location = "data/Dockerfile-bionic.template"
	case "ubuntu:18.04:libssl": // drop this variant once we move beyond bionic
		template_location = "data/Dockerfile-bionic-libssl.template"
	case "ubuntu:18.04:yjit": // used for building 3.2.0 with yjit support
		template_location = "data/Dockerfile-bionic-yjit.template"
	case "ubuntu:18.04:libssl-yjit": // used for building 3.2.0 with yjit support and new libssl
		template_location = "data/Dockerfile-bionic-libssl-yjit.template"
	default:
		template_location = "data/Dockerfile.template"
	}

	dockerfile_template, err := Asset(template_location)
	if err != nil {
		panic(err)
	}
	if len(dockerfile_template) == 0 {
		panic("Couldn't find Dockerfile template in bindata")
	}

	tmpl, err := template.New("dockerfile_template").Parse(string(dockerfile_template))
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, dockerfile_vars)
	if err != nil {
		panic(err)
	}

	return buf
}

func gemfilesFromDistro(distro string) (string, string) {
	switch distro {
	case "ubuntu:18.04":
		return "data/Gemfile.bionic", "data/Gemfile.bionic.lock"
	case "ubuntu:18.04:libssl": // drop this variant once we move beyond bionic
		return "data/Gemfile.bionic", "data/Gemfile.bionic.lock"
	case "ubuntu:18.04:yjit": // used for building 3.2.0 with yjit support
		return "data/Gemfile.bionic", "data/Gemfile.bionic.lock"
	case "ubuntu:18.04:libssl-yjit": // used for building 3.2.0 with yjit support and new libssl
		return "data/Gemfile.bionic", "data/Gemfile.bionic.lock"
	default:
		return "data/Gemfile.template", "data/Gemfile.template.lock"
	}
}

func rubyDownloadUrl(version string) string {
	// eg:
	// http://cache.ruby-lang.org/pub/ruby/2.1/ruby-2.1.1.tar.gz
	// http://cache.ruby-lang.org/pub/ruby/2.0/ruby-2.0.0-p451.tar.gz

	v := majorMinorVersionOnly(version)
	return "http://cache.ruby-lang.org/pub/ruby/" + v + "/ruby-" + version + ".tar.gz"
}

func majorMinorVersionOnly(full_version string) string {
	return strings.Join(strings.SplitN(full_version, ".", 3)[0:2], ".")
}

func fileBasePaths(full_paths []string) (base_paths []string) {
	for _, p := range full_paths {
		base_paths = append(base_paths, filepath.Base(p))
	}

	return
}
