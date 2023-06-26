.PHONY: all docs
all: speakeasy docs

speakeasy:
	speakeasy generate sdk --lang terraform -o . -s openapi.yaml

docs:
	go generate ./...

