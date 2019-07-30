
package lib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Put() {

	clienteHttp := &http.Client{}
	
	url := "https://jsonplaceholder.typicode.com/photos/"


	type Foto struct {
		AlbumID      int    `json:"albumId"`
		ID           string    `json:"id"`
		Title        string `json:"title"`
		URL          string `json:"url"`
		ThumbnailURL string `json:"thumbnailUrl"`
}

	foto := Foto{
		 1,
		"1",
		"accusamus beatae ad facilis cum similique qui sunt",
		"https://via.placeholder.com/600/92c952",
		"https://via.placeholder.com/150/92c952",
	}

	fotoComoJson, err := json.Marshal(foto)
	if err != nil {
		log.Fatalf("Error codificando usuario como JSON: %v", err)
	}

	peticion, err := http.NewRequest("PUT", url+foto.ID, bytes.NewBuffer(fotoComoJson))
	if err != nil {
		log.Fatalf("Error creando petición: %v", err)
	}

	peticion.Header.Add("Content-Type", "application/json")
	respuesta, err := clienteHttp.Do(peticion)
	if err != nil {
		log.Fatalf("Error haciendo petición: %v", err)
	}

	defer respuesta.Body.Close()

	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		log.Fatalf("Error leyendo respuesta: %v", err)
	}

	respuestaString := string(cuerpoRespuesta)
	log.Printf("Código de respuesta: %d", respuesta.StatusCode)
	log.Printf("Cuerpo de respuesta del servidor: '%s'", respuestaString)
}
