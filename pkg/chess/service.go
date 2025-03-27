package chess

import (
	"fmt"
	"go-site/pkg/global"
	"reflect"
	"strings"
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
func makeMove(moveString string) (int, error) {
    if global.Debug {
        fmt.Println("service.makeMove(): Making move ", moveString)
    }

    var move, err = parseMoveMoveFmt(moveString)
    if err != nil {
        return -1, err
    }

    // Remove piece first to avoid index changes
    var endPieceIndex = findPieceIndexAtLocation(move.End)
    if endPieceIndex != -1 {
        CurrentGame.Pieces = removeAtIndex(CurrentGame.Pieces, endPieceIndex)
    }

    var startPieceIndex = findPieceIndexAtLocation(move.Start)
    if startPieceIndex == -1 {
        return -1, fmt.Errorf("No piece is located at the location %s", move.Start.ToString())
    }

    CurrentGame.Pieces[startPieceIndex].Location = move.End
    CurrentGame.Moves = append(CurrentGame.Moves, move)
    return startPieceIndex, nil;
}

func removeAtIndex(pieces []Piece, index int) []Piece {
    var newPieces = make([]Piece, len(pieces) - 1)
    var i, j = 0, 0
    fmt.Println(len(pieces))
    for ; i < len(pieces); i++ {
        fmt.Println("i: ", i, "; j: ", j)
        if i != index {
            newPieces[j] = pieces[i]
            j++
        }
    }

    return newPieces
}

func findPieceIndexAtLocation(location Location) int {
    for i := 0; i < len(CurrentGame.Pieces); i++ {
        var piece = CurrentGame.Pieces[i]
        var isEqual = reflect.DeepEqual(piece.Location, location)
        if isEqual {
            return i
        }
    }

    return -1
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

func parseMoveMoveFmt(moveString string) (Move, error) {
    var locations = strings.Split(moveString, "-")
    if (len(locations) != 2 || len(locations[0]) != 2 || len(locations[1]) != 2) {
        return Move{}, fmt.Errorf("Invalid move string supplied for move-move fmt. Format example: a1-a2.")
    }

    var move = Move{
        Start: parseMoveMoveLocation(locations[0]),
        End: parseMoveMoveLocation(locations[1]),
    }

    return move, nil
}

func parseMoveMoveLocation(locationString string) Location {
    return Location{
        Rank: int(locationString[1]) - '0' - 1,
        File: int(locationString[0]) - 'a',
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
