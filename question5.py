# Question 5: Reverse Integer
# Write a program that takes an integer as input and returns an integer with reversed digit
# ordering.

def reverse_no(number):
    sign = 1 if number >= 0 else -1
    reverse_no = int(str(abs(number))[::-1]) * sign
    return reverse_no


input_no = int(input("Enter an integer: "))
result = reverse_no(input_no)
print(result)
