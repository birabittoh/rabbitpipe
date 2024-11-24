package rabbitpipe

import (
	"net/http"

	"github.com/birabittoh/myks"
)

type AdaptiveFormat struct {
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
}

type FormatStream struct {
	URL          string `json:"url"`
	Itag         string `json:"itag"`
	Type         string `json:"type"`
	Quality      string `json:"quality"`
	Bitrate      string `json:"bitrate"`
	FPS          int    `json:"fps"`
	Size         string `json:"size"`
	Resolution   string `json:"resolution"`
	QualityLabel string `json:"qualityLabel"`
	Container    string `json:"container"`
	Encoding     string `json:"encoding"`
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
	Type              string           `json:"type"`
	Title             string           `json:"title"`
	VideoID           string           `json:"videoId"`
	VideoThumbnails   []Thumbnail      `json:"videoThumbnails"`
	Storyboards       []Storyboard     `json:"storyboards"`
	Description       string           `json:"description"`
	DescriptionHTML   string           `json:"descriptionHtml"`
	Published         int64            `json:"published"`
	PublishedText     string           `json:"publishedText"`
	Keywords          []string         `json:"keywords"`
	ViewCount         int              `json:"viewCount"`
	LikeCount         int              `json:"likeCount"`
	DislikeCount      int              `json:"dislikeCount"`
	Paid              bool             `json:"paid"`
	Premium           bool             `json:"premium"`
	IsFamilyFriendly  bool             `json:"isFamilyFriendly"`
	AllowedRegions    []string         `json:"allowedRegions"`
	Genre             string           `json:"genre"`
	GenreURL          string           `json:"genreUrl"`
	Author            string           `json:"author"`
	AuthorID          string           `json:"authorId"`
	AuthorURL         string           `json:"authorUrl"`
	AuthorVerified    bool             `json:"authorVerified"`
	AuthorThumbnails  []Thumbnail      `json:"authorThumbnails"`
	SubCountText      string           `json:"subCountText"`
	LengthSeconds     int              `json:"lengthSeconds"`
	AllowRatings      bool             `json:"allowRatings"`
	Rating            int              `json:"rating"`
	IsListed          bool             `json:"isListed"`
	LiveNow           bool             `json:"liveNow"`
	IsPostLiveDVR     bool             `json:"isPostLiveDvr"`
	IsUpcoming        bool             `json:"isUpcoming"`
	DashURL           string           `json:"dashUrl"`
	AdaptiveFormats   []AdaptiveFormat `json:"adaptiveFormats"`
	FormatStreams     []FormatStream   `json:"formatStreams"`
	Captions          []Captions       `json:"captions"`
	RecommendedVideos []Video          `json:"recommendedVideos"`
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
