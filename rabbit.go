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

	videosEndpoint   = "videos"
	searchEndpoint   = "search"
	captionsEndpoint = "captions"
)

var (
	logger      = log.New(os.Stdin, "rabbitpipe: ", log.LstdFlags)
	expireRegex = regexp.MustCompile(`(?i)expire=(\d+)`)
)

func (c *Client) getFullEndpoint(endpoint string) string {
	c.ensureInstance()
	return fmt.Sprintf(endpointFormat, c.Instance, apiVersion, endpoint)
}

func doBareRequest(c *Client, url string) (body []byte, statusCode int) {
	resp, err := c.http.Get(url)
	if err != nil {
		logger.Println(err)
		return nil, http.StatusInternalServerError
	}
	defer resp.Body.Close()

	statusCode = resp.StatusCode

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		logger.Print(err)
	}
	return
}

func doBareInstanceRequest(c *Client, endpoint string) (body []byte, statusCode int) {
	err := c.ensureInstance()
	if err != nil {
		logger.Println("ERROR: Could not get a valid instance: ", err)
		return nil, http.StatusInternalServerError
	}

	return doBareRequest(c, c.getFullEndpoint(endpoint))
}

func doRequest[T any](c *Client, url string) (res *T, statusCode int) {
	body, statusCode := doBareRequest(c, url)
	if statusCode != http.StatusOK {
		logger.Println("Got the following status code: ", statusCode)
		return nil, statusCode
	}

	res = new(T)
	err := json.Unmarshal(body, res)
	if err != nil {
		logger.Println("ERROR:", err)
		return nil, http.StatusInternalServerError
	}

	return res, http.StatusOK
}

func doInstanceRequest[T any](c *Client, endpoint string) (*T, int) {
	err := c.ensureInstance()
	if err != nil {
		logger.Println("ERROR: Could not get a valid instance: ", err)
		return nil, http.StatusInternalServerError
	}

	return doRequest[T](c, c.getFullEndpoint(endpoint))
}

func (c *Client) fetchVideo(videoID string) (*Video, int) {
	endpoint := videosEndpoint + "/" + url.QueryEscape(videoID)
	return doInstanceRequest[Video](c, endpoint)
}

func (c *Client) fetchSearch(query string) (*[]SearchResult, int) {
	endpoint := searchEndpoint + "?q=" + url.QueryEscape(query)
	return doInstanceRequest[[]SearchResult](c, endpoint)
}

func (c *Client) fetchCaptions(videoID, language string) ([]byte, int) {
	endpoint := captionsEndpoint + "/" + videoID + "?label=" + url.QueryEscape(language)
	return doBareInstanceRequest(c, endpoint)
}

func (c *Client) ensureInstance() error {
	if c.Instance == "" {
		return c.NewInstance()
	}
	return nil
}
