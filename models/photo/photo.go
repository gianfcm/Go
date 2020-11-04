package models

import (
	"encoding/json"
	"log"
)

type Photo struct {
  AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

func (photo Photo) PhotoToJson() []byte{
	photoJson, err := json.Marshal(photo)
	if err != nil {
		log.Fatalf("Error on photoToJson: %v", err)
	}
	return photoJson
}