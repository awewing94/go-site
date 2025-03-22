package chess

import (
	"fmt"
	"go-site/pkg/global"
)

var CurrentGame = Game{}
// Chess game/board in the db will be composed of
//  - List of moves
//  - List of pieces (and their state)

func makeMove(moveString string) {
    if global.Debug {
        fmt.Println("service.makeMove():")
    }

    // TODO - parse the move and move the piece
    var move = Move{
        Start: Location{
            Rank: 0,
            File: 0,
        },
        End: Location{
            Rank: 1,
            File: 0,
        },
    }

    CurrentGame.Moves = append(CurrentGame.Moves, move)
}

func getPieces() []Piece {
    if global.Debug {
        fmt.Println("service.getMoves():")
    }

    if len(CurrentGame.Pieces) > 0{
        return CurrentGame.Pieces
    }

    initializeBoard()

    return CurrentGame.Pieces
}

func initializeBoard() {
    CurrentGame.Moves = []Move{}
    CurrentGame.Pieces = []Piece{
        // White Pieces
        Piece{ Location: Location{ Rank: 1, File: 0, }, PieceType: Pawn, Color: White, },
        Piece{ Location: Location{ Rank: 1, File: 1, }, PieceType: Pawn, Color: White, },
        Piece{ Location: Location{ Rank: 1, File: 2, }, PieceType: Pawn, Color: White, },
        Piece{ Location: Location{ Rank: 1, File: 3, }, PieceType: Pawn, Color: White, },
        Piece{ Location: Location{ Rank: 1, File: 4, }, PieceType: Pawn, Color: White, },
        Piece{ Location: Location{ Rank: 1, File: 5, }, PieceType: Pawn, Color: White, },
        Piece{ Location: Location{ Rank: 1, File: 6, }, PieceType: Pawn, Color: White, },
        Piece{ Location: Location{ Rank: 1, File: 7, }, PieceType: Pawn, Color: White, },
        Piece{ Location: Location{ Rank: 0, File: 0, }, PieceType: Rook, Color: White, },
        Piece{ Location: Location{ Rank: 0, File: 1, }, PieceType: Knight, Color: White, },
        Piece{ Location: Location{ Rank: 0, File: 2, }, PieceType: Bishop, Color: White, },
        Piece{ Location: Location{ Rank: 0, File: 3, }, PieceType: Queen, Color: White, },
        Piece{ Location: Location{ Rank: 0, File: 4, }, PieceType: King, Color: White, },
        Piece{ Location: Location{ Rank: 0, File: 5, }, PieceType: Bishop, Color: White, },
        Piece{ Location: Location{ Rank: 0, File: 6, }, PieceType: Knight, Color: White, },
        Piece{ Location: Location{ Rank: 0, File: 7, }, PieceType: Rook, Color: White, },
        // Black Pieces
        Piece{ Location: Location{ Rank: 6, File: 0, }, PieceType: Pawn, Color: Black, },
        Piece{ Location: Location{ Rank: 6, File: 1, }, PieceType: Pawn, Color: Black, },
        Piece{ Location: Location{ Rank: 6, File: 2, }, PieceType: Pawn, Color: Black, },
        Piece{ Location: Location{ Rank: 6, File: 3, }, PieceType: Pawn, Color: Black, },
        Piece{ Location: Location{ Rank: 6, File: 4, }, PieceType: Pawn, Color: Black, },
        Piece{ Location: Location{ Rank: 6, File: 5, }, PieceType: Pawn, Color: Black, },
        Piece{ Location: Location{ Rank: 6, File: 6, }, PieceType: Pawn, Color: Black, },
        Piece{ Location: Location{ Rank: 6, File: 7, }, PieceType: Pawn, Color: Black, },
        Piece{ Location: Location{ Rank: 7, File: 0, }, PieceType: Rook, Color: Black, },
        Piece{ Location: Location{ Rank: 7, File: 1, }, PieceType: Knight, Color: Black, },
        Piece{ Location: Location{ Rank: 7, File: 2, }, PieceType: Bishop, Color: Black, },
        Piece{ Location: Location{ Rank: 7, File: 3, }, PieceType: Queen, Color: Black, },
        Piece{ Location: Location{ Rank: 7, File: 4, }, PieceType: King, Color: Black, },
        Piece{ Location: Location{ Rank: 7, File: 5, }, PieceType: Bishop, Color: Black, },
        Piece{ Location: Location{ Rank: 7, File: 6, }, PieceType: Knight, Color: Black, },
        Piece{ Location: Location{ Rank: 7, File: 7, }, PieceType: Rook, Color: Black, },
    }
}
