
package lib
import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)
func Delete() {

	clienteHttp := &http.Client{}
	url := "https://jsonplaceholder.typicode.com/photos/"

	type Foto struct {
		    AlbumID      string    `json:"albumId"`
			ID           string    `json:"id"`
			Title        string `json:"title"`
			URL          string `json:"url"`
			ThumbnailURL string `json:"thumbnailUrl"`
	}

	foto := Foto{
		"1",
		"1",
		"accusamus beatae ad facilis cum similique qui sunt",
		"https://via.placeholder.com/600/92c952",
		"https://via.placeholder.com/150/92c952",
	}
	
	type Slackmsg struct {
		Text string `json:"text"`
	}

	slack_msg := Slackmsg{
		"Se elimino una nueva foto",
	 }

	fotoComoJson, err := json.Marshal(foto)
	slackJson, err2 := json.Marshal(slack_msg)
	urlSlack := "https://hooks.slack.com/services/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" 
	if err != nil {
		log.Fatalf("Error codificando como JSON: %v", err)
	}
	if err2 != nil {
		log.Fatalf("Error codificando como JSON: %v", err2)
	}

	peticion, err := http.NewRequest("DELETE", url+foto.ID, bytes.NewBuffer(fotoComoJson))
	if err != nil {
		log.Fatalf("Error creando petición: %v", err)
	}
	peticiondos, err2 := http.NewRequest("POST",urlSlack,bytes.NewBuffer(slackJson))
	if err2 != nil {
		log.Fatalf("Error creando petición: %v", err2)
	}

	peticiondos.Header.Add("Content-Type", "application/json")
	respuestados, err := clienteHttp.Do(peticiondos)
	if err != nil {
		log.Fatalf("Error haciendo petición: %v", err)
	}

	peticion.Header.Add("Content-Type", "application/json")
	respuesta, err := clienteHttp.Do(peticion)
	if err != nil {
		log.Fatalf("Error haciendo petición: %v", err)
	}

	defer respuesta.Body.Close()
	defer respuestados.Body.Close()

	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		log.Fatalf("Error leyendo respuesta: %v", err)
	}

	respuestaString := string(cuerpoRespuesta)
	log.Printf("Código de respuesta: %d", respuesta.StatusCode)
	log.Printf("Código de respuesta: %d", respuestados.StatusCode)
	contentType := respuesta.Header.Get("Content-Type")
	log.Printf("El tipo de contenido: '%s'", contentType)
	log.Printf("Cuerpo de respuesta del servidor: '%s'", respuestaString)
	log.Printf("url: '%s'", url+foto.ID)
}
