package material

type Stimuli struct {
	Array []string
}

func (s *Stimuli) Load() {
	s.Array = append(s.Array, "material/back-arrow.png")
	s.Array = append(s.Array, "material/forward.png")
}
