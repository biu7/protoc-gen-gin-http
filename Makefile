.PHONY: build generate clean install

build:
	go build -o bin/protoc-gen-gin-http .

install:
	go install .

generate: build
	protoc \
		--proto_path=proto \
		--proto_path=. \
		--go_out=. --go_opt=paths=source_relative \
		--gin-http_out=. --gin-http_opt=paths=source_relative \
		--plugin=protoc-gen-gin-http=./bin/protoc-gen-gin-http \
		example/api/greeter/greeter.proto

clean:
	rm -rf bin/
