package rabbitpipe

import (
	"net/http"
	"strings"
	"testing"
)

const (
	testInstance = "inv.nadeko.net"
	testVideoID  = "BaW_jenozKc"
	testQuery    = "youtube-dl test video"
)

func TestFetchVideo(t *testing.T) {
	client := &Client{
		http:     http.DefaultClient,
		Instance: testInstance,
	}

	video, statusCode := client.fetchVideo(testVideoID)
	if statusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, statusCode)
	}

	if video == nil {
		t.Fatal("Expected video, got nil")
	}

	var allFormats = append(video.FormatStreams, video.AdaptiveFormats...)
	for _, format := range allFormats {
		if format.URL != "" {
			return
		}
	}

	t.Fatalf("Expected at least one format to have a URL")
}

func TestFetchSearch(t *testing.T) {
	client := &Client{
		http:     http.DefaultClient,
		Instance: testInstance,
	}

	results, statusCode := client.fetchSearch(testQuery)
	if statusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, statusCode)
	}

	if results == nil {
		t.Fatal("Expected search results, got nil")
	}

	if len(*results) == 0 {
		t.Fatal("Expected non-empty search results")
	}

	for _, result := range *results {
		if strings.Contains(strings.ToLower(result.Title), testQuery) {
			return
		}
	}

	t.Fatalf("Expected search results to contain %q", testQuery)
}

func TestEnsureInstance(t *testing.T) {
	client := New("")

	err := client.ensureInstance()
	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}

	if client.Instance == "" {
		t.Fatal("Expected instance to be set")
	}
}
