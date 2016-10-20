#!/usr/bin/env ruby

if ARGV.empty?
  puts "ERROR: Please pass an new tag to tag containers"
  exit
end

subtag = ARGV[0]

tagarr = `docker ps -a | grep pics | awk '{print $2'}`.split("\n")
tagarr = tags.split("\n")

tagarr.each do |tag|
  stag = tag.gsub(/$/, ":")
  newtag = stag.gsub(/:.*/, ":#{subtag}")
  puts dockertag = "docker tag #{tag} #{newtag}"
  puts dockerpush = "docker push #{newtag}"
  puts `#{dockertag}`
  puts `#{dockerpush} &`
end
