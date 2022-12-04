package material

import "fyne.io/fyne/v2/canvas"

type Instructions struct {
	Begin       *canvas.Text
	Description *canvas.Text
	PressSpace  *canvas.Text
	Prepare     *canvas.Text
	Response    *canvas.Text
	End         *canvas.Text
}

const (
	begin       string = "This experiment is about the memory research. You will see 20 different pictures. Please remember each of them as much as possible."
	description string = "Please read the following description: there are two ways of remembering. One is first-person perspective and the other is third-person perspective. \nTake what you recall eating dinner last night as an example. When you use first-person perspective to remember, you will see plates and chopsticks,\n your hands…etc. when you use third-person perspective to remember, you will see yourself eating dinner. \nYou may see your own back, or you may see your face."
	pressSpace  string = "Press the space key to proceed."
	prepare     string = "The next page, the experiment starts, you will see continuously 20 pictures. Please try to remember the picture first as much as you can."
	response    string = "Now, please use first-person perspective to remember the picture size. Then, adjust the picture size on the screen until it matches the size you see in your memory. After you have confirmed, press the space key to proceed."
	end         string = "The experiment is over. Thanks for your participation. Press space key to exit."
)

func (instructs *Instructions) LoadEng() {
	instructs.Begin = canvas.NewText(begin, nil)
	instructs.Description = canvas.NewText(description, nil)
	instructs.PressSpace = canvas.NewText(pressSpace, nil)
	instructs.Prepare = canvas.NewText(prepare, nil)
	instructs.Response = canvas.NewText(response, nil)
	instructs.End = canvas.NewText(end, nil)
}

func nextLine(str string)
