/*
You play several games and record the information from each game (your puzzle input).
Each game is listed with its ID number (like the 11 in Game 11: ...) followed by a semicolon-separated list of subsets of cubes that were revealed from the bag (like 3 red, 5 green, 4 blue).

For example, the record of a few games might look like this:

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
In game 1, three sets of cubes are revealed from the bag (and then put back again). The first set is 3 blue cubes and 4 red cubes; the second set is 1 red cube, 2 green cubes, and 6 blue cubes; the third set is only 2 green cubes.

The Elf would first like to know which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?

In the example above, games 1, 2, and 5 would have been possible if the bag had been loaded with that configuration. However, game 3 would have been impossible because at one point the Elf showed you 20 red cubes at once;
similarly, game 4 would also have been impossible because the Elf showed you 15 blue cubes at once. If you add up the IDs of the games that would have been possible, you get 8.

Determine which games would have been possible if the bag had been loaded with only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?
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
        // convert new line char into a null terminator
        fileContents[i][strlen(fileContents[i]) - 1] = '\0';
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

        while ((strPtr = strtok(NULL, " ")) != NULL)
        {

            char *ballNum = strPtr;

            strPtr = strtok(NULL, " ");
            char *ballCol = strPtr;

            if (ballCol[strlen(ballCol) - 1] == ';' || ballCol[strlen(ballCol) - 1] == ',')
            {
                ballCol[strlen(ballCol) - 1] = '\0';
            }

            // check to see if ball limits are reached
            if ((strcmp(ballCol, "red") == 0) && atoi(ballNum) > 12)
            {
                total = total - num;
                break;
            }
            if ((strcmp(ballCol, "green") == 0) && atoi(ballNum) > 13)
            {
                total = total - num;
                break;
            }
            if ((strcmp(ballCol, "blue") == 0) && atoi(ballNum) > 14)
            {
                total = total - num;
                break;
            }
        }
        total = total + num;
    }

    printf("The final total is %d", total);
    fclose(fptr);
    return 0;
}