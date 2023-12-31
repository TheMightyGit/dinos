package cartridge

import (
	"embed"
	"image"
	"math/rand"

	"github.com/TheMightyGit/marv/marvtypes"
)

//go:embed "resources/*"
var Resources embed.FS

const (
	GfxBankFont = iota
	GfxBankDinos
	GfxBankBG
)
const (
	MapBankDinos = iota
)
const (
	MapBankDinoAnimsArea = 0
	MapBankBGArea        = 1
)
const (
	SpriteBG1 = iota
	SpriteBG2
	SpriteStart
	SpriteEnd = 127
)

var (
	API marvtypes.MarvAPI
	bg1 marvtypes.Sprite
	bg2 marvtypes.Sprite
)

func randomArenaStartPos(colour int) image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: Arena.Min.X + (colour * Arena.Size().X / 4) + rand.Intn(Arena.Size().X/4),
			Y: Arena.Min.Y + rand.Intn(Arena.Size().Y),
		},
		Max: DinoSize,
	}
}

func Start(api marvtypes.MarvAPI) {
	API = api

	bg1 = API.SpritesGet(SpriteBG1)
	bg1.ChangePos(image.Rect(0, 0, 320, 200))
	bg1.Show(GfxBankBG, API.MapBanksGet(MapBankDinos).GetArea(MapBankBGArea))

	bg2 = API.SpritesGet(SpriteBG2)
	bg2.ChangePos(image.Rect(320, 0, 320, 200))
	bg2.Show(GfxBankBG, API.MapBanksGet(MapBankDinos).GetArea(MapBankBGArea))

	for i := SpriteStart; i <= SpriteEnd; i++ {
		colour := rand.Intn(4)
		Dinos = append(Dinos, NewDino(API.SpritesGet(i), randomArenaStartPos(colour), colour))
	}
	// marv.ModBanks[0].Play()
	API.SfxBanksGet(0).PlayLooped()
}

var (
	bgX           int
	bgScrollScale = 4
)

func Update() {
	bg1.ChangePos(image.Rect((bgX / bgScrollScale), 0, 320, 200))
	bg2.ChangePos(image.Rect((bgX/bgScrollScale)+320, 0, 320, 200))
	bgX--
	if bgX == (-320 * bgScrollScale) {
		bgX = 0
	}
	for _, dino := range Dinos {
		dino.Update()
	}
	API.SpritesSort()
}
