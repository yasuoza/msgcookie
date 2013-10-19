require 'sinatra'
require 'msgpack'
require_relative 'msgpack_cookie'

configure do
  use Rack::Session::Cookie, :key => 'rack.session',
                             :coder => Rack::Session::MessagePack.new,
                             :secret => 'special_secret'
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

