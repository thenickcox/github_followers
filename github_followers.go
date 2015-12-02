package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type User struct {
	Login string
}

func retrieveGithubAPIData() []byte {
	const githubURL string = "https://api.github.com/users/"

	url := githubURL + os.Args[1] + "/followers"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2)
	}
	return body
}

func main() {
	var users []User

	err := json.Unmarshal(retrieveGithubAPIData(), &users)
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		fmt.Println(user.Login)
	}
}
