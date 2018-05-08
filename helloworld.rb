require 'sinatra'

use Rack::Auth::Basic, "Protected Area" do |username, password|
    username == 'foo' && password == 'bar'
end

get '/' do
  'Hello world!'
end