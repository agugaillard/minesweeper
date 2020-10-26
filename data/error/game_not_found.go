package error

type gameNotFound struct {
}

func (gae *gameNotFound) Error() string {
	return "game not found"
}

var GameNotFound = &gameNotFound{}
