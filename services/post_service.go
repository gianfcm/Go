package services

import (
	"os"
	"log"
	"net/http"
	"bytes"
	Photo "../models/photo"
	Global "../global"
)


func PostToSlack(message []byte) {
	
	petition, petitionError := http.NewRequest(
		"POST",
		os.Getenv("SLACK_API_URL"),
		bytes.NewBuffer(message),
	)

	if petitionError != nil {
		log.Fatalf("Error creating petition: %v", petitionError)
	}

	petition.Header.Add("Content-Type", "application/json")

	response, responseError := Global.ClientHttp.Do(petition)
	if responseError != nil {
		log.Fatalf("Error making petition: %v", responseError)
	}

	defer response.Body.Close()

	log.Printf("post message to slack response code: %d", response.StatusCode)
}

func PostToServer(photo Photo.Photo){

	petition, petitionError := http.NewRequest(
		"POST",
		os.Getenv("PHOTOS_API_URL"),
		bytes.NewBuffer(photo.PhotoToJson()),
	)

	if petitionError != nil {
		log.Fatalf("Error creating petition: %v", petitionError)
	}

	petition.Header.Add("Content-Type", "application/json")

	response, responseError := Global.ClientHttp.Do(petition)
	if responseError != nil {
		log.Fatalf("Error making petition: %v", responseError)
	}
	
	defer response.Body.Close()

	log.Printf("post message to slack response code: %d", response.StatusCode)

}

