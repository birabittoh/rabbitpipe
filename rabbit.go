package rabbitpipe

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

const (
	instancesEndpoint = "https://api.invidious.io/instances.json?sort_by=api,type"
	endpointFormat    = "https://%s/api/%s/%s"

	apiVersion = "v1"

	videosEndpoint = "videos"
	searchEndpoint = "search"
)

var (
	logger      = log.New(os.Stdin, "rabbitpipe: ", log.LstdFlags)
	expireRegex = regexp.MustCompile(`(?i)expire=(\d+)`)
)

func (c *Client) getFullEndpoint(endpoint string) string {
	return fmt.Sprintf(endpointFormat, c.Instance, apiVersion, endpoint)
}

func doRequest[T any](c *Client, endpoint string) (res *T, statusCode int) {
	err := c.ensureInstance()
	if err != nil {
		logger.Println("ERROR: Could not get a valid instance: ", err)
		return nil, http.StatusInternalServerError
	}

	resp, err := c.http.Get(c.getFullEndpoint(endpoint))
	if err != nil {
		logger.Println(err)
		return nil, http.StatusInternalServerError
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, http.StatusNotFound
	}

	if resp.StatusCode != http.StatusOK {
		logger.Println("Invidious gave the following status code: ", resp.StatusCode)
		return nil, resp.StatusCode
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Print(err)
		return nil, http.StatusInternalServerError
	}

	res = new(T)
	err = json.Unmarshal(body, res)
	if err != nil {
		logger.Println("ERROR:", err)
		return nil, http.StatusInternalServerError
	}

	return res, http.StatusOK
}

func (c *Client) fetchVideo(videoId string) (*Video, int) {
	endpoint := videosEndpoint + "/" + url.QueryEscape(videoId)
	return doRequest[Video](c, endpoint)
}

func (c *Client) fetchSearch(query string) (*[]SearchResult, int) {
	endpoint := searchEndpoint + "?q=" + url.QueryEscape(query)
	return doRequest[[]SearchResult](c, endpoint)
}

func (c *Client) ensureInstance() error {
	if c.Instance == "" {
		return c.NewInstance()
	}
	return nil
}
