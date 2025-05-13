package main

import (
	"encoding/json"
	"image"
	"null/stagerror"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Image Code
// Not used in future code

type ImageHandler struct {
	images map[string]*ebiten.Image
}

func (ih *ImageHandler) initImages() {
	ih.loadImages()
	ih.setIcon()
}

func (ih *ImageHandler) GetImage(name string) *ebiten.Image {
	i, ok := ih.images[name]
	if ok {
		return i
	}

	stagerror.New(32, "Couldn't find this image: "+name).SaveToLog(IS_RELEASE)

	// Will hard error if default doesn't exist
	return ih.images["missing"]
}

func (ih *ImageHandler) DrawImage(screen *ebiten.Image, img *ebiten.Image, x, y, w, h, rot float64, op *ebiten.DrawImageOptions) {
	ogW, ogH := img.Bounds().Dx(), img.Bounds().Dy()

	op.GeoM.Translate(-float64(ogW/2), -float64(ogH/2))
	op.GeoM.Rotate(rot)
	op.GeoM.Translate(float64(ogW/2), float64(ogH/2))

	op.GeoM.Scale(w/float64(ogW), h/float64(ogH))
	op.GeoM.Translate(x, y)
	screen.DrawImage(img, op)
}

func (ih *ImageHandler) setIcon() {
	ebiten.SetWindowIcon([]image.Image{ih.GetImage("icon")})
}

func (ih *ImageHandler) loadImages() {
	ih.images = make(map[string]*ebiten.Image)

	rawImageData, err := os.ReadFile("assets/images/images.json")
	if err != nil {
		stagerror.New(32, "Couldn't load images").SaveToLog(IS_RELEASE)
		return
	}

	var imageData [][]string
	err = json.Unmarshal(rawImageData, &imageData)
	if err != nil {
		stagerror.New(33, "Failed to load `./assets/images/images.json`, file may have been tampered with, reinstall advised").SaveToLog(IS_RELEASE)
		return
	}

	for i := 0; i < len(imageData); i++ {
		fName := imageData[i][0]
		mName := imageData[i][1]
		ih.loadImage(fName, mName)
	}
}

func (ih *ImageHandler) loadImage(fName string, mName string) {
	img, _, err := ebitenutil.NewImageFromFile("assets/images/" + fName)
	if err != nil {
		stagerror.New(34, "Couldn't load image: "+fName).SaveToLog(IS_RELEASE)
		return
	}
	ih.images[mName] = img
}
