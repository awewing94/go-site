package chess

import (
	"fmt"
	"go-site/pkg/global"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)
type ChessPage struct {
    Title string
    Board Board
}

type Board struct {
    Pieces []Piece
}

var CurrentPage = ChessPage{
    Board: Board{},
    Title: "Chess Page",
}

func HandleChess(c echo.Context) error {
    if global.Debug {
        fmt.Println("HandleChess():")
    }

    CurrentPage.Board.Pieces = getPieces()
    updateAllUiFields(CurrentPage.Board)

    return c.Render(http.StatusOK, "chess.html", CurrentPage)
}

func HandleChessMove(c echo.Context) error {
    if global.Debug {
        fmt.Println("HandleChessMove():")
    }
    
    var changeIndex, err = makeMove(c.FormValue("move"))
    if err != nil {
        return c.Render(http.StatusBadRequest, err.Error(), CurrentPage)
    }

    CurrentPage.Board.Pieces = getPieces()

    fmt.Println("\tpieces piece: ", getPieces()[changeIndex])
    fmt.Println("\tchangeIndex: ", changeIndex, "\n\tpiece: ", CurrentPage.Board.Pieces[changeIndex])

    updateUiField(CurrentPage.Board, changeIndex)

    return c.Render(http.StatusOK, "board.html", CurrentPage)
}

func updateUiField(board Board, pieceIndex int) {
    if global.Debug {
        fmt.Println("UpdateUiField(): ", pieceIndex)
    }
    var piece = board.Pieces[pieceIndex]
    piece.ClassNames = generateClassNames(piece)
    board.Pieces[pieceIndex] = piece
}

func updateAllUiFields(board Board) {
    if global.Debug {
        fmt.Println("updateAllUiFields():")
    }
    for i, piece := range board.Pieces {
        piece.ClassNames = generateClassNames(piece)
        board.Pieces[i] = piece
    }
}

func generateClassNames(piece Piece) string {
    if global.Debug {
        fmt.Println("GenerateClassNames(): ")
    }

    var svgClass string = "";
    if piece.Color == White {
        svgClass += svgClass + "w_"
    } else {
        svgClass = svgClass + "b_"
    }

    svgClass = svgClass + strings.ToLower(string(piece.PieceType))

    var rankClass = "rank_" + strconv.Itoa(piece.Location.Rank + 1)
    var fileClass = "file_" + string(piece.Location.File + 'a')

    var result = svgClass + " " + rankClass + " " + fileClass
    if global.Debug {
        fmt.Println("\t" + result)
    }

    return result
}
