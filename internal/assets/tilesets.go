package assets

import rl "github.com/gen2brain/raylib-go/raylib"

type Assets struct {
	Tilesets []Tileset `json:"tilesets"`
}

type Animation struct {
	Name   string `json:"name"`
	Frames []int  `json:"frames"`
}

type Tileset struct {
	Name           string           `json:"name"`
	Asset          string           `json:"asset"`
	TileWidth      int              `json:"tile_width"`
	TileHeight     int              `json:"tile_height"`
	FrameTime      float32          `json:"frame_time"`
	AnimationsData []Animation      `json:"animations"`
	Animations     map[string][]int `json:"-"`
	Texture        rl.Texture2D
}
