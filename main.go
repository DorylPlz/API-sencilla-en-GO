package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
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
	Content []Content
	Msg     string
}
type Content struct {
	NumberA int
	NumberB int
	Cond    string
	Value   int
}

func mainResponse(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)   //Obtiene el body del request CODIFICADA
	var post requestAPI                    //Crea la variable post de tipo Article (Data estructurada)
	json.Unmarshal([]byte(reqBody), &post) //Parsea/decodifica el body en la variable post

	post.Msg = getMsgReq() //Asigna el mensaje que se obtiene mediante request de otro endpoint (localhost:8080/getMsg)

	for k, _ := range post.Content { //Recorre post.Content
		match, _ := regexp.MatchString("[aeiou]", post.Content[k].Cond)

		if !match {
			post.Content[k].Value = post.Content[k].NumberA + post.Content[k].NumberB
		} else {
			post.Content[k].Value = post.Content[k].NumberA * post.Content[k].NumberB
		}
	}

	json.NewEncoder(w).Encode(post)    //Ingresa post al output de la respuesta, siendo w la variable que almacena el output
	newData, err := json.Marshal(post) //Marshal codifica nuevamente el objeto

	if err != nil { //control de errores
		fmt.Println(err) //imprime el error por consola
	} else { //Imprime la respuesta correcta por consola
		fmt.Println(string(newData)) //Imprime la respuesta correcta por consola
	}
}
func getMsg(w http.ResponseWriter, r *http.Request) { //Endpoint que retorna un string
	fmt.Fprintf(w, "Let's fucking GO")
}
func getMsgReq() string { //Request al endpoint getMsg
	resp, _ := http.Get("http://localhost:8080/getMsg")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

//Handlers
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(false)
	myRouter.HandleFunc("/main", mainResponse).Methods("POST")
	myRouter.HandleFunc("/getMsg", getMsg).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
func main() {
	handleRequests()
}
