nums = []
File.open('input.txt').each do |line|
	nums.append(line.to_i)
end

nums.each do |num1|
	nums.each do |num2|
		nums.each do |num3|
			if num1 + num2 + num3 == 2020
				puts "#{num1} + #{num2} + #{num3} == 2020! Product: #{num1 * num2 * num3}"
			end
		end
	end
end