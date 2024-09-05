/*The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover.
On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define LINECOUNT 1000
#define LINELIMIT 100

int returnNumbers(char arr[])
{

	int firstNum;
	int lastNum;
	int totalNum;

	// get the first number of the given string
	for (int x = 0; x < strlen(arr); x++)
	{
		int charToInt = arr[x] - '0';
		if (charToInt >= 0 && charToInt <= 9)
		{
			firstNum = charToInt;
			break;
		}
	}

	// Get the last number from the given string by starting at the end and going backwards
	for (int y = strlen(arr) - 1; y >= 0; y--)
	{
		int charToInt = arr[y] - '0';
		if (charToInt >= 0 && charToInt <= 9)
		{
			lastNum = charToInt;
			break;
		}
	}

	totalNum = (firstNum * 10) + lastNum;

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
