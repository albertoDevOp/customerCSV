package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

//Callback to the CRM
func post(customer []byte) (string, error) {
	resp, err := http.Post("http://crm_integrator:5000/", "application/json", bytes.NewBuffer(customer))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
