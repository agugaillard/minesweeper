package error

type gameAlreadyExists struct{}

func (gae *gameAlreadyExists) Error() string {
	return "game already exists"
}

var GameAlreadyExists = &gameAlreadyExists{}
