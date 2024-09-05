/*
Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.
*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

#define LINECOUNT 1000
#define LINELIMIT 100

int returnNumbers(char arr[])
{

    int firstNum;
    int lastNum;
    int totalNum;

    char *numbers[] = {"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"};
    char *firstWord = NULL;
    char *lastWord = NULL;
    int lastWordPos = -1;
    int firstWordPos = INT_MAX;

    // find the first and last occurance of a word with its index to compare to first occurance of number
    for (int i = 0; i < 9; i++)
    {
        char *position = strstr(arr, numbers[i]);
        // while loop to check whole string other wise it will settle for the first occurance of a string sevensixjczjhjzbj8fnsnrsevenfive2seven would ignore last seven at just grab first one if we dont do this
        while (position != NULL)
        {
            int pos = position - arr;

            if (strstr(arr, numbers[i]) && pos > lastWordPos)
            {
                lastWordPos = pos;
                if (lastWord != NULL)
                {
                    free(lastWord);
                }
                lastWord = (char *)malloc(strlen(numbers[i]) + 1);
                strcpy(lastWord, numbers[i]);
            }
            if (strstr(arr, numbers[i]) && pos < firstWordPos)
            {
                firstWordPos = pos;
                if (firstWord != NULL)
                {
                    free(firstWord);
                }
                firstWord = (char *)malloc(strlen(numbers[i]) + 1);
                strcpy(firstWord, numbers[i]);
            }

            position = strstr(position + 1, numbers[i]);
        }
    }

    printf("First Word: %s pos: %d\tfinal Word: %s pos: %d\n", firstWord, firstWordPos, lastWord, lastWordPos);

    // get the first number of the given string and see whether it is before or after the first word
    for (int x = 0; x < strlen(arr); x++)
    {
        int charToInt = arr[x] - '0';
        if (charToInt >= 0 && charToInt <= 9)
        {
            if (x < firstWordPos)
            {
                firstNum = charToInt;
                break;
            }
            else
            {
                for (int z = 0; z < 9; z++)
                {
                    if (strcmp(firstWord, numbers[z]) == 0)
                    {
                        firstNum = z + 1;
                        break;
                    }
                }
            }
        }
    }

    // Get the last number from the given string and compare to index of final word in string
    for (int y = strlen(arr) - 1; y >= 0; y--)
    {
        int charToInt = arr[y] - '0';
        if (charToInt >= 0 && charToInt <= 9)
        {
            if (y > lastWordPos)
            {
                lastNum = charToInt;
                break;
            }
            else
            {
                for (int z = 0; z < 9; z++)
                {
                    if (strcmp(lastWord, numbers[z]) == 0)
                    {
                        lastNum = z + 1;
                        break;
                    }
                }
            }
        }
    }

    totalNum = (firstNum * 10) + lastNum;
    printf("%d\n", totalNum);

    free(lastWord);
    free(firstWord);

    return totalNum;
}

int main()
{
    FILE *fptr;

    char fileContents[LINECOUNT][LINELIMIT];

    // puts the contents of each line into its own member of an arry
    fptr = fopen("input.txt", "r");
    printf("File had been opened\n");

    for (int i = 0; i < LINECOUNT; i++)
    {
        if (fgets(fileContents[i], LINELIMIT, fptr) == NULL)
        {
            break;
        }
    }

    int total = 0;

    for (int x = 0; x < LINECOUNT; x++)
    {
        total = total + returnNumbers(fileContents[x]);
    }

    printf("The final total is %d", total);

    fclose(fptr);

    return 0;
}