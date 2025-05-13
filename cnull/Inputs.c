#include "InputCodes.h"

#include <corecrt.h>
#include <stdbool.h>
#include <stdio.h>
#include <string.h>

uint64_t runTime = 0;
int runInputLen = 0;

InputCode runInputs[MAX_FRAMES][MAX_INPUTS];

char *inputCodeString(InputCode b) {
  switch (b) {
  case LEFT:
    return "LEFT";
  case RIGHT:
    return "RIGHT";
  case BOOST:
    return "BOOST";
  case JUMP:
    return "JUMP";
  default:
    return "";
  }
}

bool isInputDown(InputCode b) {
  for (int i = 0; i < MAX_INPUTS; i++) {
    if (runInputs[runTime][i] == b) {
      return true;
    }
  }

  return false;
}

bool strEqual(char *s1, char *s2) { return !strcmp(s1, s2); }

#define MAX_LINE_LEN 100

errno_t loadInputs(char *fileName) {
  FILE *fptr;
  char data[MAX_LINE_LEN];

  memset(data, 0, MAX_LINE_LEN);

  errno_t err = fopen_s(&fptr, fileName, "r");
  if (err) {
    return err;
  }

  while (fgets(data, MAX_LINE_LEN, fptr)) {
    int inputNum = 0;

    int cmdStart = 0;
    int cmdEnd = 0;

    for (; cmdEnd < MAX_LINE_LEN; cmdEnd++) {
      if (data[cmdEnd] == 0) {
        // EOL
      }

      if (data[cmdEnd] == ' ') {
        // Finished a word
        InputCode b;

        // Check for matches
        if (strEqual(data + cmdStart, "LEFT")) {
          b = LEFT;
        } else if (strEqual(data + cmdStart, "JUMP")) {
          b = JUMP;
        } else if (strEqual(data + cmdStart, "BOOST")) {
          b = BOOST;
        } else if (strEqual(data + cmdStart, "RIGHT")) {
          b = RIGHT;
        }

        runInputs[runInputLen][inputNum] = b;
        inputNum++;

        cmdStart = cmdEnd + 1;
      }
    }
    runInputLen++;

    // Reset the line
    memset(data, 0, MAX_LINE_LEN);
  }

  fclose(fptr);

  runInputLen++;

  return 0;
}

void addInput(InputCode b) {
  for (int i = 0; i < MAX_INPUTS; i++) {
    if (runInputs[runTime][i] == NO_INPUT) {
      runInputs[runTime][i] = b;
      return;
    }
  }
}

void resetInputs() {
  for (int row = 0; row < MAX_FRAMES; row++) {
    memset(runInputs[row], NO_INPUT, MAX_INPUTS);
  }
  runTime = 0;
}

errno_t saveInputs() {
  FILE *fptr;

  char data[MAX_LINE_LEN];
  char *inputName;

  memset(data, 0, MAX_LINE_LEN);

  // So we don't get any collisions
  remove("runc.nulldemo");

  errno_t err = fopen_s(&fptr, "runc.nulldemo", "w");
  if (err) {
    return err;
  }

  for (int frame = 0; frame < runTime; frame++) {
    int lineIndex = 0;

    for (int i = 0; i < MAX_INPUTS; i++) {
      // End of inputs
      if (runInputs[frame][i] == NO_INPUT) {
        break;
      }

      // Copy the name of the input into the string
      inputName = inputCodeString(runInputs[frame][i]);

      for (int p = 0; p < strlen(inputName); p++) {
        data[lineIndex] = inputName[p];
        lineIndex++;
      }

      data[lineIndex] = ' ';
      lineIndex++;
    }

    data[lineIndex] = '\n';

    fprintf(fptr, "%s", data);

    // Reset the string
    memset(data, 0, MAX_LINE_LEN);
  }
  fclose(fptr);

  printf("Successful write\n");

  return 0;
}
