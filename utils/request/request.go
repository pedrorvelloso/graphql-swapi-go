package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func DoGet(resource string) map[string]interface{} {
	resp, err := http.Get("https://swapi.co/api/" + resource)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	return result
}

func DoGetFullLink(link string) map[string]interface{} {
	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	return result
}
