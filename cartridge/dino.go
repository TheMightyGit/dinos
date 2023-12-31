package cartridge

import (
	"image"
	"math/rand"

	"github.com/TheMightyGit/marv/marvtypes"
)

type Dino struct {
	rect   image.Rectangle
	colour int
	sprite marvtypes.Sprite
	frame  float64
	anim   []image.Point
	dY     int
	dX     int
}

func (d *Dino) Show() {
	d.sprite.SetSortIdx(d.rect.Min.Y)
	d.sprite.ChangePos(d.rect)
	d.sprite.Show(GfxBankDinos, API.MapBanksGet(MapBankDinos).GetArea(MapBankDinoAnimsArea))
	d.sprite.ChangePalette(rand.Intn(4))
	// d.anim = Anims[rand.Intn(len(Anims))]
	if rand.Float64() <= 0.5 {
		d.anim = AnimWalk
		d.dX = 1
	} else {
		d.anim = AnimRun
		d.dX = 2
	}
}

func (d *Dino) Update() {
	if d.dX == 2 || rand.Float64() < 0.4 {
		d.rect.Min = d.rect.Min.Add(image.Point{1, d.dY})
		if d.rect.Min.X > 320+DinoSize.X {
			d.rect.Min.X = -DinoSize.X
		}
		d.sprite.ChangePos(d.rect)
		d.sprite.SetSortIdx(d.rect.Min.Y)
	}
	fn := int(d.frame) % len(d.anim)
	animFrame := d.anim[fn]

	if fn == 0 {
		if d.dX == 1 && rand.Float64() <= 0.2 {
			d.dY = rand.Intn(3) - 1
		}
		if rand.Float64() <= 0.1 {
			if rand.Float64() <= 0.8 {
				d.anim = AnimWalk
				d.dX = 1
			} else {
				d.anim = AnimRun
				d.dX = 2
				d.dY = 0
			}
		}
	}
	if d.rect.Min.Y < 96 {
		d.dY = 1
	} else if d.rect.Min.Y > 200-DinoSize.Y {
		d.dY = -1
	}

	d.sprite.ChangeViewport(
		image.Point{
			animFrame.X * DinoSize.X,
			((animFrame.Y * 4) + d.colour) * DinoSize.Y,
		},
	)
	d.frame += 0.2
}

var (
	DinoSize   = image.Point{24, 18}
	Dinos      []*Dino
	Arena      = image.Rectangle{Min: image.Point{Y: 96}, Max: image.Point{320, 200 - DinoSize.Y}}
	AnimIdle   = []image.Point{image.Point{0, 0}, image.Point{1, 0}, image.Point{2, 0}, image.Point{3, 0}}
	AnimWalk   = []image.Point{image.Point{0, 1}, image.Point{1, 1}, image.Point{2, 1}, image.Point{3, 1}, image.Point{4, 1}, image.Point{5, 1}}
	AnimKick   = []image.Point{image.Point{0, 2}, image.Point{1, 2}, image.Point{2, 2}}
	AnimHurt   = []image.Point{image.Point{0, 3}, image.Point{1, 3}, image.Point{2, 3}, image.Point{3, 3}}
	AnimCrouch = []image.Point{image.Point{0, 4}}
	AnimRun    = []image.Point{image.Point{0, 5}, image.Point{1, 5}, image.Point{2, 5}, image.Point{3, 5}, image.Point{4, 5}, image.Point{5, 5}}

	Anims = [][]image.Point{
		AnimIdle,
		AnimWalk,
		AnimKick,
		AnimHurt,
		AnimCrouch,
		AnimRun,
	}
)

func NewDino(sprite marvtypes.Sprite, pos image.Rectangle, colour int) *Dino {
	dino := &Dino{
		sprite: sprite,
		rect:   pos,
		colour: colour,
		frame:  rand.Float64(), // so we're not all in the same place of each anim
	}
	dino.Show()
	return dino
}
