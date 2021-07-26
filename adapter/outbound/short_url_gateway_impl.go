package outbound

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ShortURLGatewayImpl struct {
	host string
}

func NewShortURLGatewayImpl(host string) *ShortURLGatewayImpl {
	return &ShortURLGatewayImpl{host}
}

type shortURLCreateResponse struct {
	Shortcut string `json:"shortcut"`
}

func (s ShortURLGatewayImpl) CreateShortURL(longURL string) (string, error) {
	client := http.Client{}
	req := fmt.Sprintf(`{"url": %q}`, longURL)
	resp, err := client.Post(fmt.Sprintf("%s/create", s.host), "application/json", strings.NewReader(req))
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("short URL service error with status: %d", resp.StatusCode)
	}
	in, err := io.ReadAll(resp.Body)
	res := shortURLCreateResponse{}
	err = json.Unmarshal(in, &res)
	if err != nil {
		///
	}
	return res.Shortcut, nil
}
