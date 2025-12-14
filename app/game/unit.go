package game

type Unit interface {
	AttackPower() int
	Health() int
	Attacked(power int)
	Coordinates() Coordinates
	Name() string
	ToChar() string
	PerformTurn(game *Game)
}
