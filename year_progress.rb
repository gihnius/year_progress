#!/usr/bin/env ruby
# coding: utf-8

passed_days = Time.now.yday
total_days = Time.local(Time.now.year, 12, 31, 23, 59, 59).yday
progress = ((passed_days.to_f / total_days) * 100).round
bar = '▓' * progress + '░' * (100 - progress)
puts "#{bar} #{progress}%(#{passed_days}/#{total_days})"
