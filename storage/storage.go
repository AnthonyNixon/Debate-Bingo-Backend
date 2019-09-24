package storage

import (
	"cloud.google.com/go/storage"
	"encoding/json"
	"fmt"
	"github.com/AnthonyNixon/Debate-Bingo-Backend/customerrors"
	"github.com/AnthonyNixon/Debate-Bingo-Backend/types"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var client *storage.Client
var ctx context.Context
var bucketname string
var bucket *storage.BucketHandle

func Initialize() {
	log.Print("Initializing Storage")
	ctx = context.Background()
	myClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failec to initialize storage client: %s", err)
	}

	client = myClient

	bucketname = os.Getenv("DEBATEBINGO_BUCKET_NAME")
	if bucketname == "" {
		log.Fatal("DEBATEBINGO_BUCKET_NAME not set.")
	}

	bucket = client.Bucket(bucketname)

}

func GetConfigFile(division string) (types.Config, customerrors.Error) {
	objectName := fmt.Sprintf("config/%s.json", division)
	var config types.Config

	obj := bucket.Object(objectName)
	r, err := obj.NewReader(ctx)
	if err != nil {
		return config, customerrors.New(http.StatusNotFound, err.Error())

	}
	defer r.Close()

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return config, customerrors.New(http.StatusInternalServerError, err.Error())
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, customerrors.New(http.StatusInternalServerError, err.Error())
	}

	return config, nil
}

func SaveBoardFile(board types.Board) customerrors.Error {
	objectName := fmt.Sprintf("boards/%s.json", board.Code)
	wc := bucket.Object(objectName).NewWriter(ctx)
	wc.ContentType = "application/json"

	content, err := json.Marshal(board)
	if err != nil {
		return customerrors.New(http.StatusInternalServerError, err.Error())
	}

	if _, err := wc.Write(content); err != nil {
		return customerrors.New(http.StatusInternalServerError, err.Error())
	}

	if err := wc.Close(); err != nil {
		return customerrors.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func GetBoardFromFile(id string) (types.Board, customerrors.Error) {
	objectName := fmt.Sprintf("boards/%s.json", id)
	var board types.Board

	obj := bucket.Object(objectName)
	r, err := obj.NewReader(ctx)
	if err != nil {
		return board, customerrors.New(http.StatusNotFound,err.Error())
	}
	defer r.Close()

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return board, customerrors.New(http.StatusInternalServerError, err.Error())
	}

	err = json.Unmarshal(data, &board)
	if err != nil {
		return board, customerrors.New(http.StatusInternalServerError, err.Error())
	}

	return board, nil

}


