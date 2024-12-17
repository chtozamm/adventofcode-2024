from part1 import partOne
from colorama import Fore, Style

test_case = """
Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
"""
expected = "4,6,3,5,6,3,5,2,1,0"

def run_tests():
  got = partOne(test_case.strip())
  if got != expected: fail(got, expected)
  else: success()

def success(): print(Fore.GREEN + "All tests passed" + Style.RESET_ALL)
def fail(got, expected): print(Fore.RED + "Test case failed\n" + f"  expected:\n    {expected}\n  got:\n    {got}" + Style.RESET_ALL)

if __name__ == "__main__": run_tests()
