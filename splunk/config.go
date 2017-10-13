package splunk

import (
	"log"

	"github.com/denniswebb/go-splunk/splunk"
)

type Config struct {
	URL                string
	Username           string
	Password           string
	InsecureSkipVerify bool
}

// Client() returns a new client for accessing Splunk.
func (c *Config) Client() (*splunk.Client, error) {
	client := splunk.New(c.URL, c.Username, c.Password, c.InsecureSkipVerify)
	log.Printf("[INFO] Splunk Client configured for: %s@%s", c.Username, c.URL)
	return client, nil
}
