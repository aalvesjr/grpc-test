build:
	protoc -I . post.proto --go_out=plugins=grpc:.

build-ruby:
	grpc_tools_ruby_protoc -I . post.proto --ruby_out=. --grpc_out=.
