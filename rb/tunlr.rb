#!/usr/bin/env ruby
require 'json'

# Read in app ports
tules_home = ENV['TULES_HOME']
conf_path = "#{tules_home}/etc/ports.json"

# Check for ConfigFile existence
has_config = File.exist?(conf_path)

puts conf_path if has_config
f = File.read(conf_path) if has_config
ports = JSON.parse(f) if has_config

host = ARGV.shift
# if host exists in config use that one
host = ports['envs'][host] if ports['envs'].has_key? host and has_config

host_port = ARGV.shift
# If it's not literal port, look for it in config
host_port = ports[host_port] if host_port =~ /[A-Za-z]/ and has_config

port_prefix = ARGV.shift

# Fetch ip from config
host_ip = `cat ~/.ssh/config | grep #{host}$ -A 3 | grep Hostname | awk '{print $2}'`.strip

puts tunnel = "ssh -N -L #{port_prefix}#{host_port}:#{host_ip}:#{host_port} #{host}"
puts `#{tunnel}`
