package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("WELCOME TO THE SERVER")
    let myJson = req.body;      // your JSON
	
	res.status(200).send(myJson);

	http.HandleFunc("/get", hiiHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}

}

func hiiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, TI am SDE from tezminds")
}    let myJson = req.body;      // your JSON
	
res.status(200).send(myJson);

let myJson = req.body;      // your JSON
	
res.status(200).send(myJson);
