package bootdevapi

import (
	"time"
)

type GlobalStats struct {
	LessonCompletions      int `json:"LessonCompletions"`
	XPEarned               int `json:"XPEarned"`
	BootsChats             int `json:"BootsChats"`
	SolutionsPeeked        int `json:"SolutionsPeeked"`
	DiscordMessagesSent    int `json:"DiscordMessagesSent"`
	RegisteredUsersAlltime int `json:"RegisteredUsersAlltime"`
}

type ActivityFeedUser struct {
	UUID string `json:"UUID"`
	User struct {
		Name   string `json:"Name"`
		Handle string `json:"Handle"`
		UUID   string `json:"UUID"`
		Level  int    `json:"Level"`
	} `json:"User"`
	CourseName string `json:"CourseName"`
}

// string pointers are used for nullable fields
type LeaderboardUser struct {
	DiscordUserHandle *string   `json:"DiscordUserHandle"`
	SyncedGoogleID    *string   `json:"SyncedGoogleID"`
	SyncedGithubID    *string   `json:"SyncedGithubID"`
	Currency          *string   `json:"Currency"`
	Xp                int       `json:"XP"`
	Level             int       `json:"Level"`
	XPForLevel        int       `json:"XPForLevel"`
	XPTotalForLevel   int       `json:"XPTotalForLevel"`
	Gems              int       `json:"Gems"`
	Armor             int       `json:"Armor"`
	UUID              string    `json:"UUID"`
	CreatedAt         time.Time `json:"CreatedAt"`
	UpdatedAt         time.Time `json:"UpdatedAt"`
	FirstName         string    `json:"FirstName"`
	LastName          string    `json:"LastName"`
	Handle            string    `json:"Handle"`
	Bio               string    `json:"Bio"`
	JobTitle          *string   `json:"JobTitle"`
	Location          string    `json:"Location"`
	City              *string   `json:"City"`
	Country           *string   `json:"Country"`
	GithubHandle      string    `json:"GithubHandle"`
	WebsiteURL        *string   `json:"WebsiteURL"`
	ProfileImageURL   string    `json:"ProfileImageURL"`
	XPEarned          int       `json:"XPEarned"`
}

type Archmage struct {
	LeaderboardUser
	ArchmageUnlockedAt time.Time `json:"ArchmageUnlockedAt"`
}

// Discord community members. It was a decent way to distinguish
// it with other `user` types in the code.
type Archon struct {
	LeaderboardUser
	Karma int `json:"Karma"`
}

type UserData struct {
	DiscordUserHandle       string    `json:"DiscordUserHandle"`
	SyncedGoogleID          any       `json:"SyncedGoogleID"`
	SyncedGithubID          int       `json:"SyncedGithubID"`
	ManualProSubExpiresAt   time.Time `json:"ManualProSubExpiresAt"`
	LifetimeProSubCreatedAt any       `json:"LifetimeProSubCreatedAt"`
	MembershipExpiresAt     time.Time `json:"MembershipExpiresAt"`
	Email                   string    `json:"Email"`
	Currency                string    `json:"Currency"`
	Xp                      int       `json:"XP"`
	Level                   int       `json:"Level"`
	XPForLevel              int       `json:"XPForLevel"`
	XPTotalForLevel         int       `json:"XPTotalForLevel"`
	Role                    string    `json:"Role"`
	Gems                    int       `json:"Gems"`
	Armor                   int       `json:"Armor"`
	CreatedAt               time.Time `json:"CreatedAt"`
	UpdatedAt               time.Time `json:"UpdatedAt"`
	FirstName               string    `json:"FirstName"`
	LastName                string    `json:"LastName"`
	Handle                  string    `json:"Handle"`
	Bio                     string    `json:"Bio"`
	JobTitle                any       `json:"JobTitle"`
	Location                string    `json:"Location"`
	City                    any       `json:"City"`
	Country                 any       `json:"Country"`
	GithubHandle            string    `json:"GithubHandle"`
	WebsiteURL              string    `json:"WebsiteURL"`
	ProfileImageURL         string    `json:"ProfileImageURL"`
	IsSubscribed            bool      `json:"IsSubscribed"`
	GithubSynced            bool      `json:"GithubSynced"`
}
