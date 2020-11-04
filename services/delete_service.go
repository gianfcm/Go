package services

import (
	"os"
	"log"
	"net/http"
	"bytes"
	Photo "../models/photo"
	Global "../global"
)

func DeletePhoto(photo Photo.Photo){

	petition, petitionError := http.NewRequest(
		"DELETE",
		os.Getenv("PHOTOS_API_URL") + string(photo.ID),
		bytes.NewBuffer(photo.PhotoToJson()),
	)

	if petitionError != nil {
		log.Fatalf("Error creating the petition: %v", petitionError)
	}

	petition.Header.Add("Content-Type", "application/json")

	response, responseError := Global.ClientHttp.Do(petition)

	if responseError != nil {
		log.Fatalf("Error making the petition: %v", responseError)
	}

	defer response.Body.Close()

	log.Printf("delete response code : %d", response.StatusCode)

}

