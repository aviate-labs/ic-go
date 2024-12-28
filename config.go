package ic

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Config struct {
	Args      map[string]string                `json:"args"`
	Candid    string                           `json:"candid"`
	Locations map[string]map[string]ConfigLink `json:"locations"`
}

func (c *Config) UnmarshalJSON(raw []byte) error {
	var config struct {
		Args      map[string]string
		Candid    string
		Locations map[string]map[string]string
	}
	if err := json.Unmarshal(raw, &config); err != nil {
		return err
	}
	c.Args = config.Args
	c.Candid = config.Candid
	c.Locations = make(map[string]map[string]ConfigLink)
	for dir, loc := range config.Locations {
		locDir, ok := c.Locations[dir]
		if !ok {
			locDir = make(map[string]ConfigLink)
			c.Locations[dir] = locDir
		}
		for k, l := range loc {
			parts := strings.SplitN(l, "://", 2)
			prefix, ok := c.Args[parts[0]]
			if !ok {
				return fmt.Errorf("invalid prefix")
			}
			url, err := url.JoinPath(prefix, parts[1])
			if err != nil {
				return err
			}
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("failed to get %s: %d", url, resp.StatusCode)
			}
			raw, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			locDir[k] = ConfigLink{
				URL:       url,
				CandidRaw: raw,
			}
		}
	}
	return nil
}

type ConfigLink struct {
	URL       string
	CandidRaw []byte
}

func (c ConfigLink) String() string {
	return c.URL
}
