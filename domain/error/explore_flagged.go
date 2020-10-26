package error

type exploreFlagged struct {
}

func (ef *exploreFlagged) Error() string {
	return "can't explore flagged cell"
}

var ExploreFlagged = &exploreFlagged{}
