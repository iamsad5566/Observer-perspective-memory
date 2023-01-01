package material

import (
	"math/rand"
	"strconv"
	"time"
)

const prefix string = "material/"
const suffix string = ".jpg"

type PictureFile struct {
	Mask           string
	CurrentPicture string
	Slice          []string
}

type InstructFile struct {
	Begin              string
	Description        string
	Prepare            string
	Response           string
	CurrentInstruction string
}

func (pic *PictureFile) Load() {
	pic.Mask = prefix + "Mask" + suffix
	pic.CurrentPicture = prefix + "Mask" + suffix

	prefixPic := "material/P"
	suffixPic := ".jpg"

	for i := 1; i <= 2; i++ {
		pic.Slice = append(pic.Slice, prefixPic+strconv.Itoa(i)+suffixPic)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(pic.Slice), func(i, j int) { pic.Slice[i], pic.Slice[j] = pic.Slice[j], pic.Slice[i] })
}

func (pic *PictureFile) ShuffleSlice() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(pic.Slice), func(i, j int) { pic.Slice[i], pic.Slice[j] = pic.Slice[j], pic.Slice[i] })
}

func (instruct *InstructFile) Load() {
	instruct.Begin = prefix + "Begin" + suffix
	instruct.Description = prefix + "Description" + suffix
	instruct.Prepare = prefix + "Prepare" + suffix
	instruct.Response = prefix + "Response" + suffix
}
