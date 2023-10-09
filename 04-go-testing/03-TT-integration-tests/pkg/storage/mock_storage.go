package storage

type mockStorage struct {
	KeysGetValue map[string]float64
	ErrGetValue  error
}

func NewMockStorage(keys map[string]float64) Storage {
	return &mockStorage{KeysGetValue: keys}
}

func (s *mockStorage) GetValue(key string) interface{} {
	if s.ErrGetValue != nil {
		return nil
	}

	if v, ok := s.KeysGetValue[key]; ok {
		return v
	}

	return nil
}
