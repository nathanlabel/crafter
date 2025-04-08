package assets

import rl "github.com/gen2brain/raylib-go/raylib"

type FrameSet struct {
	Frames  []rl.RectangleInt32
	Time    float32
	TileSet *Tileset
}
