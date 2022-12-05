package material

import (
	"strconv"
)

type Stimuli struct {
	Mask               string
	CurrentPicture     string
	Array              []string
	CurrentInstruction string
}

func (s *Stimuli) Load() {
	s.Mask = "material/Mask.png"
	s.CurrentInstruction = "material/Begin.png"
	s.CurrentPicture = "material/P1.jpg"

	prefix := "material/P"
	suffix := ".jpg"

	for i := 1; i <= 26; i++ {
		s.Array = append(s.Array, prefix+strconv.Itoa(i)+suffix)
	}
}
