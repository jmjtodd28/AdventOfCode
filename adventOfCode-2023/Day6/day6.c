#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define LINECOUNT 2
#define LINELENGTH 40
#define ARRAYLENGTH 4

void populateArrays(int *arr, char *line)
{
    char *token = strtok(line, ":");
    int x = 0;
    while ((token = strtok(NULL, " ")) != NULL)
    {
        arr[x] = atoi(token);
        x++;
    }
}

int main()
{
    FILE *fptr;
    char fileContents[LINECOUNT][LINELENGTH];
    // CHANGE THE ARRAYLENGTH DEPEDING ON INPUT, IF INTPUT.TXT = 4, TEST_INPUT.TXT = 3
    fptr = fopen("input.txt", "r");

    printf("file has been opened\n");

    // puts contents of file into an array
    for (int i = 0; i < LINECOUNT; i++)
    {
        if (fgets(fileContents[i], LINELENGTH, fptr) == NULL)
        {
            break;
        }
        // convert new line char into a null terminator
        if (fileContents[i][strlen(fileContents[i]) - 1] == '\n')
        {
            fileContents[i][strlen(fileContents[i]) - 1] = '\0';
        }
    }

    int times[ARRAYLENGTH];
    int distance[ARRAYLENGTH];

    printf("%lu\n", sizeof(fileContents[0]));

    populateArrays(times, fileContents[0]);
    populateArrays(distance, fileContents[1]);

    int totalNumOfWays = 1;

    for (int y = 0; y < ARRAYLENGTH; y++)
    {
        int numOfWays = 0;
        for (int z = 0; z < times[y]; z++)
        {
            int holdTime = z;
            int dist = (times[y] - holdTime) * holdTime;

            if (dist > distance[y])
            {
                numOfWays += 1;
            }
        }
        totalNumOfWays *= numOfWays;
    }

    printf("Final total: %d", totalNumOfWays);

    fclose(fptr);
    return 0;
}