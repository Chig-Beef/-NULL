#ifndef PLAYER_H_
#define PLAYER_H_

#include "include/raylib.h"

typedef struct Player {
  double x, y;
  double w, h;

  // image?
  Color clr;

  double dx, dy;

  double speed;

  bool inAir;
} Player;

#endif
