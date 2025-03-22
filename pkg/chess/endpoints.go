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


func HandleChess(c echo.Context) error {
    if global.Debug {
        fmt.Println("HandleChess():")
    }

    var board = Board{
        Pieces: getPieces(),
    }
    PopulateUiFields(board)

    return c.Render(http.StatusOK, "chess.html", ChessPage{
        Title: "Chess",
        Board: board,
    })
}

func HandleChessMove(c echo.Context) error {
    if global.Debug {
        fmt.Println("HandleChessMove():")
    }
    
    makeMove(c.FormValue("move"))

    return c.Render(http.StatusOK, "board", Board{
        Pieces: getPieces(),
    })
}

func PopulateUiFields(board Board) {
    if global.Debug {
        fmt.Println("PopulateUiFields(): ")
    }
    for i, piece := range board.Pieces {
        piece.ClassNames = GenerateClassNames(piece)
        board.Pieces[i] = piece
    }

    if global.Debug {
        fmt.Println(board.Pieces)
    }
}

func GenerateClassNames(piece Piece) string {
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
