package capwatch

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/internal/logging"
)

const capwatchURL = "https://www.capnhq.gov/CAP.CapWatchAPI.Web/api/cw"

type Client struct {
	orgID            string
	capwatchUsername string
	capwatchPassword []byte
	refresh          int
	logger           logging.Logger
}

func NewClient(orgID, capwatchUsername string, refresh int, logger logging.Logger) (c *Client) {
	nc := Client{
		orgID:            orgID,
		capwatchUsername: capwatchUsername,
		refresh:          refresh,
		logger:           logger,
	}

	c = &nc
	return c
}

func (c *Client) Fetch(filename string, refresh bool) (dump *Dump, err error) {
	if !refresh {
		if !c.WillRefreshCache(filename) {
			dump, err = c.readCache(filename)
			if err == nil {
				c.logger.Info().Msg("CAPWATCH cache loaded")
				return dump, nil
			}
			c.logger.Warn().Err(err).Msg("failed to find CAPWATCH cache, re-querying from CAPWATCH")
		}
		c.logger.Info().Msg("refreshing CAPWATCH cache")
	}

	dump, err = c.queryCapwatch()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = c.writeCache(dump, filename)
	if err != nil {
		c.logger.Error().Err(err).Msg("CAPWATCH cache write failed, proceeding anyway")
	}

	c.logger.Info().Msg("CAPWATCH cache loaded")
	return dump, nil
}

func (c *Client) WillRefreshCache(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return true
	}

	now := time.Now()
	diff := int((now.Sub(info.ModTime()).Hours()) / 24)

	if diff >= c.refresh {
		c.logger.Info().Int("age", diff).Int("refresh", c.refresh).Msg("CAPWATCH cache is old")
		return true
	}

	c.logger.Info().Int("age", diff).Int("refresh", c.refresh).Msg("CAPWATCH cache is new enough")
	return false
}

func (c *Client) SetCapwatchPassword(password []byte) {
	c.capwatchPassword = password
}

func (c *Client) queryCapwatch() (dump *Dump, err error) {
	// Create new request
	req, err := http.NewRequest("GET", capwatchURL, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Set up auth header
	rawCreds := append([]byte(c.capwatchUsername), ':')
	rawCreds = append(rawCreds, c.capwatchPassword...)
	creds := base64.StdEncoding.EncodeToString(rawCreds)
	auth := fmt.Sprintf("Basic %s", creds)
	req.Header.Add("Authorization", auth)

	// Set up URL params
	params := req.URL.Query()
	params.Add("ORGID", c.orgID)
	params.Add("unitOnly", "1")
	req.URL.RawQuery = params.Encode()

	// Send request
	client := &http.Client{}
	c.logger.Info().Msg("querying CAPWATCH")
	resp, err := client.Do(req)
	if err != nil {
		c.logger.Error().Err(err).Msg("capwatch query failed")
		return nil, errors.WithStack(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// CAPWATCH is kind of dumb, so anything that's not a 200 is an error
	if resp.StatusCode != 200 {
		err = errors.Errorf("CAPWATCH returned response %d %s: %s", resp.StatusCode, resp.Status, body)
		c.logger.Error().Err(err).Msg("capwatch query failed")
		return nil, err
	}

	dump = NewDump(body, time.Now())

	return dump, nil
}

func (c *Client) writeCache(dump *Dump, filename string) (err error) {
	c.logger.Info().Str("filename", filename).Msg("writing CAPWATCH cache")
	file, err := os.Create(filename)
	if err != nil {
		c.logger.Error().Err(err).Str("filename", filename).Msg("failed to create CAPWATCH cache")
		return errors.WithStack(err)
	}

	_, err = file.Write(dump.raw)
	if err != nil {
		c.logger.Error().Err(err).Str("filename", filename).Msg("failed to write CAPWATCH cache")
		return errors.WithStack(err)
	}

	return nil
}

func (c *Client) readCache(filename string) (dump *Dump, err error) {
	c.logger.Info().Str("filename", filename).Msg("loading CAPWATCH data from cache")
	file, err := os.Open(filename)
	if err != nil {
		c.logger.Error().Err(err).Str("filename", filename).Msg("failed to open CAPWATCH cache")
		return nil, errors.WithStack(err)
	}
	info, err := file.Stat()
	if err != nil {
		c.logger.Error().Err(err).Str("filename", filename).Msg("failed to stat CAPWATCH cache")
		return nil, errors.WithStack(err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		c.logger.Error().Err(err).Str("filename", filename).Msg("failed to read CAPWATCH cache")
		return nil, errors.WithStack(err)
	}

	dump = NewDump(content, info.ModTime())

	return dump, nil
}
