monthsOfTheYear = ('Jan','Feb','Mar','Apr','May','June','July','Aug','Sep','Oct','Nov','Dec')

firtsMonth = monthsOfTheYear[0]

lastMonth = monthsOfTheYear[-1]

slicedMonth = monthsOfTheYear[1:6]

slicedStepperMonth = monthsOfTheYear[2:8:2]

print('Months of the year:',monthsOfTheYear)
print('Month in the first index:',firtsMonth)
print('Month in the last index:',lastMonth)
print('Months between index 0 and 6:',slicedMonth)
print('Months between index 1 and 8:',slicedStepperMonth)