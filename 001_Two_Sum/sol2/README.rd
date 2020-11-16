
Input:
    Array - [6,5,2,7]
    Target - 9

Output:
    Array - [2, 3]


Idea:
    We will create dictionary and store number's indices as values and numbers themselves we will use as keys

    To create our dictionary we will loop once through given array and use index and value to fill dictionary

    dict[6] = 0
    dict[5] = 1
    dict[2] = 2
    dict[7] = 3

    Then we will loop through given array and for each number we will check if we have number which adds up to target.
    this key will be target minus current number.

    key = target - num

    If we find key in dictionary we will get index of current number from loop and index of number from dictionary.
    And finally we will return those indices.

