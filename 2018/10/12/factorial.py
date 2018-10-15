def factorial(val):
    if val > 1:
        return factorial(val-1)*(val)
    else:
        return 1

print(factorial(3))
print(factorial(5))