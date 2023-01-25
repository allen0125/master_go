package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type LingoCloudPayLoad struct {
	Source    []string `json:"source"`
	TransType string   `json:"trans_type"`
	RequestID string   `json:"request_id"`
	Detect    bool     `json:"detect"`
}

type LingoCloudResp struct {
	Target     []string `json:"target"`
	RC         int8     `json:"rc"`
	Confidence float32  `json:"confidence"`
}

func LingoCloud(payload *LingoCloudPayLoad) *LingoCloudResp {
	url := "http://api.interpreter.caiyunai.com/v1/translator"
	b, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("x-authorization", os.Getenv("LINGO_CLOUD_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: time.Second * 15}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyIO, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var lingoCloudResp LingoCloudResp
	if err := json.Unmarshal(bodyIO, &lingoCloudResp); err != nil { // Parse []byte to go struct pointer
		log.Fatal("Can not unmarshal JSON")
	}
	return &lingoCloudResp

}

func Translate(content []string, transType string) []string {
	payload := &LingoCloudPayLoad{Source: content, TransType: transType, RequestID: "tootbot", Detect: true}
	return LingoCloud(payload).Target
}
