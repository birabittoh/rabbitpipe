package rabbitpipe

import (
	"net/http"
	"time"

	"github.com/birabittoh/myks"
)

type Format struct {
	Init            string `json:"init"`
	Index           string `json:"index"`
	Bitrate         string `json:"bitrate"`
	URL             string `json:"url"`
	Itag            string `json:"itag"`
	Type            string `json:"type"`
	Clen            string `json:"clen"`
	Lmt             string `json:"lmt"`
	ProjectionType  string `json:"projectionType"`
	Container       string `json:"container"`
	Encoding        string `json:"encoding"`
	AudioQuality    string `json:"audioQuality"`
	AudioSampleRate int    `json:"audioSampleRate"`
	AudioChannels   int    `json:"audioChannels"`
	Quality         string `json:"quality"`
	FPS             int    `json:"fps"`
	Size            string `json:"size"`
	Resolution      string `json:"resolution"`
	QualityLabel    string `json:"qualityLabel"`
}

type Captions struct {
	Label        string `json:"label"`
	LanguageCode string `json:"language_code"`
	URL          string `json:"url"`
}

type Thumbnail struct {
	Quality string `json:"quality"`
	URL     string `json:"url"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
}

type Storyboard struct {
	URL              string `json:"url"`
	TemplateURL      string `json:"templateUrl"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	Count            int    `json:"count"`
	Interval         int    `json:"interval"`
	StoryboardWidth  int    `json:"storyboardWidth"`
	StoryboardHeight int    `json:"storyboardHeight"`
	StoryboardCount  int    `json:"storyboardCount"`
}

type Client struct {
	http     *http.Client
	timeouts *myks.KeyStore[error]
	videos   *myks.KeyStore[Video]
	search   *myks.KeyStore[[]SearchResult]
	Instance string
}

type Video struct {
	Type              string       `json:"type"`
	Title             string       `json:"title"`
	VideoID           string       `json:"videoId"`
	VideoThumbnails   []Thumbnail  `json:"videoThumbnails"`
	Storyboards       []Storyboard `json:"storyboards"`
	Description       string       `json:"description"`
	DescriptionHTML   string       `json:"descriptionHtml"`
	Published         int64        `json:"published"`
	PublishedText     string       `json:"publishedText"`
	Keywords          []string     `json:"keywords"`
	ViewCount         int          `json:"viewCount"`
	LikeCount         int          `json:"likeCount"`
	DislikeCount      int          `json:"dislikeCount"`
	Paid              bool         `json:"paid"`
	Premium           bool         `json:"premium"`
	IsFamilyFriendly  bool         `json:"isFamilyFriendly"`
	AllowedRegions    []string     `json:"allowedRegions"`
	Genre             string       `json:"genre"`
	GenreURL          string       `json:"genreUrl"`
	Author            string       `json:"author"`
	AuthorID          string       `json:"authorId"`
	AuthorURL         string       `json:"authorUrl"`
	AuthorVerified    bool         `json:"authorVerified"`
	AuthorThumbnails  []Thumbnail  `json:"authorThumbnails"`
	SubCountText      string       `json:"subCountText"`
	LengthSeconds     int          `json:"lengthSeconds"`
	AllowRatings      bool         `json:"allowRatings"`
	Rating            int          `json:"rating"`
	IsListed          bool         `json:"isListed"`
	LiveNow           bool         `json:"liveNow"`
	IsPostLiveDVR     bool         `json:"isPostLiveDvr"`
	IsUpcoming        bool         `json:"isUpcoming"`
	DashURL           string       `json:"dashUrl"`
	AdaptiveFormats   []Format     `json:"adaptiveFormats"`
	FormatStreams     []Format     `json:"formatStreams"`
	Captions          []Captions   `json:"captions"`
	RecommendedVideos []Video      `json:"recommendedVideos"`
}

type SearchResult struct {
	Type            string      `json:"type"`
	Title           string      `json:"title"`
	VideoID         string      `json:"videoId"`
	Author          string      `json:"author"`
	AuthorID        string      `json:"authorId"`
	AuthorURL       string      `json:"authorUrl"`
	AuthorVerified  bool        `json:"authorVerified"`
	VideoThumbnails []Thumbnail `json:"videoThumbnails"`
	Description     string      `json:"description"`
	DescriptionHTML string      `json:"descriptionHtml"`
	ViewCount       int         `json:"viewCount"`
	ViewCountText   string      `json:"viewCountText"`
	Published       int         `json:"published"`
	PublishedText   string      `json:"publishedText"`
	LengthSeconds   int         `json:"lengthSeconds"`
	LiveNow         bool        `json:"liveNow"`
	Premium         bool        `json:"premium"`
	IsUpcoming      bool        `json:"isUpcoming"`
	IsNew           bool        `json:"isNew"`
	Is4k            bool        `json:"is4k"`
	Is8k            bool        `json:"is8k"`
	IsVR180         bool        `json:"isVr180"`
	IsVR360         bool        `json:"isVr360"`
	Is3D            bool        `json:"is3d"`
	HasCaptions     bool        `json:"hasCaptions"`
}

// Instances
type Server [2]interface{}

type Detail struct {
	Flag    string   `json:"flag"`
	Region  string   `json:"region"`
	Stats   *Stats   `json:"stats,omitempty"`
	CORS    *bool    `json:"cors,omitempty"`
	API     *bool    `json:"api,omitempty"`
	Type    string   `json:"type"`
	URI     string   `json:"uri"`
	Monitor *Monitor `json:"monitor,omitempty"`
}

type Stats struct {
	Version  string    `json:"version"`
	Software Software  `json:"software"`
	Usage    Usage     `json:"usage"`
	Metadata Metadata  `json:"metadata"`
	Playback *Playback `json:"playback,omitempty"`
}

type Software struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Branch  string `json:"branch"`
}

type Usage struct {
	Users Users `json:"users"`
}

type Users struct {
	Total          int `json:"total"`
	ActiveHalfYear int `json:"activeHalfyear"`
	ActiveMonth    int `json:"activeMonth"`
}

type Metadata struct {
	UpdatedAt              int64 `json:"updatedAt"`
	LastChannelRefreshedAt int64 `json:"lastChannelRefreshedAt"`
}

type Playback struct {
	TotalRequests      int     `json:"totalRequests"`
	SuccessfulRequests int     `json:"successfulRequests"`
	Ratio              float64 `json:"ratio"`
}

type Monitor struct {
	Token             string            `json:"token"`
	URL               string            `json:"url"`
	Alias             string            `json:"alias"`
	LastStatus        int               `json:"last_status"`
	Uptime            float64           `json:"uptime"`
	Down              bool              `json:"down"`
	DownSince         *time.Time        `json:"down_since,omitempty"`
	UpSince           time.Time         `json:"up_since"`
	Error             *string           `json:"error,omitempty"`
	Period            int               `json:"period"`
	ApdexT            float64           `json:"apdex_t"`
	StringMatch       string            `json:"string_match"`
	Enabled           bool              `json:"enabled"`
	Published         bool              `json:"published"`
	DisabledLocations []string          `json:"disabled_locations"`
	Recipients        []string          `json:"recipients"`
	LastCheckAt       time.Time         `json:"last_check_at"`
	NextCheckAt       time.Time         `json:"next_check_at"`
	CreatedAt         time.Time         `json:"created_at"`
	MuteUntil         *string           `json:"mute_until,omitempty"`
	FaviconURL        string            `json:"favicon_url"`
	CustomHeaders     map[string]string `json:"custom_headers"`
	HTTPVerb          string            `json:"http_verb"`
	HTTPBody          string            `json:"http_body"`
	SSL               SSL               `json:"ssl"`
}

type SSL struct {
	TestedAt  time.Time `json:"tested_at"`
	ExpiresAt time.Time `json:"expires_at"`
	Valid     bool      `json:"valid"`
	Error     *string   `json:"error,omitempty"`
}
