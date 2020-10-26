package error

type interactFinished struct{}

func (i *interactFinished) Error() string {
	return "interacting with finished game"
}

var InteractFinished = &interactFinished{}
