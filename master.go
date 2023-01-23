package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		fmt.Println("Can not unmarshal JSON")
	}
	return &lingoCloudResp

}

func Translate(content []string) {
	payload := &LingoCloudPayLoad{Source: content, TransType: "auto2zh", RequestID: "tootbot", Detect: true}
	result := LingoCloud(payload)
	for i := len(result.Target) - 1; i >= 0; i-- {
		fmt.Println(result.Target[i])
	}
}
