package main

import (
	f "fmt"
	s "net/http"
	l "log"
	e "encoding/json"
	m "math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type Anime struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Studio *Studio `json:"studio"`
}

type Studio struct{
	 Name string `json:"name"`
}


var animes []Anime

func getanimes(w s.ResponseWriter , r *s.Request){
	w.Header().Set("content-type", "appliation/json")
	e.NewEncoder(w).Encode(animes)
}

func createanime(w s.ResponseWriter, r *s.Request){
	w.Header().Set("content-type", "application/json")
	var anime Anime
	_=e.NewDecoder(r.Body).Decode(&anime)
	anime.ID = strconv.Itoa(m.Intn(1000000000))
	animes = append(animes , anime)
}

func deleteanime(w s.ResponseWriter , r *s.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for index , item := range animes {

		if item.ID == params["id"]{
			animes = append(animes[:index], animes[index+1:]...)
			e.NewEncoder(w).Encode(animes)
			break
		}
	}
}
func getanime (w s.ResponseWriter , r *s.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for _,item := range animes{
		if item.ID ==params ["id"]{
			e.NewEncoder(w).Encode(item)
			return
		}
	}

}

func updateanime(w s.ResponseWriter, r *s.Request) {
	//set json content
	w.Header().Set("content-type", "application/json")
	//params
	params:= mux.Vars(r)
	// range over animes and delete the anime 
	 for index , item := range animes{
		if item.ID == params["id"]{
			animes = append(animes[index:], animes[index+1:]...)
			var anime Anime 
			_ = e.NewDecoder(r.Body).Decode(&anime)
			anime.ID = params["id"]
			animes = append(animes, anime)
			e.NewEncoder(w).Encode(anime)
			return
		}
		
	 }
}



func main(){

	r:=mux.NewRouter()

	animes = append(animes, Anime{ID:"1", Isbn : "123667", Title: "I want to eat your pancreas" , Studio : &Studio{Name:"VOLN"}})
	animes = append(animes, Anime{ID:"2", Isbn : "114557", Title: "Your Name" , Studio : &Studio{Name:"CoMix Wave Films"}})
	r.HandleFunc("/anime",getanimes).Methods("GET")
	r.HandleFunc("/anime/{id}",getanime).Methods("GET")
	r.HandleFunc("/anime",createanime).Methods("POST")
	r.HandleFunc("/anime/{id}",updateanime).Methods("PUT")
	r.HandleFunc("/anime/{id}",deleteanime).Methods("DELETE")


	f.Printf("Staring server at port 8080")
	l.Fatal(s.ListenAndServe(":8080",r))

}