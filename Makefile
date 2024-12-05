all: assets
	go build -o bin/build_ruby

setup:
	mise install
	go mod tidy

# failing:
test: assets
	go test

assets:
	go-bindata data/...

deps:
	go mod tidy

update_deps:
	go mod -u
