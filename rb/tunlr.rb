#!/usr/bin/env ruby
require 'json'

# Read in app ports
rbpath = ENV['RBPATH']
f = File.read("#{rbpath}/../etc/ports.json")
appports = JSON.parse(f)

host = ARGV.shift
hostport = ARGV.shift
# If it's not literal port, look for it in config
if hostport =~ /[A-Za-z]/
  hostport = appports[hostport]
end
portprefix = ARGV.shift

# Fetch ip from config
hostip = `cat ~/.ssh/config | grep #{host}$ -A 3 | grep Hostname | awk '{print $2}'`.strip

tunnel = "ssh -N -L #{portprefix}#{hostport}:#{hostip}:#{hostport} #{host}"
puts tunnel
puts `#{tunnel}`
