#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define LINECOUNT 2
#define LINELENGTH 40
#define ARRAYLENGTH 4

long long constructNums(char *line)
{

    char *token = strtok(line, ":");
    int x = 0;
    long long fullNum = 0;
    while ((token = strtok(NULL, " ")) != NULL)
    {
        int num = atoi(token);
        if (num < 10)
        {
            fullNum = fullNum * 10 + num;
        }
        else if (num > 9 && num < 100)
        {
            fullNum = fullNum * 100 + num;
        }
        else if (num > 99 && num < 1000)
        {
            fullNum = fullNum * 1000 + num;
        }
        else
        {
            fullNum = fullNum * 10000 + num;
        }
        x++;
    }
    return fullNum;
}

int main()
{
    FILE *fptr;
    char fileContents[LINECOUNT][LINELENGTH];
    fptr = fopen("input.txt", "r");

    printf("The file has been opened\n");

    for (int i = 0; i < LINECOUNT; i++)
    {
        if (fgets(fileContents[i], LINELENGTH, fptr) == NULL)
        {
            break;
        }

        if (fileContents[i][strlen(fileContents[i]) - 1] == '\n')
        {
            fileContents[i][strlen(fileContents[i]) - 1] = '\0';
        }
    }

    long long time, distance;

    time = constructNums(fileContents[0]);
    distance = constructNums(fileContents[1]);

    printf("Time: %lld\tDistance: %lld\n", time, distance);

    int numOfWays = 0;
    for (int x = 0; x < time; x++)
    {
        int holdTime = x;
        long long dist = (time - holdTime) * holdTime;

        if (dist > distance)
        {
            numOfWays += 1;
        }
    }

    printf("final number of ways: %d", numOfWays);

    fclose(fptr);
    return 0;
}