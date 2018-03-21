#!/usr/bin/env ruby
# coding: utf-8

past_days = Time.now.yday
total_days = Time.local(Time.now.year, 12, 31, 23, 59, 59).yday
progress = ((past_days.to_f / total_days) * 100).round
bar = '▓' * progress + '░' * (100 - progress)
puts "#{bar} #{progress}%(#{past_days}/#{total_days})"
