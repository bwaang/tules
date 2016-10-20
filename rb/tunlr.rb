#!/usr/bin/env ruby
require 'json'

# Read in app ports
tules_home = ENV['TULES_HOME']
conf_path = "#{tules_home}/etc/ports.json"

# Check for ConfigFile
has_config = File.exist?(conf_path)

puts conf_path if has_config
f = File.read(conf_path) if has_config
ports = JSON.parse(f) if has_config

host = ARGV.shift
host_port = ARGV.shift
# If it's not literal port, look for it in config
if host_port =~ /[A-Za-z]/
  host_port = ports[host_port] if has_config
end
port_prefix = ARGV.shift

# Fetch ip from config
host_ip = `cat ~/.ssh/config | grep #{host}$ -A 3 | grep Hostname | awk '{print $2}'`.strip

puts tunnel = "ssh -N -L #{port_prefix}#{host_port}:#{host_ip}:#{host_port} #{host}"
puts `#{tunnel}`
