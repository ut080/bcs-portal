package capwatch

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const capwatchURL = "https://www.capnhq.gov/CAP.CapWatchAPI.Web/api/cw"

type Client struct {
	orgID            string
	capwatchUsername string
	capwatchPassword string
	refresh          int
}

func NewClient(orgID, capwatchUsername, capwatchPassword string, refresh int) Client {
	return Client{
		orgID:            orgID,
		capwatchUsername: capwatchUsername,
		capwatchPassword: capwatchPassword,
		refresh:          refresh,
	}
}

func (c Client) Fetch(filename string, refresh bool) (dump *Dump, err error) {
	if !refresh {
		if !c.refreshCache(filename) {
			dump, err = c.readCache(filename)
			if err == nil {
				log.Info().Msg("CAPWATCH cache loaded")
				return dump, nil
			}
			log.Warn().Err(err).Msg("failed to find CAPWATCH cache, re-querying from CAPWATCH")
		}
	}

	dump, err = c.queryCapwatch()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = c.writeCache(dump, filename)
	if err != nil {
		log.Error().Err(err).Msg("CAPWATCH cache write failed, proceeding anyway")
	}

	log.Info().Msg("CAPWATCH cache loaded")
	return dump, nil
}

func (c Client) refreshCache(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return true
	}

	now := time.Now()
	diff := int((now.Sub(info.ModTime()).Hours()) / 24)

	if diff >= c.refresh {
		log.Info().Int("age", diff).Int("refresh", c.refresh).Msg("CAPWATCH cache is old, refreshing")
		return true
	}

	log.Info().Int("age", diff).Int("refresh", c.refresh).Msg("CAPWATCH cache is new enough, continuing")
	return false
}

func (c Client) queryCapwatch() (dump *Dump, err error) {
	// Create new request
	req, err := http.NewRequest("GET", capwatchURL, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Set up auth header
	credString := fmt.Sprintf("%s:%s", c.capwatchUsername, c.capwatchPassword)
	creds := base64.StdEncoding.EncodeToString([]byte(credString))
	auth := fmt.Sprintf("Basic %s", creds)
	req.Header.Add("Authorization", auth)

	// Set up URL params
	params := req.URL.Query()
	params.Add("ORGID", c.orgID)
	params.Add("unitOnly", "1")
	req.URL.RawQuery = params.Encode()

	// Send request
	client := &http.Client{}
	log.Info().Msg("querying CAPWATCH")
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("capwatch query failed")
		return nil, errors.WithStack(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// CAPWATCH is kind of dumb, so anything that's not a 200 is an error
	if resp.StatusCode != 200 {
		err = errors.Errorf("CAPWATCH returned response %d %s: %s", resp.StatusCode, resp.Status, body)
		log.Error().Err(err).Msg("capwatch query failed")
		return nil, err
	}

	dump = NewDump(body, time.Now())

	return dump, nil
}

func (c Client) writeCache(dump *Dump, filename string) (err error) {
	log.Info().Str("filename", filename).Msg("writing CAPWATCH cache")
	file, err := os.Create(filename)
	if err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to create CAPWATCH cache")
		return errors.WithStack(err)
	}

	_, err = file.Write(dump.raw)
	if err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to write CAPWATCH cache")
		return errors.WithStack(err)
	}

	return nil
}

func (c Client) readCache(filename string) (dump *Dump, err error) {
	log.Info().Str("filename", filename).Msg("loading CAPWATCH data from cache")
	file, err := os.Open(filename)
	if err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to open CAPWATCH cache")
		return nil, errors.WithStack(err)
	}
	info, err := file.Stat()
	if err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to stat CAPWATCH cache")
		return nil, errors.WithStack(err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to read CAPWATCH cache")
		return nil, errors.WithStack(err)
	}

	dump = NewDump(content, info.ModTime())

	return dump, nil
}
