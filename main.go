package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type requestAPI struct { //Al declarar las claves del objeto JSON, estas DEBEN empezar con mayuscula, sino, no serán consideradas
	Mensaje string
	Id      int
	Key     int
	Boolean bool
	String  string
	Date    time.Time
	File    string
	Null    string
	Content Content
}
type Content struct {
	Title  string
	Number int
	Desc   string
}

func dataHandler(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)   //Obtiene el body del request CODIFICADA
	var post []requestAPI                  //Crea la variable post de tipo Article (Data estructurada)
	json.Unmarshal([]byte(reqBody), &post) //Parsea/decodifica el body en la variable post
	json.NewEncoder(w).Encode(post)        //Ingresa post al output de la respuesta, siendo w la variable que almacena el output
	newData, err := json.Marshal(post)     //Marshal codifica nuevamente el objeto

	if err != nil { //control de errores
		fmt.Println(err) //imprime el error por consola
	} else {
		fmt.Println(time.Now())      //Imprime la respuesta correcta por consola
		fmt.Println(string(newData)) //Imprime la respuesta correcta por consola
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/articles", dataHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
