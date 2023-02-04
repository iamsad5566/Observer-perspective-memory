package material

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

const prefix string = "material/I/I"
const suffix string = ".jpg"

var trialsNum int

type PictureFile struct {
	Mask           string
	CurrentPicture string
	Slice          []string
	OPSlice        []string
	FPSlice        []string
}

type InstructFile struct {
	Instructions []string
}

func (pic *PictureFile) Load() {
	prefixPic := "material/P/P"
	suffixPic := ".jpg"

	pic.Mask = "material/P/" + "Mask" + suffixPic
	pic.CurrentPicture = "material/P/" + "Mask" + suffixPic

	loadEnv()
	for i := 1; i <= trialsNum; i++ {
		pic.Slice = append(pic.Slice, prefixPic+strconv.Itoa(i)+suffixPic)
	}

	prefixPic = "material/OP/OP"
	for i := 1; i <= 6; i++ {
		pic.OPSlice = append(pic.OPSlice, prefixPic+strconv.Itoa(i)+suffixPic)
	}

	prefixPic = "material/FP/FP"
	for i := 1; i <= 6; i++ {
		pic.FPSlice = append(pic.FPSlice, prefixPic+strconv.Itoa(i)+suffixPic)
	}
}

func (pic *PictureFile) ShuffleSlice() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(pic.Slice), func(i, j int) { pic.Slice[i], pic.Slice[j] = pic.Slice[j], pic.Slice[i] })
}

func (instruct *InstructFile) Load() {
	for i := 1; i <= 16; i++ {
		instruct.Instructions = append(instruct.Instructions, prefix+strconv.Itoa(i)+suffix)
	}
}

func loadEnv() {
	trialsNum, _ = strconv.Atoi(os.Getenv("TRIALS_NUM"))
}
