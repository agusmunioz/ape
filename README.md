# Ape

Ape is a Go module that provides a handler that encapsulates Json management for Go handlers


```go
package main

import (
	"github.com/agusmunioz/ape"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
   Id     string   `json:"id"`
   Title  string   `json:"title"`
}

//GetArticles returns business objects in an ape response, no json encoding is need it.
func GetArticles(r *http.Request) ape.Response {
   articles:=  []Article{
                 {Id: "1234", Title: "An interesting article"},
		 {Id: "5678", Title: "Another interesting article"},
		} 
   return ape.NewOk(articles)
}

func main() {
   r := mux.NewRouter()
   r.Handle("/articles", ape.Handler(GetArticles)).Methods(http.MethodGet)
   log.Fatal(http.ListenAndServe(":8080", r))
}
```
