package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Passanger struct {
	Passanger_Id int    `json:"passanger_id"`
	First_Name   string `json:"first_name"`
	Last_Name    string `json:"last_name"`
	Mobile_No    string `json:"mobile_no"`
	Email        string `json:"email"`
}

func getAllPassanger() ([]Passanger, error) {
	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/passanger", nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				if res.StatusCode == http.StatusBadRequest {
					err = errors.New("ERROR: Bad Request")
					return nil, err
				}
				var allPassanger []Passanger
				json.Unmarshal(body, &allPassanger)
				return allPassanger, nil
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func getPassanger(id string) ([]Passanger, error) {
	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/passanger/"+id, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				if res.StatusCode == http.StatusBadRequest {
					err = errors.New("ERROR: Bad Request")
					return nil, err
				}
				var allPassanger []Passanger
				json.Unmarshal(body, &allPassanger)
				return allPassanger, nil
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func createPassanger(passanger Passanger) error {
	client := &http.Client{}
	postBody, _ := json.Marshal(passanger)
	resBody := bytes.NewBuffer(postBody)
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/passanger", resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == http.StatusAccepted {
				return nil
			} else {
				err = errors.New("ERROR: Bad Request")
				return err
			}
		} else {
			return err
		}
	} else {
		return err
	}
}

func updatePassanger(passanger Passanger) error {
	client := &http.Client{}
	postBody, _ := json.Marshal(passanger)
	resBody := bytes.NewBuffer(postBody)
	if req, err := http.NewRequest(http.MethodPut, fmt.Sprint("http://localhost:5000/api/v1/passanger/", passanger.Passanger_Id), resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == http.StatusAccepted {
				return nil
			} else {
				err = errors.New("ERROR: Bad Request")
				return err
			}
		} else {
			return err
		}
	} else {
		return err
	}
}
