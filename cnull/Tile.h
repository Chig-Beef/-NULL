#ifndef TILE_H_
#define TILE_H_

#include "TileCodes.h"
#include "include/raylib.h"

typedef struct Tile {
  TileCode code;
  Color color;
  bool solid;

  // Boost
  double bx, by;

  // Where the tile points to
  int px, py;
} Tile;

bool getTileSolid(TileCode code);

Color getTileColor(TileCode code);

double getTileHorBoost(TileCode code);

double getTileVerBoost(TileCode code);

#endif
