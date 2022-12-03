package material

import (
	"strconv"
)

type Stimuli struct {
	Mask  string
	Array []string
}

func (s *Stimuli) Load() {
	s.Mask = "material/Mask.png"

	prefix := "material/P"
	suffix := ".jpg"

	for i := 1; i <= 26; i++ {
		s.Array = append(s.Array, prefix+strconv.Itoa(i)+suffix)
	}
}
