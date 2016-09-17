#!/usr/bin/env ruby

def hazumi?
  b = @a.sort
  @a != b && @a != b.reverse
end

x = 0
i = 0

loop do
  i += 1
  @a = i.to_s.split('')
  next unless hazumi?
  x += 1
  if (x * 1.0 / i) == 0.99
    puts i
    break
  end
end

=begin
$ ruby -v
ruby 2.3.0p0 (2015-12-25 revision 53290) [x86_64-darwin14]
$ time ruby hazumi.rb
1587000

real    0m21.346s
user    0m20.645s
sys     0m0.549s
=end
