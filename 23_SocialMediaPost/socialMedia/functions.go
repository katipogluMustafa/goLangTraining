package socialMedia

import (
	"time"
)

//go:generate stringer -type=MoodState

type MoodState int

// define all the possible mood states using an iota enumerator.
const (
	MoodStateNeutral MoodState = iota
	MoodStateHappy
	MoodStateSad
	MoodStateAngry
	MoodStateHopeful
	MoodStateThrilled
	MoodStateBored
	MoodStateShy
	MoodStateComical
	MoodStateOnCloudnine
)

type AuditableContent struct {
	TimeCreated  time.Time
	TimeModified time.Time
	CreatedBy    string
	ModifiedBy   string
}

// Social Media Post
type Post struct {
	AuditableContent // Embedded type
	Caption          string
	MessageBody      string
	URL              string
	ImageURI         string
	ThumbnailURI     string
	Keywords         []string
	Likers           []string
	AuthorMood       MoodState
}

// we use map to remember moodstates
var Moods map[string]MoodState

// init() responsible for initializing the mood state
func init() {
	Moods = map[string]MoodState{
		"neutral":   MoodStateNeutral,
		"happy":     MoodStateHappy,
		"sad":       MoodStateSad,
		"angry":     MoodStateAngry,
		"hopeful":   MoodStateHopeful,
		"thrilled":  MoodStateThrilled,
		"bored":     MoodStateBored,
		"shy":       MoodStateShy,
		"comical":   MoodStateComical,
		"cloudnine": MoodStateOnCloudnine,
	}
}

// constructor for new social media post

func NewPost(username string, mood MoodState, caption string, messageBody string, url string, imageURI string, thumbnailURI string, keywords []string) *Post {
	auditableContent := AuditableContent{CreatedBy: username, TimeCreated: time.Now()}
	return &Post{Caption: caption, MessageBody: messageBody, URL: url, ImageURI: imageURI, ThumbnailURI: thumbnailURI, AuthorMood: mood, Keywords: keywords, AuditableContent: auditableContent}
}
