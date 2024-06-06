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

// LeaderboardUser is a struct to hold users from
// the leaderboard_stats endpoint
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

// Archmage holds users from the leaderboard_archmage endpoint
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

// UserData holds data about the user from the users endpoint
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

// CourseData holds information about all courses and the user's progress in them
type CourseData struct {
	UUID        string    `json:"UUID"`
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	ImageURL    string    `json:"ImageURL"`
	Category    string    `json:"Category"`
	Order       int       `json:"Order"`
	UnlockAtVal int       `json:"UnlockAtVal"`
	UnlockedAt  time.Time `json:"UnlockedAt"`
}

// Boss Battle information
type BossBattle struct {
	Event struct {
		Boss struct {
			UUID         string `json:"UUID"`
			Name         string `json:"Name"`
			Description  string `json:"Description"`
			LoreSlug     string `json:"LoreSlug"`
			ImageURL     string `json:"ImageURL"`
			ThumbnailURL string `json:"ThumbnailURL"`
			Rewards      []struct {
				UUID             string `json:"UUID"`
				Type             string `json:"Type"`
				UnlockedAt       int    `json:"UnlockedAt"`
				BossRewardDataXP struct {
					Amount int `json:"Amount"`
				} `json:"BossRewardDataXP"`
				BossRewardDataGems any `json:"BossRewardDataGems"`
				BossRewardChest    any `json:"BossRewardChest"`
			} `json:"Rewards"`
			HealthPoints int `json:"HealthPoints"`
		} `json:"Boss"`
		UUID               string     `json:"UUID"`
		CreatedAt          time.Time  `json:"CreatedAt"`
		UpdatedAt          time.Time  `json:"UpdatedAt"`
		StartsAt           time.Time  `json:"StartsAt"`
		ExpiresAt          time.Time  `json:"ExpiresAt"`
		AchievementsSentAt *time.Time `json:"AchievementsSentAt"`
		BossUUID           string     `json:"BossUUID"`
		AnnouncementSentAt time.Time  `json:"AnnouncementSentAt"`
		DefeatedAt         time.Time  `json:"DefeatedAt"`
	} `json:"Event"`
	XPBonus        int `json:"XPBonus"`
	XPTotal        int `json:"XPTotal"`
	XPUser         int `json:"XPUser"`
	NumActiveUsers int `json:"NumActiveUsers"`
}
