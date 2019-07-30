package lib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)
func Post() {
	
	clienteHttp := &http.Client{}

	url := "https://jsonplaceholder.typicode.com/photos"

	type Foto struct {
			AlbumID      int    `json:"albumId"`
			ID           int    `json:"id"`
			Title        string `json:"title"`
			URL          string `json:"url"`
			ThumbnailURL string `json:"thumbnailUrl"`
	}

	type Slackmsg struct {
			Text string `json:"text"`
	}

	foto := Foto{
		1,
		5001,
		"foto title",
		"https://via.placeholder.com/600/e21cb",
		"https://via.placeholder.com/150/9e66a4",
	}
	
	slack_msg := Slackmsg{
		"Se creo una nueva foto",
	 }
	urlSlack := "https://hooks.slack.com/services/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"


	fotoComoJson, err := json.Marshal(foto)
	slackJson, err2 := json.Marshal(slack_msg)

	
	if err != nil {
		log.Fatalf("Error codificando foto como JSON: %v", err)
	}

	if err2 != nil {
		log.Fatalf("Error codificando foto como JSON: %v", err2)
	}

	peticion, err := http.NewRequest("POST", url, bytes.NewBuffer(fotoComoJson))
	peticiondos, err2 := http.NewRequest("POST",urlSlack,bytes.NewBuffer(slackJson))
	

	if err != nil {
		log.Fatalf("Error creando petición: %v", err)
	}
	
	if err2 != nil {
		log.Fatalf("Error creando petición: %v", err2)
	}

	
	peticion.Header.Add("Content-Type", "application/json")
	peticiondos.Header.Add("Content-Type", "application/json")
	
	respuesta, err := clienteHttp.Do(peticion)
	respuestaDos, err2 := clienteHttp.Do(peticiondos)
	
	if err != nil {
		log.Fatalf("Error haciendo petición: %v", err)
	}
	if err2 != nil {
		log.Fatalf("Error haciendo petición: %v", err2)
	}

	defer respuesta.Body.Close()
	defer respuestaDos.Body.Close()

	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		log.Fatalf("Error leyendo respuesta: %v", err)
	}
	cuerpoRespuestados, err2 := ioutil.ReadAll(respuestaDos.Body)
	if err != nil {
		log.Fatalf("Error leyendo respuesta: %v", err2)
	}

	respuestaString := string(cuerpoRespuesta)
	
	log.Printf("Código de respuesta: %d", respuesta.StatusCode)
	log.Printf("Código de respuesta DOS: %d", respuestaDos.StatusCode)
	contentType := respuesta.Header.Get("Content-Type")
	log.Printf("El tipo de contenido: '%s'", contentType)
	contentTypedos := respuestaDos.Header.Get("Content-Type")
	log.Printf("El tipo de contenido: '%s'", contentTypedos)
	log.Printf("Cuerpo de respuesta del servidor: '%s'", respuestaString)
	log.Printf("Cuerpo de respuesta del servidor: '%s'", slack_msg)
	log.Printf("Cuerpo dos: '%s'", cuerpoRespuestados)
}
