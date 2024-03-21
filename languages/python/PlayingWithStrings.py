firstName = "Lamech "
lastName = 'Dete'
fullName = firstName + lastName
brand = "Apple"
price = 2000.97
rate = 1099
tax = 30
vat = '16%'

message = "%s has bought %s laptop at %4.2f USD at the rate of %d\n"%(fullName,brand, price, rate)

receipt = 'Full Name: {0:s}\nProduct:{1:s}\nPrice:{2:4.2f}\nrate:{3:d}'.format(fullName,brand, price, rate)

code = 'The code standard vat {0}, tax {1}'.format(tax, vat)

codeReverse = 'The reversed valuation is {1} to {0} annually'.format(tax, vat)

stable = (price * 16)/100

vatPerProduct = 'The vat per product is {:4.2f} vat is percentage is {:d}'.format(stable, 16)

total = price + stable

totalPrice = 'Total price:{}'.format(total)

print(message)
print(receipt)
print(code)
print(codeReverse)
print(vatPerProduct)
print(totalPrice)





