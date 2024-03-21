dataSet = [10,20,30,40,50]

dataSet1 = [dataSet[0],dataSet[1],dataSet[2],dataSet[-2],dataSet[-1]]

dataSet2 = dataSet1[1:4]

dataSet3 = dataSet1[1:5:2]

dataSet4 = dataSet1.append(100)

dataSet5 , dataSet6 = dataSet1, dataSet1

dataSet5[0] = 70

del dataSet6[-1]

print('Initial list values:',dataSet)
print('The result of the index 0,1,2,-2,-1:', dataSet1)
print('The result of slice 1:4 is:', dataSet2)
print('The result of the slice stepper 1:5:2 is:', dataSet3)
print('The result of adding items to list with append function is:', dataSet4)
print('The result of modifying the list after adding item is:',dataSet5)
print('The result of deletion after modification of the list is:', dataSet6)
