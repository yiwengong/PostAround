package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"strconv"
)

/*
{
	lat:100,
	lon:10
}
映射
 */
type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

/*
{
	user:"jack",
	message:"hello",
	location:
	{
		lat:100,
		lon:10
	}
}
映射
 */
type Post struct{
	// `json:"user"` is for the json parsing of this User field. Otherwise, by default it's 'User'.
	User string `json:"user"`
	Message string `json:"message"`
	Location Location `json:"location"`
}

const (
	DISTANCE = "200km"
)

func main() {
	fmt.Println("started-service")
	// connect the handlerPost to url, but not exec
	// when exec, the value will pass into the function auto
	http.HandleFunc("/post", handlerPost)
	http.HandleFunc("/search", handlerSearch)
	log.Fatal(http.ListenAndServe(":8080", nil)) //if err, log the err in fatal
}

//* ONLY address, & look at the address
func handlerPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a request for post")
	decoder := json.NewDecoder(r.Body)
	var p Post
	if err := decoder.Decode(&p); err!= nil{
		panic(err)
		return
	}
	fmt.Fprintf(w, "Post received: %s\n", p.Message)
}

func handlerSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a search request")

	lat := r.URL.Query().Get("lat")
	//_ => err, written in _ to let go know we don't need that
	lt, _ := strconv.ParseFloat(lat, 64)

	lon := r.URL.Query().Get("lon")
	ln, _ := strconv.ParseFloat(lon, 64)

	ran := DISTANCE
	if val := r.URL.Query().Get("range"); val != "" {
		ran = val + "km"
	}
	fmt.Println("range is", ran)

	//first create a new object and then send the address of the post
	p := &Post{
		User:"1111",
		Message:"This is whatever",
		Location:Location{
			Lat:lt,
			Lon:ln,
		},
	}

	js,err := json.Marshal(p);
	if (err != nil) {
		return
	}
	w.Header().Set("Content-Type", "application-json")
	w.Write(js)
}

