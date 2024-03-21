import sys

def maximumMinimumElementsTest(matrix, result):

    s = set()

    for i in range(0, len(matrix[0]), 1):

        maximumRows = -sys.maxsize - 1

        for j in range(0, len(matrix), 1):

            maximumRows = max(maximumRows, matrix[i][j])

        s.add(maximumRows)
 
    for j in range(0, len(matrix), 1):
        
        minimumColumns = sys.maxsize
 
        for i in range(0, len(matrix[j]), 1):

            minimumColumns = min(minimumColumns, matrix[i][j])

        if (maximumRows in s):
            result.append(maximumRows)
 
    return result

if __name__ == '__main__':
     
    mat = [ [ 1, 10, 4 ],
            [ 9, 3, 8 ],
            [ 15, 16, 17 ] ]
 
    ans = []
 
    maximumMinimumElementsTest(mat, ans)

    if (len(ans) == 0):
        print("-1")
 
    for i in range(len(ans)):
        print(ans[i])
