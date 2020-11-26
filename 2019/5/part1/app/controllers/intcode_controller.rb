class IntcodeController < ApplicationController
  def run

    ## get input
    input = params[:input] # || ''
    program = input.split(",").map(&:to_i)

    @output = []
    programCounter = 0
    while programCounter < program.length()
    
      # parse instruction
      instruction = program[programCounter].to_s
      opcode_right = instruction[-1]
      opcode_left = instruction[-2] != nil ? instruction[-2] : "0"
      opcode = opcode_left + opcode_right
      mode_first_param = instruction[-3] != nil ? instruction[-3] : "0"
      mode_second_param = instruction[-4] != nil ? instruction[-4] : "0"

      case opcode
      when "01"
        param1 = mode_first_param == "1" ? program[programCounter+1] : program[program[programCounter+1]]
        param2 = mode_second_param == "1" ? program[programCounter+2] : program[program[programCounter+2]]
        resultAddress = program[programCounter+3]
        program[resultAddress] = param1 + param2
        programCounter += 4
      
      when "02"
        param1 = mode_first_param == "1" ? program[programCounter+1] : program[program[programCounter+1]]
        param2 = mode_second_param == "1" ? program[programCounter+2] : program[program[programCounter+2]]
        resultAddress = program[programCounter+3]
        program[resultAddress] = param1 * param2
        programCounter += 4
      
      when "03"
        address = program[programCounter+1]
        program[address] = 1
        programCounter += 2

      when "04"
        param = mode_first_param == "1" ? program[programCounter+1] : program[program[programCounter+1]]
        puts "Output: #{param}"
        @output.append(param)
        programCounter += 2

      when "99"
        puts "Halting"
        break
      
      else
        puts "Unknown opcode '#{opcode}'"
        break
      end
    end

  end
end
