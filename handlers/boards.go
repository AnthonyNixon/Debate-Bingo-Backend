package handlers

import (
	"fmt"
	"github.com/AnthonyNixon/Debate-Bingo-Backend/storage"
	"github.com/AnthonyNixon/Debate-Bingo-Backend/types"
	"github.com/AnthonyNixon/Debate-Bingo-Backend/utils"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

const NUM_CHARS_FOR_CODE = 5

type newBoard struct {
	Division string `json:'division'`
}

func NewBoard(c *gin.Context) {
	var newBoard newBoard
	binderr := c.BindJSON(&newBoard)
	if binderr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad JSON Input, could not bind."})
		return
	}

	rand.Seed(time.Now().UnixNano())

	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	code := make([]rune, NUM_CHARS_FOR_CODE)
	for i := range code {
		code[i] = letterRunes[rand.Intn(len(letterRunes))]
	}


	config, err := storage.GetConfigFile(newBoard.Division)
	if err != nil {
		c.JSON(err.StatusCode(), gin.H{"Error": err.Description()})
		return
	}

	var board types.Board
	board.Code = string(code)

	board.Bingo = false

	boxes := make([][]types.Box, 5)
	for i := range boxes {
		boxes[i] = make([]types.Box, 5)
	}

	fields := utils.ShuffleFields(config.Fields)
	freeSpace := types.Box{
		Content:     "Free Space!",
		Checked:     true,
		Coordinates: types.Coordinates{2,2},
	}

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			field := fields[(y*5) + x]
			boxes[x][y] = types.Box{Content: field.Content, Checked:false, Coordinates: types.Coordinates{x, y}}
		}
	}
	boxes[2][2] = freeSpace

	board.Boxes = boxes

	err = storage.SaveBoardFile(board)
	if err != nil {
		c.JSON(err.StatusCode(), gin.H{"Error": err.Description()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"board": board})
}

func GetBoard(c *gin.Context) {
	boardID, present := c.Params.Get("board_id")
	if !present {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No Board ID"})
		return
	}

	board, err := storage.GetBoardFromFile(boardID)
	if err != nil {
		fmt.Printf("%v", err)
		c.JSON(err.StatusCode(), gin.H{"Error": err.Description()})
		return
	}

	c.JSON(http.StatusOK, board)
}
