package error

type invalidNumberOfMines struct {
}

func (inm *invalidNumberOfMines) Error() string {
	return "invalid number of mines"
}

var InvalidNumberOfMines = &invalidNumberOfMines{}
