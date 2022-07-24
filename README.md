# Introduction to Protocol Buffers

## Prerequisites

* **required** Go **1.18**, _and_
* **recommended** `direnv`, to allow all Go-based binaries to be local to this folder and not installed globally. For more details you can refer to [this post](https://mariocarrion.com/2020/11/20/golang-go-tool-direnv.html).

## Tools

Install the following tools:

* **required** Protocol Buffers Compiler, `protoc` (to date version `3.19.4`):
    * Homebrew: `brew install protobuf`
    * Alpine 3.15: `apk add protobuf-dev protobuf`
    * Ubuntu 21.10: `apt-get install protobuf-compiler libprotobuf-dev`
* **required** `buf` CLI for linting and compiling:
    * `go install github.com/bufbuild/buf/cmd/buf@v1.0.0-rc12`
* **required** Protocol Buffer Plugin for Go:
    * `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1`
* **required** gRPC Plugin for Go:
    * `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0`
* **recommended** gRPC Health Check tester using grpc-health-probe:
*   * `go install github.com/grpc-ecosystem/grpc-health-probe@v0.4.8`
* **recommended** Code Formatting, `clang-format`, you can use `find . -name '*.proto' | xargs clang-format -i`
    * Homebrew: `brew install clang-format` (to date version `13.0.1`):
    * Alpine 3.15: `apk add clang-extra-tools`
    * Ubuntu 21.10: `apt-get install clang-format`

## Interoperability example
- Writer: `:! go run writer/main.go`
- Reader: `:! ruby reader/main.rb` --> output: `#<User::V1::User uuid="1-2-3-4" full_name="Anh" birth_year=2999>`
