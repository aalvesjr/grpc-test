# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: post.proto for package 'post'

require 'grpc'
require_relative 'post_pb'

module Post
  module Posts
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'post.Posts'

      rpc :Get, SearchParams, Post
    end

    Stub = Service.rpc_stub_class
  end
end
