package player

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	position rl.Vector2
}

func NewPlayer(pos rl.Vector2, textureName string) *Player {
	p := Player{
		position: rl.NewVector2(0, 0),
	}
	return &p
}

func (p *Player) Position() rl.Vector2 {
	return p.position
}

func (p *Player) SetPosition(pos rl.Vector2) {
	p.position = pos
}
