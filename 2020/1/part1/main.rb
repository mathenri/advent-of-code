nums = []
File.open('input.txt').each do |line| 
	nums.append(line.to_i)
end

nums.each do |num1|
	nums.each do |num2| 
		if num1 + num2 == 2020
			puts "#{num1} + #{num2} equals 2020! Product equals #{num1 * num2}"
		end
	end 
end