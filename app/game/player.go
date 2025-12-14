package game

type Player struct {
	Warrior    Warrior
	lastHealth int
}

func (p *Player) PlayTurn(game *Game) {
	feel := p.Warrior.Feel(*game)
	if feel.monster {
		p.Warrior.Attack(game)
	} else if p.Warrior.Health < MAX_HEALTH && p.lastHealth == p.Warrior.Health {
		p.Warrior.Heal()
	} else {
		p.Warrior.Walk(*game)
	}
	p.lastHealth = p.Warrior.Health
}

func NewPlayer(name string) *Player {
	Coordinates := Coordinates{X: 0, Y: 0}
	warrior := NewWarrior(name, Coordinates)
	p := Player{Warrior: *warrior}
	return &p
}
