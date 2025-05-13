#ifndef INPUTCODES_H_
#define INPUTCODES_H_

#include <corecrt.h>
#include <stdbool.h>
#include <stdint.h>

#define MAX_FRAMES 4096
#define MAX_INPUTS 7

typedef enum InputCode {
  NO_INPUT,
  LEFT,
  RIGHT,
  JUMP,
  BOOST,
  NUM_INPUT_CODES,
} InputCode;

extern uint64_t runTime;
extern int runInputLen;

extern InputCode runInputs[MAX_FRAMES][MAX_INPUTS];

bool isInputDown(InputCode b);

errno_t loadInputs(char *fileName);

void addInput(InputCode b);

errno_t saveInputs();

void resetInputs();

#endif
