require 'grpc'
require_relative '../pb-post/post_services_pb'

def main
  stub = Post::Posts::Stub.new('localhost:8080', :this_channel_is_insecure)

  begin
  id = ARGV[0].to_i
  post = stub.get(Post::SearchParams.new(id: id))

  p "PostID: #{post.id}"
  p "Post Title: #{post.title}"
  p "Post Description: #{post.description}"
  rescue
    p "PostID #{id} not found"
  end
end

main
