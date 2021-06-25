package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"order-service/src/model"
	"os"
)

func GetUserById(id uint) model.User {
	url := fmt.Sprint(os.Getenv("USER_SERVICE_URL"), id)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	var parsedUser model.User
	json.Unmarshal([]byte(sb), &parsedUser)

	return parsedUser
}
