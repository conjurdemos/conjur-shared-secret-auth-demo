require 'sinatra'
require "sinatra/reloader"
require 'rack/token_auth'
require 'securerandom'
require 'conjur-api'

secret = SecureRandom.hex(20)

conjur_username = ENV['CONJUR_AUTHN_LOGIN']
conjur_api_key = ENV['CONJUR_AUTHN_API_KEY']
conjur_account = ENV['CONJUR_ACCOUNT']
conjur_cert = ENV['CONJUR_CERT_FILE']

OpenSSL::SSL::SSLContext::DEFAULT_CERT_STORE.add_file conjur_cert if conjur_cert

api = Conjur::API.new_from_key conjur_username, conjur_api_key
api.resource("#{conjur_account}:variable:helloworld-secret").add_value secret

use Rack::TokenAuth do |token, options, env|
  token == secret
end

get '/' do
  'Hello world!'
end
