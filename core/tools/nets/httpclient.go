package nets

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpJson(method, url, req_data string) (map[string]interface{}, error) {
	//Url & Param
	client := &http.Client{}

	//Request
	req, err := http.NewRequest(method, url, strings.NewReader(req_data))
	if err != nil {
		return nil, err
	}
	//Content-Type
	req.Header.Set("Content-Type", "application/json")

	//Response
	rep, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(rep.Body)
	rep.Body.Close()
	if err != nil {
		return nil, err
	}

	//Json Data
	var jsons map[string]interface{}
	if err = json.Unmarshal(data, &jsons); err != nil {
		return nil, err
	}
	return jsons, nil
}

func HttpData(method, url, req_data string) (string, error) {
	//Url & Param
	client := &http.Client{}

	//Request
	req, err := http.NewRequest(method, url, strings.NewReader(req_data))
	if err != nil {
		return "", err
	}
	//Content-Type
	req.Header.Set("Content-Type", "application/json")

	//Response
	rep, err := client.Do(req)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(rep.Body)
	rep.Body.Close()
	if err != nil {
		return "", err
	}

	return string(data), nil
}
