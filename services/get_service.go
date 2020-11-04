package services
import (
	"os"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	Global "../global"
	Photo "../models/photo"
)

func GetPhotos(){

	petition, petitionError := http.NewRequest(
		"GET",
		os.Getenv("PHOTOS_API_URL"),
		nil,
	)

	if petitionError != nil {
		log.Fatalf("Error creating petition: %v", petitionError)
	}

	petition.Header.Add("Content-Type", "application/json")

	response, responseError := Global.ClientHttp.Do(petition)

	if responseError != nil {
		log.Fatalf("Error making petition")
	}
	
	defer response.Body.Close()

	log.Printf("get response code: %d", response.StatusCode)

	bodyResponse , bodyResponseError := ioutil.ReadAll(response.Body)

	if bodyResponseError != nil {
		log.Printf("Error on bodyResponse: %d" , bodyResponseError)
	}

	var responsePhotoJson []Photo.Photo
	json.Unmarshal(bodyResponse,&responsePhotoJson)
	log.Println("photos : ", responsePhotoJson)

}