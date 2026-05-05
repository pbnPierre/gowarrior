package game

type Unit interface {
	AttackPower() int
	IsFoe() bool
	IsCaptive() bool
	Health() int
	Attacked(power int)
	Coordinates() Coordinates
	Name() string
	ToChar() string
	PerformTurn(game *Game)
}
