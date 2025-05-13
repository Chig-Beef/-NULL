package main

import (
	"slices"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	x, y float64
	w, h float64

	img *ebiten.Image

	dx, dy float64

	speed float64

	inAir bool
}

func newPlayer(img *ebiten.Image) *Player {
	return &Player{
		x:     0,
		y:     0,
		w:     tileSize,
		h:     tileSize,
		speed: 10,
		img:   img,
	}
}

func (p *Player) jump(g *Game) {
	if p.inAir {
		return
	}

	g.audio.playEffect("jump")
	p.dy -= 32
	p.inAir = true
}

func (p *Player) bound(l *Level) {
	// Out of bounds check (HOR)
	if p.x < 0 {
		p.x = 0
	} else if p.x+p.w > float64(len(l.data[0])*tileSize) {
		p.x = float64(len(l.data[0])*tileSize) - p.w
	}

	// Out of bounds check (VER)
	if p.y < 0 {
		p.y = 0
	} else if p.y+p.h > float64(len(l.data)*tileSize) {
		p.y = float64(len(l.data)*tileSize) - p.h
	}
}

func (p *Player) assistedBoost(g *Game, l *Level, nx, ny float64) {
	// Boosting
	for row := range len(l.data) {
		for col, t := range l.data[row] {
			if !t.solid {
				continue
			}

			if t.code != UPWARD_BOOST {
				continue
			}

			// Collision check
			if int(nx) >= tileSize*(col+1) || int(nx+p.w) <= tileSize*col ||
				int(ny) >= tileSize*(row+1) || int(ny+p.h) <= tileSize*row {
				continue
			}

			p.dy += t.by
			g.audio.playEffect("boost")
		}
	}
}

// Returns true if the level was finished
func (p *Player) verCollision(g *Game, l *Level, nx, ny float64) bool {
	// OPTIMIZE:
	// At some point use nx and ny as indecies into the 2D array
	// rather than searching everything for no reason
	for row := 0; row < len(l.data); row++ {
		for col := 0; col < len(l.data[row]); col++ {
			t := l.data[row][col]

			if !t.solid {
				continue
			}

			// Collision check
			if int(nx) >= tileSize*(col+1) || int(nx+p.w) <= tileSize*col ||
				int(ny) >= tileSize*(row+1) || int(ny+p.h) <= tileSize*row {
				continue
			}

			switch t.code {
			case UPWARD_BOOST:
				continue
			case END_LEVEL:
				g.winPhase()
				return true
			}

			// Was above and moving down
			// if int(p.y+p.h) < row*tileSize {
			if p.dy >= 0 {
				p.inAir = false
				p.dy = 0
				p.y = float64(tileSize*row) - p.h
			} else {
				p.dy = 0
				// p.dy = -p.dy
				p.y = float64(tileSize * (row + 1))
			}

			return false
		}
	}

	return false
}

func (p *Player) horCollision(g *Game, l *Level, nx, ny float64) bool {
	// OPTIMIZE:
	for row := 0; row < len(l.data); row++ {
		for col := 0; col < len(l.data[row]); col++ {
			t := l.data[row][col]

			if !t.solid {
				continue
			}

			// Collision check
			if int(nx) >= tileSize*(col+1) || int(nx+p.w) <= tileSize*col ||
				int(ny) >= tileSize*(row+1) || int(ny+p.h) <= tileSize*row {
				continue
			}

			switch t.code {
			case UPWARD_BOOST:
				continue
			case END_LEVEL:
				g.winPhase()
				return true
			}

			// Was above and moving down
			// if int(p.y+p.h) < row*tileSize {
			if p.dx >= 0 {
				p.dx = 0
				p.x = float64(tileSize*col) - p.w
			} else {
				// p.dy = 0
				p.dx = 0
				p.x = float64(tileSize * (col + 1))
			}

			return false

		}
	}

	return false
}

// Will return true on segfault
func (p *Player) update(g *Game, l *Level) {
	// gravity
	p.dy++

	// Jump
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeySpace) {
		runInputs[runTime-1] = append(runInputs[runTime-1], JUMP)
		p.jump(g)
	}

	nx := p.x
	ny := p.y + p.dy

	p.assistedBoost(g, l, nx, ny)

	nx = p.x
	ny = p.y + p.dy

	// If the level ended, don't keep working
	if p.verCollision(g, l, nx, ny) {
		return
	}

	p.y += p.dy

	// Horizontal movement
	boost := 1.0
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		boost *= 2
		runInputs[runTime-1] = append(runInputs[runTime-1], BOOST)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.dx = -(p.speed * boost)
		runInputs[runTime-1] = append(runInputs[runTime-1], LEFT)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.dx = p.speed * boost
		runInputs[runTime-1] = append(runInputs[runTime-1], RIGHT)
	}

	nx = p.x + p.dx
	ny = p.y

	// If the level ended, don't keep working
	if p.horCollision(g, l, nx, ny) {
		return
	}

	// Execute the movement
	p.x += p.dx
	p.dx = 0

	p.bound(l)
	p.updateAudio(g)
}

// Will return true on segfault
func (p *Player) updateAsReplay(g *Game, l *Level) {
	// gravity
	p.dy++

	// Jump
	if slices.Contains(runInputs[runTime-1], JUMP) {
		p.jump(g)
	}

	nx := p.x
	ny := p.y + p.dy

	p.assistedBoost(g, l, nx, ny)

	nx = p.x
	ny = p.y + p.dy

	// If the level ended, don't keep working
	if p.verCollision(g, l, nx, ny) {
		return
	}

	p.y += p.dy

	// Horizontal movement
	boost := 1.0
	if slices.Contains(runInputs[runTime-1], BOOST) {
		boost *= 2
	}

	if slices.Contains(runInputs[runTime-1], LEFT) {
		p.dx = -(p.speed * boost)
	}
	if slices.Contains(runInputs[runTime-1], RIGHT) {
		p.dx = p.speed * boost
	}

	nx = p.x + p.dx
	ny = p.y

	// If the level ended, don't keep working
	if p.horCollision(g, l, nx, ny) {
		return
	}

	p.x += p.dx
	p.dx = 0

	p.bound(l)
	p.updateAudio(g)
}

func (p *Player) updateAudio(g *Game) {
	y := 5 - int(p.y/(16*tileSize))

	if y == 0 {
		return
	}

	// Turn on audio that we're at
	for i := 1; i <= y; i++ {
		name := "theme" + strconv.Itoa(i)
		if !slices.Contains(g.audio.curAudio, name) {
			g.audio.playAudio(name)
			g.audio.soundTracks[name].audioPlayer.SetPosition(g.audio.soundTracks["theme0"].audioPlayer.Position())
		}
	}

	// Turn off audio we shouldn't be hearing
	for i := y + 1; i <= 5; i++ {
		name := "theme" + strconv.Itoa(i)
		if slices.Contains(g.audio.curAudio, name) {
			g.audio.stopAudio(name)
			index := slices.Index(g.audio.curAudio, name)
			g.audio.curAudio = slices.Delete(g.audio.curAudio, index, index+1)
		}
	}
}

func (p *Player) draw(screen *ebiten.Image, g *Game) {
	g.image.DrawImage(screen, p.img, p.x, halfHeight-p.h/2, p.w, p.h, 0, &ebiten.DrawImageOptions{})
}
