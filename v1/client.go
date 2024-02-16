package v1

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type Client struct {
	cli *http.Client
}

func NewClient() *Client {
	return &Client{
		cli: &http.Client{}, // @TODO inject
	}
}

// Route calls the route endpoint returning route information and the call data object
func (c *Client) Route(params RouteRequestParameters) (*RouteResponse, error) {
	if params.Slippage <= 0.0 || params.Slippage > 99.99 {
		return nil, errors.New("slippage must be > 0 and < 99.99")
	}

	resp, err := c.cli.Get("https://api.0xsquid.com/v1/route?" + params.ToQuery().Encode())
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errorForResponse(resp.Body)
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

func errorForResponse(body io.Reader) error {
	buff, err := io.ReadAll(body)
	if err != nil {
		return errors.Wrap(err, "failed to read error response body")
	}

	er := map[string][]ErrorResponse{}
	err = json.Unmarshal(buff, &er)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal error response body")
	}

	return er["errors"][0].Err()
}
