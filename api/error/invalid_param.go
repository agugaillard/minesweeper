package error

type invalidParameter struct{}

func (ip *invalidParameter) Error() string {
	return "invalid parameter"
}

var InvalidParameter = &invalidParameter{}
