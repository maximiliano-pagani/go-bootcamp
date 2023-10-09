package prey

type StubPrey struct {
	Speed float64
}

func (s *StubPrey) GetSpeed() float64 {
	return s.Speed
}
