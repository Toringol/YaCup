package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := `[
			8,  
  			6,  
  			-2,  
			2,  
			4,  
			17,  
			256,  
			1024,  
			-17,  
			-19  
		]`

		fmt.Fprintf(w, data)
	})
	http.ListenAndServe(":7777", nil)
}
