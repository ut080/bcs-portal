package airtable

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/internal/logging"
)

const airtableBaseURL = "https://api.airtable.com/v0"

type Client struct {
	baseId string
	token  string
	logger logging.Logger
}

func NewClient(baseId, token string, logger logging.Logger) Client {
	return Client{
		baseId: baseId,
		token:  token,
		logger: logger,
	}
}

func (c *Client) queryAirtable(method, table string, params map[string]string, body []byte) error {
	// Build URL
	url := fmt.Sprintf("%s/%s/%s", airtableBaseURL, c.baseId, table)

	// Set up request body
	var b io.Reader
	if body != nil {
		b = bytes.NewReader(body)
	}

	// Create a new request
	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return errors.WithStack(err)
	}

	// Set up authorization header
	auth := fmt.Sprintf("Bearer %s", c.token)
	req.Header.Add("Authorization", auth)

	// Set up URL params
	p := req.URL.Query()
	for k, v := range params {
		p.Add(k, v)
	}
	req.URL.RawQuery = p.Encode()

	// Send request
	client := &http.Client{}
	c.logger.Info().Msg("querying Airtable")
	resp, err := client.Do(req)
	if err != nil {
		c.logger.Error().Err(err).Msg("Airtable query failed")
		return errors.WithStack(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	return nil
}
