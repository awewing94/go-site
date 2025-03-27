package chess

import "strconv"

type Game struct {
	Pieces  []Piece
	Moves   []Move
}

type Piece struct {
	PieceType PieceType
    Color       Color
    Location    Location
    ClassNames  string
}

type Location struct {
	Rank int    // zero indexed
	File int    // a=0; b=1; ...
}

type Move struct {
	Start Location
	End   Location
}

type PieceType string

const (
	Pawn   PieceType = "pawn"
	King             = "king"
	Queen            = "queen"
	Rook             = "rook"
	Bishop           = "bishop"
	Knight           = "knight"
)

type Color rune

const (
    White Color = 'w'
    Black       = 'b'
)

func (location *Location) ToString() string {
    return string(location.File + 'a') + strconv.Itoa(int(location.Rank + 1))
}

