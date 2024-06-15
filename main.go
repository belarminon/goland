package main

import (
	"net/http"
	"fmt"
	"time"
)

func main()  { // thread 1
	// http.HandleFunc("/", HelloWorld)
	// http.ListenAndServe(":8888", nil)
	go contador(10) // thread 2
	go contador(10) // thread 3
	contador(10)
}

func HelloWorld(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello, World!"))
}

func contador(qtd int)  {
		for i := range qtd{
			fmt.Println(i)
			time.Sleep(time.Second)
		};
}