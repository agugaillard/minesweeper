package error

type invalidPosition struct{}

func (ip *invalidPosition) Error() string {
	return "invalid position"
}

var InvalidPosition = &invalidPosition{}
