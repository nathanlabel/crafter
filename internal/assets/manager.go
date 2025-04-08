package assets

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	DefaultAssets = "assets.json"
)

var (
	Manager *AssetManager
)

type AssetManager struct {
	tilesets  map[string]*Tileset
	frameSets map[string]*FrameSet
}

func InitManager() {
	Manager = newAssetManager()
}

func newAssetManager() *AssetManager {
	am := AssetManager{
		tilesets:  make(map[string]*Tileset),
		frameSets: make(map[string]*FrameSet),
	}
	am.LoadAssets()
	return &am
}

// LoadAssets will try and load all assets listed in a file called assets.txt (DefaultAssets) if the file
// doesnt exist or is erraneous, it may partially load. Outputs to log
func (am *AssetManager) LoadAssets() {
	assetData, err := os.ReadFile(DefaultAssets)
	if err != nil {
		log.Printf("unable to read %s as sheets file", DefaultAssets)
		return
	}

	var assets Assets
	err = json.Unmarshal(assetData, &assets)
	if err != nil {
		log.Printf("error parsing json from %s", DefaultAssets)
	}

	for _, tileset := range assets.Tilesets {
		texture := rl.LoadTexture(tileset.Asset)
		ts := tileset
		ts.Texture = texture

		// Create Animations map
		ts.Animations = make(map[string][]int)
		for _, anim := range ts.AnimationsData {
			ts.Animations[anim.Name] = anim.Frames
		}

		am.tilesets[tileset.Name] = &ts

		log.Printf("Loaded Tileset: %s, Animations: %d, Name: %s", ts.Asset, len(ts.AnimationsData), ts.Name)
	}

	err = am.loadFrames()
	if err != nil {
		log.Printf("unable to load frames: %s", err)
	}
}

func (am *AssetManager) loadFrames() error {
	for name, tileset := range am.tilesets {
		frames_across := int(tileset.Texture.Width) / tileset.TileWidth
		frames_down := int(tileset.Texture.Height) / tileset.TileHeight

		fs := FrameSet{
			Frames:  make([]rl.RectangleInt32, frames_across*frames_down),
			Time:    tileset.FrameTime,
			TileSet: tileset,
		}
		i := 0
		for y := 0; y < frames_down; y++ {
			for x := 0; x < frames_across; x++ {
				i = int(y*frames_across + x)
				rec := rl.RectangleInt32{
					X:      int32(x * tileset.TileWidth),
					Y:      int32(y * tileset.TileHeight),
					Width:  int32(tileset.TileWidth),
					Height: int32(tileset.TileHeight),
				}
				fs.Frames[i] = rec
			}
		}
		am.frameSets[name] = &fs
		log.Printf("Loaded Frameset: %s, Frames: %d", name, len(fs.Frames))
	}

	return nil
}

func (am *AssetManager) GetTileset(name string) (*Tileset, error) {
	if ts, ok := am.tilesets[name]; ok {
		return ts, nil
	}
	return nil, fmt.Errorf("unable to find tileset %s", name)
}

func (am *AssetManager) GetAnimationFrames(assetName, animName string) (*[]int, error) {
	if ts, ok := am.tilesets[assetName]; ok {
		if anim, ok := ts.Animations[animName]; ok {
			return &anim, nil
		}
		return nil, fmt.Errorf("unable to find animation %s", animName)
	}
	return nil, fmt.Errorf("unable to find tileset %s", assetName)
}
