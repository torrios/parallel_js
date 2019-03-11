package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/testit", processRequest)
	//Define Handler for html and js assets
	http.Handle("/", http.FileServer(http.Dir(".")))
	//http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//Responsible for receiving a test request, introducing a delay
//and returning.
func processRequest(w http.ResponseWriter, r *http.Request) {

	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Needs to be a POST", http.StatusBadRequest)
	// 	return
	// }

	// payload, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
	// 	return
	// }

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "calling ParseForm caused an error", http.StatusBadRequest)
	}

	delay, _ := strconv.Atoi(r.Form.Get("delay"))
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Fprintln(w, r.Form.Get("ajax"))
}
