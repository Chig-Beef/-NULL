#include "Tile.h"

bool getTileSolid(TileCode code) {
  switch (code) {
  case NO_TILE:
  case NULL_SPAWN:
  case PLAYER_START:
    return false;

  case BLOCK:
  case UPWARD_BOOST:
  case END_LEVEL:
    return true;

  default:
    return true;
  }
}

Color getTileColor(TileCode code) {
  switch (code) {
  case NO_TILE:
  case NULL_SPAWN:
  case PLAYER_START:
    return (Color){0, 0, 0, 0};

  case BLOCK:
    return (Color){255, 255, 255, 255};
  case UPWARD_BOOST:
    return (Color){255, 255, 0, 255};
  case END_LEVEL:
    return (Color){0, 0, 255, 255};

  // Bad Code
  default:
    return (Color){0, 0, 0, 0};
  }
}

double getTileHorBoost(TileCode code) {
  switch (code) {
  default:
    return 0;
  }
}

double getTileVerBoost(TileCode code) {
  switch (code) {
  case UPWARD_BOOST:
    return -5;
  default:
    return 0;
  }
}
