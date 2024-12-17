import re

def partOne(data: str) -> str:
  A, B, C, *program = map(int, re.findall(r"\d+", data))

  pointer = 0
  output  = []

  while pointer < len(program):
    opcode       = program[pointer]
    operand      = program[pointer + 1]
    combooperand = operand

    if   combooperand == 4: combooperand = A
    elif combooperand == 5: combooperand = B
    elif combooperand == 6: combooperand = C

    if   opcode == 0: A = A >> combooperand
    elif opcode == 1: B = B ^ operand
    elif opcode == 2: B = combooperand % 8
    elif opcode == 3: 
      if A != 0: pointer = operand; continue
    elif opcode == 4: B = B ^ C
    elif opcode == 5: output.append(combooperand % 8)
    elif opcode == 6: B = A >> combooperand
    elif opcode == 7: C = A >> combooperand

    pointer += 2

  return ",".join([str(item) for item in output])
