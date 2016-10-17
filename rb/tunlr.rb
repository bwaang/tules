#!/usr/bin/env ruby
require 'json'

# Read in app ports
rbpath = ENV['RBPATH']
confpath = "#{rbpath}/../etc/ports.json"

# Check for ConfigFile
hasConfig = File.exist?(confpath)

puts confpath if hasConfig
f = File.read(confpath) if hasConfig
ports = JSON.parse(f) if hasConfig

host = ARGV.shift
hostport = ARGV.shift
# If it's not literal port, look for it in config
if hostport =~ /[A-Za-z]/
  hostport = ports[hostport] if hasConfig
end
portprefix = ARGV.shift

# Fetch ip from config
hostip = `cat ~/.ssh/config | grep #{host}$ -A 3 | grep Hostname | awk '{print $2}'`.strip

puts tunnel = "ssh -N -L #{portprefix}#{hostport}:#{hostip}:#{hostport} #{host}"
puts `#{tunnel}`
