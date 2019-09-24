package handlers

import (
	"github.com/AnthonyNixon/Debate-Bingo-Backend/storage"
	"github.com/AnthonyNixon/Debate-Bingo-Backend/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type toggleEvent struct {
	Coordinates types.Coordinates `json:"coordinates"`
	BoardID string `json:"board_id"`
}

func ToggleBox(c *gin.Context) {
	var toggleEvent toggleEvent
	binderr := c.BindJSON(&toggleEvent)
	if binderr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": binderr})
		return
	}

	x, y := toggleEvent.Coordinates.X, toggleEvent.Coordinates.Y

	board, err := storage.GetBoardFromFile(toggleEvent.BoardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}

	board.Boxes[x][y].Checked = !board.Boxes[x][y].Checked
	board.CheckBingo()

	err = storage.SaveBoardFile(board)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, board)
}