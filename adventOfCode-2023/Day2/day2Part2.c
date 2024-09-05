/*
As you continue your walk, the Elf poses a second question:
in each game you played, what is the fewest number of cubes of each color that could have been in the bag to make the game possible?

Again consider the example games from earlier:

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

In game 1, the game could have been played with as few as 4 red, 2 green, and 6 blue cubes. If any color had even one fewer cube, the game would have been impossible.
Game 2 could have been played with a minimum of 1 red, 3 green, and 4 blue cubes.
Game 3 must have been played with at least 20 red, 13 green, and 6 blue cubes.
Game 4 required at least 14 red, 3 green, and 15 blue cubes.
Game 5 needed no fewer than 6 red, 3 green, and 2 blue cubes in the bag.

The power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together.
The power of the minimum set of cubes in game 1 is 48. In games 2-5 it was 12, 1560, 630, and 36, respectively. Adding up these five powers produces the sum 2286.

For each game, find the minimum set of cubes that must have been present. What is the sum of the power of these sets?
*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define LINECOUNT 100
#define LINELIMIT 200

int main()
{
    FILE *fptr;
    char fileContents[LINECOUNT][LINELIMIT];
    fptr = fopen("input.txt", "r");

    printf("file has been opened\n");
    // putting contents of file into a array
    for (int i = 0; i < LINECOUNT; i++)
    {
        if (fgets(fileContents[i], LINELIMIT, fptr) == NULL)
        {
            break;
        }
        // convert new line char into a null terminator except last entry as that doesnt have a new line
        if (i != (LINECOUNT - 1))
        {
            fileContents[i][strlen(fileContents[i]) - 1] = '\0';
        }
    }

    int total = 0;

    for (int x = 0; x < LINECOUNT; x++)
    {
        // Example input: Game 100: 8 green; 2 red, 20 green; 12 green, 1 red, 1 blue; 4 red, 1 blue; 1 blue, 6 red
        char *strPtr = strtok(fileContents[x], " ");
        // Fisrt skip past 'Game'
        strPtr = strtok(NULL, " ");
        // Obtain number and remove : from end and convert to int
        char *lineNum = strPtr;
        lineNum[strlen(lineNum) - 1] = '\0';
        int num = atoi(lineNum);

        int redNum = 0, blueNum = 0, greenNum = 0;

        while ((strPtr = strtok(NULL, " ")) != NULL)
        {

            char *ballNum = strPtr;

            strPtr = strtok(NULL, " ");
            char *ballCol = strPtr;

            if (ballCol[strlen(ballCol) - 1] == ';' || ballCol[strlen(ballCol) - 1] == ',')
            {
                ballCol[strlen(ballCol) - 1] = '\0';
            }

            // Get number of balls for each colour to get the max number of each colour
            if ((strcmp(ballCol, "red") == 0) && atoi(ballNum) > redNum)
            {
                redNum = atoi(ballNum);
            }
            if ((strcmp(ballCol, "green") == 0) && atoi(ballNum) > greenNum)
            {
                greenNum = atoi(ballNum);
            }
            if ((strcmp(ballCol, "blue") == 0) && atoi(ballNum) > blueNum)
            {
                blueNum = atoi(ballNum);
            }
        }
        printf("For line %d blue:%d, red %d, green %d\n", num, blueNum, redNum, greenNum);

        total = total + (blueNum * redNum * greenNum);
    }

    printf("The final total is %d", total);
    fclose(fptr);
    return 0;
}