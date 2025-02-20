package chess

import (
	"fmt"
	"go-site/pkg/global"
)

var CurrentGame = Game{}
// Chess game/board in the db will be composed of
//  - List of moves
//  - List of pieces (and their state)

func makeMove(move string) {
    if global.Debug {
        fmt.Println("service.makeMove():")
    }

    // TODO - parse the move and move the piece

    Moves = append(Moves, move)
}

func getChessBoard() []Piece {
    if global.Debug {
        fmt.Println("service.getMoves():")
    }

    return Pieces
}

func initializeBoard() {
    Moves = []Move{}
    Pieces = []Piece{
        {
            Rank: 0,
            File: 0,
        },

    }
}
