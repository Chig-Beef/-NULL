#include "Level.h"
#include "Config.h"
#include "Images.h"
#include "Player.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void levelDraw(Level *l, double yOffset) {
  for (int row = 0; row < 96; row++) {
    for (int col = 0; col < 16; col++) {
      if (!l->data[row][col].solid) {
        continue;
      }

      int x = col * tileSize;
      int y = row * tileSize - (int)yOffset;

      switch (l->data[row][col].code) {
      case BLOCK:
        DrawTexture(blockTexture, x, y, WHITE);
        break;

      case UPWARD_BOOST:
        DrawTexture(boostTexture, x, y, WHITE);
        break;

      default:
        DrawRectangle(x, y, (int)tileSize, (int)tileSize, l->data[y][x].color);
        break;
      }
    }
  }
}

void loadLevel(Level *retLevel, Player *p) {
  FILE *fptr;
  errno_t err = fopen_s(&fptr, "assets/levels/level.txt", "r");
  if (err) {
    printf("%i\n", err);
    exit(err);
  }

  char data[4096];
  memset(data, 0, 4096);

  int rowIndex = 0;
  int colIndex = 0;

  while (fgets(data, 4096, fptr)) {
    // Start parsing
    for (int i = 0; i < 4096; i++) {
      if (data[i] == 0) {
        break;
      }

      switch (data[i]) {
      case ',':
        colIndex++;
        break;

      case '\n':
        // Do nothing
        break;

      case '\r':
        // Do nothing
        break;

      default:
        byte code = data[i] - '0';

        if (code > 9) {
          printf("Error in level loading, loaded code: %i\n", code);
          exit(1);
        }

        switch (code) {
        case PLAYER_START:
          p->x = (double)(colIndex * tileSize);
          p->y = (double)(rowIndex * tileSize);
          printf("%f\n", p->y);
          break;
        case NULL_SPAWN:
          code = NO_TILE;
          break;
        }

        // printf("%i\n", code);

        retLevel->data[rowIndex][colIndex] = (Tile){
            code,
            getTileColor(code),
            getTileSolid(code),
            getTileHorBoost(code),
            getTileVerBoost(code),
            0,
            0,
        };
      }
    }

    colIndex = 0;
    rowIndex++;

    memset(data, 0, 4096);
  }

  fclose(fptr);
}
