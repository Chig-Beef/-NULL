#ifndef LEVEL_H_
#define LEVEL_H_

#include "Player.h"
#include "Tile.h"

typedef struct Level {
  Tile data[96][16];
} Level;

void levelDraw(Level *l, double yOffset);

void loadLevel(/*string fName, */ Level *retLevel, Player *p);

#endif
