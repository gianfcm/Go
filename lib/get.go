package lib

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Get() {

	clienteHttp := &http.Client{}

	url := "https://jsonplaceholder.typicode.com/photos"
	
	peticion, err := http.NewRequest("GET", url, nil)
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
