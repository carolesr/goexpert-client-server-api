package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type USDBRLExchangeRate struct {
	USDBRL struct {
		Code       string `json:"code"`
		CodeIn     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		VarBID     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		BID        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func GetExchangeRateAPI() (*USDBRLExchangeRate, error) {
	client := http.Client{
		Timeout: time.Millisecond * 200,
	}

	req, err := client.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var exchangeRate USDBRLExchangeRate
	err = json.Unmarshal(res, &exchangeRate)
	if err != nil {
		return nil, err
	}

	return &exchangeRate, nil
}
