package utility

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


func TaskCheckPremium() {

	client := http.Client{}

	req, err := http.NewRequest("PUT","http://localhost:8080/api/v1/checks/cron", nil)
	//get, err := http.Get("http://localhost:8080/api/v1/checks/cron")
	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()
	Body, _ := ioutil.ReadAll(res.Body)

	var response map[string]interface{}
	err = json.Unmarshal(Body, &response)

	fmt.Println(response)
	//return get

	//	fmt.Println("testing cron")
}
