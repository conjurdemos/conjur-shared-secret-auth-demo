require 'sinatra'
require 'rack/token_auth'
require 'securerandom'
require 'conjur-api'

# Get service configuration inputs from environment
conjur_account = ENV['CONJUR_ACCOUNT']
conjur_username = ENV['CONJUR_AUTHN_LOGIN']
conjur_api_key = ENV['CONJUR_AUTHN_API_KEY']
conjur_cert = ENV['CONJUR_CERT_FILE']
secret_id = ENV['DEMO_SERVICE_KEY_ID']
resource_id = "#{conjur_account}:variable:#{secret_id}"

OpenSSL::SSL::SSLContext::DEFAULT_CERT_STORE.add_file conjur_cert if conjur_cert

# Generate service key value and store it in the attached
# Conjur service
secret = SecureRandom.hex(20)
api = Conjur::API.new_from_key conjur_username, conjur_api_key
api.resource(resource_id).add_value secret

# Set up token authentication for our service with
# the generated service key
use Rack::TokenAuth do |token, _options, _env|
  token == secret
end

# Our service simply responds to an authenticated GET
# request on the root path with the string, "Hello world!"
get '/' do
  'Hello world!'
end
