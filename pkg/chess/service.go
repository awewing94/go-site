package chess

import (
	"fmt"
	"go-site/pkg/global"
)

var CurrentGame = Game{}
// Chess game/board in the db will be composed of
//  - List of moves
//  - List of pieces (and their state)

/*
Make a move
Input:
    moveString string - String representation of move
Output:
    int - index of the Piece that was moved
*/
func makeMove(moveString string) int {
    if global.Debug {
        fmt.Println("service.makeMove(): Making move ", moveString)
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

    // hard code a move to test
    CurrentGame.Pieces[0].Location.Rank = 3
    // end 
    CurrentGame.Moves = append(CurrentGame.Moves, move)
    return 0;
}

func getPieces() []Piece {
    if global.Debug {
        fmt.Println("service.getPieces():")
    }

    if len(CurrentGame.Pieces) > 0{
        return CurrentGame.Pieces
    } else {
        initializeBoard()
        return CurrentGame.Pieces
    }
}

func initializeBoard() {
    CurrentGame.Moves = []Move{}
    CurrentGame.Pieces = []Piece{
        // White Pieces
        { Location: Location{ Rank: 1, File: 0, }, PieceType: Pawn, Color: White, },
        { Location: Location{ Rank: 1, File: 1, }, PieceType: Pawn, Color: White, },
        { Location: Location{ Rank: 1, File: 2, }, PieceType: Pawn, Color: White, },
        { Location: Location{ Rank: 1, File: 3, }, PieceType: Pawn, Color: White, },
        { Location: Location{ Rank: 1, File: 4, }, PieceType: Pawn, Color: White, },
        { Location: Location{ Rank: 1, File: 5, }, PieceType: Pawn, Color: White, },
        { Location: Location{ Rank: 1, File: 6, }, PieceType: Pawn, Color: White, },
        { Location: Location{ Rank: 1, File: 7, }, PieceType: Pawn, Color: White, },
        { Location: Location{ Rank: 0, File: 0, }, PieceType: Rook, Color: White, },
        { Location: Location{ Rank: 0, File: 1, }, PieceType: Knight, Color: White, },
        { Location: Location{ Rank: 0, File: 2, }, PieceType: Bishop, Color: White, },
        { Location: Location{ Rank: 0, File: 3, }, PieceType: Queen, Color: White, },
        { Location: Location{ Rank: 0, File: 4, }, PieceType: King, Color: White, },
        { Location: Location{ Rank: 0, File: 5, }, PieceType: Bishop, Color: White, },
        { Location: Location{ Rank: 0, File: 6, }, PieceType: Knight, Color: White, },
        { Location: Location{ Rank: 0, File: 7, }, PieceType: Rook, Color: White, },
        // Black Pieces
        { Location: Location{ Rank: 6, File: 0, }, PieceType: Pawn, Color: Black, },
        { Location: Location{ Rank: 6, File: 1, }, PieceType: Pawn, Color: Black, },
        { Location: Location{ Rank: 6, File: 2, }, PieceType: Pawn, Color: Black, },
        { Location: Location{ Rank: 6, File: 3, }, PieceType: Pawn, Color: Black, },
        { Location: Location{ Rank: 6, File: 4, }, PieceType: Pawn, Color: Black, },
        { Location: Location{ Rank: 6, File: 5, }, PieceType: Pawn, Color: Black, },
        { Location: Location{ Rank: 6, File: 6, }, PieceType: Pawn, Color: Black, },
        { Location: Location{ Rank: 6, File: 7, }, PieceType: Pawn, Color: Black, },
        { Location: Location{ Rank: 7, File: 0, }, PieceType: Rook, Color: Black, },
        { Location: Location{ Rank: 7, File: 1, }, PieceType: Knight, Color: Black, },
        { Location: Location{ Rank: 7, File: 2, }, PieceType: Bishop, Color: Black, },
        { Location: Location{ Rank: 7, File: 3, }, PieceType: Queen, Color: Black, },
        { Location: Location{ Rank: 7, File: 4, }, PieceType: King, Color: Black, },
        { Location: Location{ Rank: 7, File: 5, }, PieceType: Bishop, Color: Black, },
        { Location: Location{ Rank: 7, File: 6, }, PieceType: Knight, Color: Black, },
        { Location: Location{ Rank: 7, File: 7, }, PieceType: Rook, Color: Black, },
    }
}
