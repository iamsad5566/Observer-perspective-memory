package win

type Stimulus struct {
	Current string
}

func (s *Stimulus) initialize(str string) {
	s.Current = str
}
