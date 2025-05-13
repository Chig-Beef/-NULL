#include "include/raylib.h"

#include "Config.h"
#include "Images.h"
#include "InputCodes.h"
#include "Level.h"
#include "Phase.h"
#include "Player.h"
#include "Tile.h"
#include "TileCodes.h"

#include <stdint.h>
#include <stdlib.h>
#include <string.h>

Phase phase = GAME;

char endText[9];

const Color buttonColor = {32, 32, 32, 255};

void getPlayerPosition(Player *p, Level *l) {
  for (int row = 0; row < 96; row++) {
    for (int col = 0; col < 16; col++) {
      if (l->data[row][col].code == PLAYER_START) {
        p->x = (double)(col * tileSize);
        p->y = (double)(row * tileSize);
      }
    }
  }
}

void createWinPage() {
  int milliseconds = (int)((double)runTime * 1000.0F / 60.0F);
  int seconds = (int)runTime / 60;
  int minutes = seconds / 60;

  // limit numbers (effectively a mod)
  milliseconds -= seconds * 1000;
  seconds -= minutes * 60;

  // NOTE: Assuming the text format M:SS.MMM

  // Minutes
  endText[0] = minutes + '0';
  endText[1] = ':';

  // Seconds
  endText[2] = seconds / 10 + '0';
  endText[3] = seconds % 10 + '0';
  endText[4] = '.';

  // Milliseconds
  endText[5] = milliseconds / 100 + '0';
  endText[6] = milliseconds / 10 % 10 + '0';
  endText[7] = milliseconds % 10 + '0';

  // End of string
  endText[8] = 0;
}

bool updatePlayer(Player *p, Level *l) {
  // gravity
  p->dy++;

  // Jump
  if (IsKeyDown(KEY_W) || IsKeyDown(KEY_SPACE)) {
    addInput(JUMP);

    if (!p->inAir) {
      p->dy -= 32;
      p->inAir = true;
    }
  }

  double nx = p->x;
  double ny = p->y + p->dy;

  Tile t;

  // Boosting
  for (int row = 0; row < 96; row++) {
    for (int col = 0; col < 16; col++) {
      t = l->data[row][col];

      if (!t.solid) {
        continue;
      }

      if (t.code != UPWARD_BOOST) {
        continue;
      }

      // Collision check
      if ((int)(nx) >= tileSize * (col + 1) ||
          (int)(nx + p->w) <= tileSize * col ||
          (int)(ny) >= tileSize * (row + 1) ||
          (int)(ny + p->h) <= tileSize * row) {
        continue;
      }

      p->dy += t.by;
    }
  }

  nx = p->x;
  ny = p->y + p->dy;

  // At some point use nx and ny as indecies into the 2D array
  // rather than searching everything for no reason
  for (int row = 0; row < 96; row++) {
    for (int col = 0; col < 16; col++) {
      t = l->data[row][col];

      if (!t.solid) {
        continue;
      }

      // Collision check
      if ((int)(nx) >= tileSize * (col + 1) ||
          (int)(nx + p->w) <= tileSize * col ||
          (int)(ny) >= tileSize * (row + 1) ||
          (int)(ny + p->h) <= tileSize * row) {
        continue;
      }

      switch (t.code) {
      case UPWARD_BOOST:
        continue;
      case END_LEVEL:
        phase = WIN;
        createWinPage();
        return false;
      }

      // Was above and moving down
      if (p->dy >= 0) {
        p->inAir = false;
        p->dy = 0;
        p->y = (double)(tileSize * row) - p->h;
      } else {
        p->dy = 0;
        p->y = (double)(tileSize * (row + 1));
      }

      goto DONE_VER_COLLISION_CHECK;
    }
  }

DONE_VER_COLLISION_CHECK:

  p->y += p->dy;

  // Horizontal movement

  double boost = 1.0;
  if (IsKeyDown(KEY_LEFT_SHIFT) || IsKeyDown(KEY_RIGHT_SHIFT)) {
    boost *= 2;
    addInput(BOOST);
  }

  if (IsKeyDown(KEY_A)) {
    p->dx = -(p->speed * boost);
    addInput(LEFT);
  }

  if (IsKeyDown(KEY_D)) {
    p->dx = p->speed * boost;
    addInput(RIGHT);
  }

  nx = p->x + p->dx;
  ny = p->y;

  for (int row = 0; row < 96; row++) {
    for (int col = 0; col < 16; col++) {
      t = l->data[row][col];

      if (!t.solid) {
        continue;
      }

      // Collision check
      if ((int)(nx) >= tileSize * (col + 1) ||
          (int)(nx + p->w) <= tileSize * col ||
          (int)(ny) >= tileSize * (row + 1) ||
          (int)(ny + p->h) <= tileSize * row) {
        continue;
      }

      switch (t.code) {
      case UPWARD_BOOST:
        continue;
      case END_LEVEL:
        phase = WIN;
        createWinPage();
        return false;
      }

      // Was above and moving down
      if (p->dx >= 0) {
        p->dx = 0;
        p->x = (double)(tileSize * col) - p->w;
      } else {
        p->dx = 0;
        p->x = (double)(tileSize * (col + 1));
      }

      goto DONE_HOR_COLLISION_CHECK;
    }
  }

DONE_HOR_COLLISION_CHECK:

  p->x += p->dx;
  p->dx = 0;

  if (p->x < 0) {
    p->x = 0;
  } else if (p->x + p->w > (double)screenWidth) {
    p->x = (double)screenWidth - p->w;
  }

  if (p->y < 0) {
    p->y = 0;
  }

  return false;
}

bool updatePlayerReplay(Player *p, Level *l) {
  // gravity
  p->dy++;

  // Jump
  if (isInputDown(JUMP)) {
    if (!p->inAir) {
      p->dy -= 32;
      p->inAir = true;
    }
  }

  double nx = p->x;
  double ny = p->y + p->dy;

  Tile t;

  // Boosting
  for (int row = 0; row < 96; row++) {
    for (int col = 0; col < 16; col++) {
      t = l->data[row][col];

      if (!t.solid) {
        continue;
      }

      if (t.code != UPWARD_BOOST) {
        continue;
      }

      // Collision check
      if ((int)(nx) >= tileSize * (col + 1) ||
          (int)(nx + p->w) <= tileSize * col ||
          (int)(ny) >= tileSize * (row + 1) ||
          (int)(ny + p->h) <= tileSize * row) {
        continue;
      }

      p->dy += t.by;
    }
  }

  nx = p->x;
  ny = p->y + p->dy;

  // At some point use nx and ny as indecies into the 2D array
  // rather than searching everything for no reason
  for (int row = 0; row < 96; row++) {
    for (int col = 0; col < 16; col++) {
      t = l->data[row][col];

      if (!t.solid) {
        continue;
      }

      // Collision check
      if ((int)(nx) >= tileSize * (col + 1) ||
          (int)(nx + p->w) <= tileSize * col ||
          (int)(ny) >= tileSize * (row + 1) ||
          (int)(ny + p->h) <= tileSize * row) {
        continue;
      }

      switch (t.code) {
      case UPWARD_BOOST:
        continue;
      case END_LEVEL:
        phase = WIN;
        createWinPage();
        return false;
      }

      // Was above and moving down
      // if int(p.y+p.h) < row*tileSize {
      if (p->dy >= 0) {
        p->inAir = false;
        p->dy = 0;
        p->y = (double)(tileSize * row) - p->h;
      } else {
        p->dy = 0;
        // p.dy = -p.dy
        p->y = (double)(tileSize * (row + 1));
      }

      goto DONE_VER_COLLISION_CHECK;
    }
  }

DONE_VER_COLLISION_CHECK:

  p->y += p->dy;

  // Horizontal movement

  double boost = 1.0;
  if (isInputDown(BOOST)) {
    boost *= 2;
  }

  if (isInputDown(LEFT)) {
    p->dx = -(p->speed * boost);
  }

  if (isInputDown(RIGHT)) {
    p->dx = p->speed * boost;
  }

  nx = p->x + p->dx;
  ny = p->y;

  for (int row = 0; row < 96; row++) {
    for (int col = 0; col < 16; col++) {
      t = l->data[row][col];

      if (!t.solid) {
        continue;
      }

      // Collision check
      if ((int)(nx) >= tileSize * (col + 1) ||
          (int)(nx + p->w) <= tileSize * col ||
          (int)(ny) >= tileSize * (row + 1) ||
          (int)(ny + p->h) <= tileSize * row) {
        continue;
      }

      switch (t.code) {
      case UPWARD_BOOST:
        continue;
      case END_LEVEL:
        phase = WIN;
        createWinPage();
        return false;
      }

      // Was above and moving down
      // if int(p.y+p.h) < row*tileSize {
      if (p->dx >= 0) {
        p->dx = 0;
        p->x = (double)(tileSize * col) - p->w;
      } else {
        // p.dy = 0
        p->dx = 0;
        p->x = (double)(tileSize * (col + 1));
      }

      goto DONE_HOR_COLLISION_CHECK;
    }
  }

DONE_HOR_COLLISION_CHECK:

  p->x += p->dx;
  p->dx = 0;

  if (p->x < 0) {
    p->x = 0;
  } else if (p->x + p->w > (double)screenWidth) {
    p->x = (double)screenWidth - p->w;
  }

  if (p->y < 0) {
    p->y = 0;
  }

  return false;
}

void drawPlayer(Player *p) {
  DrawTexture(playerTexture, (int)p->x, halfHeight - (int)(p->h / 2), WHITE);
}

void update(Player *p, Level *l) {
  switch (phase) {
  case GAME:
    updatePlayer(p, l);
    break;
  case REPLAY:
    updatePlayerReplay(p, l);
    break;
  case WIN:
    int x = GetMouseX();
    int y = GetMouseY();

    if (IsMouseButtonPressed(MOUSE_BUTTON_LEFT)) {
      if (x > 10 && x < 110) { // Back
        if (y > screenHeight - 60 && y < screenHeight - 10) {
          phase = GAME;
          p->dx = 0;
          p->dy = 0;
          getPlayerPosition(p, l);
          resetInputs();
        }
      } else if (x > 120 && x < 220) { // Save
        if (y > screenHeight - 60 && y < screenHeight - 10) {
          errno_t err = saveInputs();
          if (err != 0) {
            exit(err);
          }
        }
      }
    }

    break;
  }
}

void draw(Player *p, Level *l) {
  BeginDrawing();
  ClearBackground((Color){0, 0, 0, 255});

  switch (phase) {
  case GAME:
  case REPLAY:
    levelDraw(l, p->y + (p->h / 2) - (double)halfHeight);
    drawPlayer(p);
    runTime++;
    break;

  case WIN:
    DrawText(endText, 50, 50, 50, WHITE);
    DrawRectangle(10, screenHeight - 60, 100, 50, GRAY);  // Restart
    DrawRectangle(120, screenHeight - 60, 100, 50, GRAY); // Save
    break;
  }

  EndDrawing();
}

int main(int argc, char *argv[]) {
  // Set inputs to NO INPUTS
  for (int i = 0; i < 4096; i++) {
    memset(runInputs[i], NO_INPUT, 7);
  }

  SetTraceLogLevel(LOG_NONE);

  InitWindow(screenWidth, screenHeight, "*NULL");
  SetTargetFPS(60);

  playerImg = LoadImage("assets/images/player.png");
  ImageResize(&playerImg, tileSize, tileSize);
  playerTexture = LoadTextureFromImage(playerImg);

  blockImg = LoadImage("assets/images/block.png");
  ImageResize(&blockImg, tileSize, tileSize);
  blockTexture = LoadTextureFromImage(blockImg);

  boostImg = LoadImage("assets/images/upwardBoost.png");
  ImageResize(&boostImg, tileSize, tileSize);
  boostTexture = LoadTextureFromImage(boostImg);

  Player p = (Player){0, 0, tileSize, tileSize, BLUE, 0, 0, 10, false};

  Level l;

  loadLevel(&l, &p);

  // Attempt replay-mode
  if (argc == 2) {
    phase = REPLAY;

    errno_t err = loadInputs(argv[1]);
    if (err != 0) {
      return err;
    }
  }

  while (!WindowShouldClose()) {
    update(&p, &l);
    draw(&p, &l);
    // printf("%d\n", GetFPS());
  }

  UnloadTexture(playerTexture);
  UnloadImage(playerImg);
  UnloadTexture(blockTexture);
  UnloadImage(blockImg);
  UnloadTexture(boostTexture);
  UnloadImage(boostImg);

  CloseWindow();

  return 0;
}
