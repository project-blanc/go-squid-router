package v2

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type Client struct {
	cli          *http.Client
	integratorID string
}

func NewClient(integratorID string) *Client {
	return &Client{
		integratorID: integratorID,
		cli:          &http.Client{}, // @TODO inject
	}
}

// Route calls the route endpoint returning route information and the call data object
func (c *Client) Route(params RouteRequestParameters) (*RouteResponse, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "https://v2.api.squidrouter.com/v2/route", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-integrator-id", c.integratorID)

	resp, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("request failed") // @TODO
	}

	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	rr := map[string]*RouteResponse{}
	err = json.Unmarshal(buff, &rr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response body")
	}

	return rr["route"], nil
}
