package material

type Stimuli struct {
	Mask  string
	Array []string
}

func (s *Stimuli) Load() {
	s.Mask = "material/Mask.png"
	s.Array = append(s.Array, "material/P1.jpg")
	s.Array = append(s.Array, "material/P2.jpg")
}
