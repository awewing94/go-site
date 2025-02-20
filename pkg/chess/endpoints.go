package chess

import (
	"fmt"
	"net/http"
    "go-site/pkg/global"

	"github.com/labstack/echo/v4"
)
type ChessPage struct {
    Title string

    Data string
}

type Board struct {
    Board []Piece
}

func HandleChess(c echo.Context) error {
    if global.Debug {
        fmt.Println("HandleChess():")
    }

    return c.Render(http.StatusOK, "chess.html", ChessPage{
        Title: "Chess",
        Data: "Here's some data having to do with chess",
    })
}

func HandleChessMove(c echo.Context) error {
    if global.Debug {
        fmt.Println("HandleChessMove():")
    }
    
    makeMove(c.FormValue("move"))
    var allPieces = getChessBoard()

    return c.Render(http.StatusOK, "board", Board{
        Board: allPieces,
    })
}

