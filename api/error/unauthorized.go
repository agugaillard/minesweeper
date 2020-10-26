package error

type unauthorized struct{}

func (u *unauthorized) Error() string {
	return "unauthorized"
}

var Unauthorized = &unauthorized{}
