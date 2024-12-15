package rabbitpipe

import (
	"net/http"
	"strings"
	"testing"

	"github.com/birabittoh/myks"
)

const (
	testInstance = "inv.nadeko.net"
	testVideoID  = "qRY0m96ESZU"
	testQuery    = "youtube dl test video"
	testCaptions = "English (auto-generated)"
)

func TestFetchVideo(t *testing.T) {
	client := New(testInstance)

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
	client := New(testInstance)

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
	client := &Client{
		http:     http.DefaultClient,
		timeouts: &myks.KeyStore[error]{},
	}

	err := client.ensureInstance()
	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}

	if client.Instance == "" {
		t.Fatal("Expected instance to be set")
	}
}

func TestGetCachedVideos(t *testing.T) {
	client := New("")

	client.GetVideo(testVideoID)

	videos := client.GetCachedVideos()

	for id := range videos {
		if id == testVideoID {
			return
		}
	}

	t.Fatalf("Expected videos to contain %q", testVideoID)
}

func TestGetCaptions(t *testing.T) {
	client := New(testInstance)

	captions, err := client.GetCaptions(testVideoID, testCaptions)
	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}

	if captions == nil {
		t.Fatal("Expected captions, got nil")
	}
}
