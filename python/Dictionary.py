nameAge = {'Lamech':27, 'Faith':23, 'Keith':22, 'Tamaia':17}

print('The initial data is:', nameAge)

highestAge = nameAge['Lamech']

print('Highest age is:', highestAge)

nameAge['Rochelle']= 7

print('The updated list after addition is:',nameAge)

nameAge['Lamech'] = 28

print('The updated list after modification is:',nameAge)

del nameAge['Faith']

print('The updated list after deletion is:',nameAge)