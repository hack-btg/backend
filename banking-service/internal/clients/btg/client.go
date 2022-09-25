package btg

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/hack-btg/backend/banking-service/internal/domains/models"
)

type Client struct {
	http http.Client
}

func NewClient() *Client {
	return &Client{
		http: http.Client{Timeout: time.Second * 30},
	}
}

func (c *Client) GetBanks() (models.Banks, error) {
	url := "https://data.directory.openbankingbrasil.org.br/participants"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("cache-control", "no-cache")
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	banks := models.Banks{}
	return banks, json.Unmarshal(body, &banks)
}
