require 'grpc'
require_relative '../pb-post/post_services_pb'

# GreeterServer is simple server that implements the Helloworld Greeter server.
# class GreeterServer < Helloworld::Greeter::Service

class Server < Post::Posts::Service
  def get(req, _unused_call)
    puts "request id: #{req.id}"
    @posts = {
      1 => Post::Post.new(id: 1, title: "a little title 1", description: "desc 1"),
      2 => Post::Post.new(id: 2, title: "a little title 2", description: "desc 2")
    }
    @posts.fetch(req.id)
  end
end

# main starts an RpcServer that receives requests to GreeterServer at the sample
# server port.
def main
  s = GRPC::RpcServer.new
  s.add_http2_port('0.0.0.0:8080', :this_port_is_insecure)
  s.handle(Server)
  s.run_till_terminated
end

main
