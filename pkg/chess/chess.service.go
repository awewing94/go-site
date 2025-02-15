package chess

import (
	"fmt"
	"go-site/pkg/global"
)

var Moves = []string{}

func makeMove(move string) {
    if global.Debug {
        fmt.Println("chessService.makeMove():")
    }

    Moves = append(Moves, move)
}

func getMoves() []string {
    if global.Debug {
        fmt.Println("chessService.getMoves():")
    }

    return Moves
}

