#!/usr/bin/env ruby

require_relative File.join("..", "gen", "ruby", "user", "v1", "user.pb")

def main 
  message = User::V1::User.decode(File.read("user.bin"))

  puts message.inspect
end

main
