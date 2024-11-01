import re

# D: Use the auxiliar compare function to get the result and also use regex to parse the text
#    Into lists
# I: Two strings with version format "x.y.z..."
# O: If version 1 > version 2 return 1, if less -1, otherwise 0
def compare(ver1, ver2):
    return compare_aux(re.findall(r'\d+', ver1), re.findall(r'\d+', ver2))

# D: Given two arrays of versions, the first element of the array is the base version
#    For example: 1.2 => [1, 2], then it compares the first element to check which is greater
# I: A list of two versions
# O: If version 1 > version 2 return 1, if less -1, otherwise 0
def compare_aux(ver1, ver2):
    # Base case, one of the list is empty
    if len(ver1) == 0 or len(ver2) == 0:
        return int(len(ver1) != 0) - int(len(ver2) != 0)

    if int(ver1[0]) > int(ver2[0]):
        return 1
    
    if int(ver1[0]) < int(ver2[0]):
        return -1
 
    return compare_aux(ver1[1:], ver2[1:])