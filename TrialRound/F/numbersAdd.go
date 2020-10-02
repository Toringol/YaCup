package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	serverURL, serverPort, queryParamA, queryParamB := getInputValues()

	numbersArr, err := sendReqToServer(serverURL, serverPort, queryParamA, queryParamB)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(addArray(numbersArr))
}

func getInputValues() (string, string, string, string) {
	var serverURL, serverPort, queryParamA, queryParamB string

	fmt.Scan(&serverURL, &serverPort, &queryParamA, &queryParamB)

	return serverURL, serverPort, queryParamA, queryParamB
}

func sendReqToServer(serverURL, serverPort, queryParamA, queryParamB string) ([]int, error) {
	resultArr := make([]int, 0)

	req, err := http.NewRequest("GET", serverURL+":"+serverPort, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("a", queryParamA)
	q.Add("b", queryParamB)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&resultArr)
	if err != nil {
		return nil, err
	}

	return resultArr, nil
}

func addArray(numbersArr []int) int {
	resultAdd := 0

	for _, value := range numbersArr {
		resultAdd += value
	}

	return resultAdd
}
