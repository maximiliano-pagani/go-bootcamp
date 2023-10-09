package simulator

type MockCatchSimulator struct {
	ValueCanCatch          bool
	ValueGetLinearDistance float64
	CallsGetLinearDistance []struct {
		position [2]float64
		distance float64
	}
}

func (s *MockCatchSimulator) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	return s.ValueCanCatch
}

func (s *MockCatchSimulator) GetLinearDistance(position [2]float64) float64 {
	s.CallsGetLinearDistance = append(s.CallsGetLinearDistance,
		struct {
			position [2]float64
			distance float64
		}{
			position: position,
			distance: s.ValueGetLinearDistance,
		},
	)
	return s.ValueGetLinearDistance
}
