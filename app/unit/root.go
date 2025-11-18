package unit

type Unit interface {
	toUtf8Char() string
	toANSIChar() float64
	PerformTurn()
}
