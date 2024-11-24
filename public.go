package rabbitpipe

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/birabittoh/myks"
)

func New(defaultInstance string) *Client {
	client := &Client{
		http:     http.DefaultClient,
		timeouts: myks.New[error](time.Minute),
		videos:   myks.New[Video](time.Hour),
		search:   myks.New[[]SearchResult](2 * time.Hour),
		Instance: defaultInstance,
	}

	err := client.ensureInstance()
	if err != nil {
		logger.Print(err)
	}

	return client
}

func (c *Client) NewInstance() error {
	if c.Instance != "" {
		c.timeouts.Set(c.Instance, fmt.Errorf("generic error"), time.Hour)
	}

	resp, err := c.http.Get(instancesEndpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	var jsonArray [][]interface{}
	err = json.Unmarshal(body, &jsonArray)
	if err != nil {
		logger.Printf("ERROR: Could not unmarshal JSON response for instances.")
		return err
	}

	for i := range jsonArray {
		instance := jsonArray[i][0].(string)

		_, err := c.timeouts.Get(instance)
		if err != nil {
			c.Instance = instance
			logger.Print("Using new instance: ", c.Instance)
			return nil
		}
	}

	return fmt.Errorf("cannot find a valid instance")
}

func (c *Client) GetVideoNoCache(videoID string) (*Video, error) {
	c.videos.Delete(videoID)
	return c.GetVideo(videoID)
}

func (c *Client) GetVideo(videoID string) (*Video, error) {
	logger.Print("Video https://youtu.be/", videoID, " was requested.")

	video, err := c.videos.Get(videoID)
	if err == nil {
		logger.Printf("Found a valid cache entry.")
		return video, nil
	}

	video, httpErr := c.fetchVideo(videoID)

	switch httpErr {
	case http.StatusOK:
		logger.Printf("Retrieved by API.")
	case http.StatusNotFound:
		return nil, errors.New("video does not exist or can't be retrieved")
	default:
		err = c.NewInstance()
		if err != nil {
			logger.Print("ERROR: ", err)
			time.Sleep(10 * time.Second)
		}
		return c.GetVideo(videoID)
	}

	expiration := 5 * time.Hour
	if len(video.AdaptiveFormats) > 0 {
		expireString := expireRegex.FindStringSubmatch(video.AdaptiveFormats[0].URL)
		expireTimestamp, err := strconv.ParseInt(expireString[1], 10, 64)
		if err == nil {
			expiration = time.Until(time.Unix(expireTimestamp, 0))
		}
	}

	c.videos.Set(videoID, *video, expiration)

	return video, nil
}

func (c *Client) Search(query string) (*[]SearchResult, error) {
	q := strings.ToLower(strings.TrimSpace(query))

	results, err := c.search.Get(q)
	if err == nil {
		return results, nil
	}

	results, httpErr := c.fetchSearch(q)

	switch httpErr {
	case http.StatusOK:
		c.search.Set(q, *results, 3*time.Hour)
		return results, nil
	case http.StatusNotFound:
		return nil, errors.New("search did not return any valid videos")
	default:
		err := c.NewInstance()
		if err != nil {
			logger.Print("ERROR: ", err)
			time.Sleep(10 * time.Second)
		}
		return c.Search(query)
	}
}