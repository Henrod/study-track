RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)
PROJECTNAME := $(shell basename "$(PWD)")

.PHONY: help
help:
	@echo "Choose a command run in "$(PROJECTNAME)": \n"
	@find . -maxdepth 1 -type f \( -name Makefile -or -name "*.mk" \) -exec cat {} \+ | sed -n 's/^##//p' | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: deps/up
## deps/up: starts the project dependencies
deps/up:

.PHONY: generate/protos
## generate/protos: create go files from proto files
generate/protos:
	protoc \
      -I pkg/studytrack \
      -I third_party/googleapis \
      --grpc-gateway_out=logtostderr=true:pkg/studytrack \
      --go_out=plugins=grpc:pkg/studytrack \
      pkg/studytrack/*.proto
