def recursive_function(n):
    if (n != 0):
        return n * recursive_function(n-1)
    else:
        return 1
    
if __name__ == "__main__":
    number = 4
    result = recursive_function(number)
    print(f'{number} & factorial = {result}')
