package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func InsertCtoa(c *Client, ctoa Ctoa) (string, error) {

	resp := Ctoa{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(ctoa)

	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/api/ctoa", c.HostURL), buf)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Send the request
	r, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return resp.ID, nil

}

func DeleteCtoa(c *Client, id string) (string, error) {

	resp := Ctoa{}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s/api/ctoa/%s", c.HostURL, id), nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Send the request
	r, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return resp.ID, nil

}

func ReadCtoa(c *Client, id string) (Ctoa, error) {

	resp := []Ctoa{}

	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/api/ctoa", c.HostURL), nil)
	if err != nil {
		fmt.Println(err)
		return Ctoa{}, err
	}

	// Send the request
	r, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return Ctoa{}, err
	}

	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return Ctoa{}, err
	}

	// extract the ctoa
	for _, j := range resp {
		if j.ID == id {
			return j, nil
		}
	}

	return Ctoa{}, fmt.Errorf("cant find ctoa")

}

func UpdateCtoa(c *Client, ctoa Ctoa) (Ctoa, error) {

	resp := Ctoa{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(ctoa)

	req, err := http.NewRequest("PUT", fmt.Sprintf("http://%s/api/ctoa/%s", c.HostURL, ctoa.ID), buf)
	if err != nil {
		fmt.Println(err)
		return Ctoa{}, err
	}

	// Send the request
	r, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return Ctoa{}, err
	}

	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return Ctoa{}, err
	}

	return resp, nil

}
