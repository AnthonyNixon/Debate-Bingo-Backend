package utils

import (
	"github.com/AnthonyNixon/Debate-Bingo-Backend/types"
	"math/rand"
	"time"
)

func ShuffleFields(vals []types.Field) []types.Field {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]types.Field, len(vals))
	n := len(vals)
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(vals))
		ret[i] = vals[randIndex]
		vals = append(vals[:randIndex], vals[randIndex+1:]...)
	}
	return ret
}

