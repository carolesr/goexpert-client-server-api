package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type ExchangeRate struct {
	Value string `json:"value"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()

	exchangeRate, err := getExchangeRate(ctx)
	if err != nil {
		panic(err)
	}

	err = writeToFile(exchangeRate)
	if err != nil {
		panic(err)
	}
}

func getExchangeRate(ctx context.Context) (*ExchangeRate, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var exchangeRate ExchangeRate
	err = json.Unmarshal(response, &exchangeRate)
	if err != nil {
		return nil, err
	}

	return &exchangeRate, nil
}

func writeToFile(exchangeRate *ExchangeRate) error {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	data := fmt.Sprintf("Dólar: %s", exchangeRate.Value)
	_, err = file.Write([]byte(data))
	if err != nil {
		return err
	}

	return nil
}
