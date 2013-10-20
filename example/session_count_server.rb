require 'sinatra'
require 'msgpack'
require 'json'
require_relative 'msgpack_cookie'

configure do
  config = JSON.load(File.open(File.join(__dir__, 'config.json')))
  use Rack::Session::Cookie, :key => config['cookie_name'],
                             :coder => Rack::Session::MessagePack.new,
                             :secret => config['cookie_secret']
end

class Visits
  @@global = 0
  def self.global; @@global; end
  def self.add
    @@global +=1
  end
end

get '/' do
  session[:visits] ||= 0
  session[:visits] += 1

  Visits.add #una visita global (server)

  "#{Visits.global} and #{session[:visits]}"
end

