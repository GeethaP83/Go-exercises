package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Name      string
	Id        string
	Job       string
	CreatedAt string
}

func main() {
	user := User{Name: "ABC", Job: "XYZ"}

	marshalledData, err := json.Marshal(user)
	fmt.Println("marshalledData : " + string(marshalledData))
	if err != nil {
		log.Printf("Marshal Failed: %s", err)
		return
	}

	resp, err := http.Post("https://reqres.in/api/users", "json", bytes.NewBuffer(marshalledData))
	if err != nil {
		log.Printf("Post Request Failed: %s", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}
	fmt.Println("Responsebody : " + string(body))

	post := User{}
	Unmarshalerr := json.Unmarshal(body, &post)
	if Unmarshalerr != nil {
		log.Printf("Unmarshal Failed: %s", Unmarshalerr)
		return
	}
	fmt.Println("id : " + post.Id)
	fmt.Println("createdAt : " + post.CreatedAt)
}
