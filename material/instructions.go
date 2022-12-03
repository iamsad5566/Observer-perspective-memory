package material

type Instructions struct {
	Begin       string
	Description string
	PressSpace  string
	Prepare     string
	Response    string
	End         string
}

func (instructs *Instructions) LoadEng() {
	instructs.Begin = "This experiment is about the memory research. You will see 20 different pictures. Please remember each of them as much as possible."
	instructs.Description = "Please read the following description: there are two ways of remembering. One is first-person perspective and the other is third-person perspective. Take what you recall eating dinner last night as an example. When you use first-person perspective to remember, you will see plates and chopsticks, your hands…etc. when you use third-person perspective to remember, you will see yourself eating dinner. You may see your own back, or you may see your face."
	instructs.PressSpace = "Press the space key to proceed."
	instructs.Prepare = "The next page, the experiment starts, you will see continuously 20 pictures. Please try to remember the picture first as much as you can."
	instructs.Response = "Now, please use first-person perspective to remember the picture size. Then, adjust the picture size on the screen until it matches the size you see in your memory. After you have confirmed, press the space key to proceed."
	instructs.End = "The experiment is over. Thanks for your participation. Press space key to exit."
}
