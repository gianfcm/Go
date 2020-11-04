package services
import (
	"log"
	"net/http"
	"os"
	"bytes"
	Global "../global"
	Photo "../models/photo"
)

func PutPhoto(photo Photo.Photo){
	petition, petitionError := http.NewRequest(
		"PUT",
		os.Getenv("PHOTOS_API_URL") + string(photo.ID),
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

	log.Printf("put photo on server code: %d", response.StatusCode)

}