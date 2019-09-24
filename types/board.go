package types

type Board struct {
	OwnerID string `json:"ownerID"`
	Code string `json:"code"`
	Bingo bool `json:"bingo"`
	Boxes [][]Box `json:"boxes"`
}

func (board *Board) CheckBingo() {
	boxes := board.Boxes
	horizontal := false
	vertical := false
	diagonal := false

	// Check Horizontal
	for y := 0; y < 5; y++ {
		if boxes[0][y].Checked && boxes[1][y].Checked && boxes[2][y].Checked && boxes[3][y].Checked && boxes[4][y].Checked {
			horizontal = true
		}
	}

	// Check Vertical
	for x := 0; x < 5; x++ {
		if boxes[x][0].Checked && boxes[x][1].Checked && boxes[x][2].Checked && boxes[x][3].Checked && boxes[x][4].Checked {
			vertical = true
		}
	}

	// Check Diagonal
	if boxes[0][0].Checked && boxes[1][1].Checked && boxes[2][2].Checked && boxes[3][3].Checked && boxes[4][4].Checked {
		diagonal = true
	}

	if boxes[4][0].Checked && boxes[3][1].Checked && boxes[2][2].Checked && boxes[1][3].Checked && boxes[0][4].Checked {
		diagonal = true
	}

	if horizontal || vertical || diagonal {
		board.Bingo = true
	} else {
		board.Bingo = false
	}
}
