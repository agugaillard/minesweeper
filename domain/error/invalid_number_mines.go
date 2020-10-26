package error

type invalidBoardProperties struct {
}

func (inm *invalidBoardProperties) Error() string {
	return "invalid number of mines"
}

var InvalidBoardProperties = &invalidBoardProperties{}
