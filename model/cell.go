package model

type Cell struct {
	Mined    bool
	Flagged  flag
	Revealed bool
}

func NewMinedCell() *Cell {
	cell := newBasicCell()
	cell.Mined = true
	return cell
}

func NewUnminedCell() *Cell {
	return newBasicCell()
}

func newBasicCell() *Cell {
	return &Cell{Flagged: None}
}
