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
	XPBonus        float64 `json:"XPBonus"`
	XPTotal        int     `json:"XPTotal"`
	XPUser         int     `json:"XPUser"`
	NumActiveUsers int     `json:"NumActiveUsers"`
}

// Full course response from courses endpoint
type Course struct {
	UUID                         string     `json:"UUID"`
	ArchivedAt                   *time.Time `json:"ArchivedAt"`
	Slug                         string     `json:"Slug"`
	Title                        string     `json:"Title"`
	ShortDescription             string     `json:"ShortDescription"`
	Description                  string     `json:"Description"`
	ImageURL                     string     `json:"ImageURL"`
	ThumbnailURL                 string     `json:"ThumbnailURL"`
	Difficulty                   float64    `json:"Difficulty"`
	PrerequisiteCourseUUIDS      any        `json:"PrerequisiteCourseUUIDS"`
	EstimatedCompletionTimeHours int        `json:"EstimatedCompletionTimeHours"`
	TypeDescription              string     `json:"TypeDescription"`
	LastUpdated                  string     `json:"LastUpdated"`
	SlugAliases                  []string   `json:"SlugAliases"`
	AuthorUUIDs                  []string   `json:"AuthorUUIDs"`
	MaintainerUUIDs              []string   `json:"MaintainerUUIDs"`
	Alternatives                 struct {
	} `json:"Alternatives,omitempty"`
	Draft              bool `json:"Draft"`
	NumRequiredLessons int  `json:"NumRequiredLessons"`
	NumOptionalLessons int  `json:"NumOptionalLessons"`
	Chapters           []struct {
		UUID            string `json:"UUID"`
		Slug            string `json:"Slug"`
		Title           string `json:"Title"`
		Description     string `json:"Description"`
		RequiredLessons []struct {
			UUID           string `json:"UUID"`
			Slug           string `json:"Slug"`
			Type           string `json:"Type"`
			CourseUUID     string `json:"CourseUUID"`
			CourseTitle    string `json:"CourseTitle"`
			CourseImageURL string `json:"CourseImageURL"`
			ChapterUUID    string `json:"ChapterUUID"`
			IsFree         bool   `json:"IsFree"`
			LastMod        string `json:"LastMod"`
			CompletionType string `json:"CompletionType"`
			Title          string `json:"Title"`
		} `json:"RequiredLessons"`
		OptionalLessons    []any  `json:"OptionalLessons"`
		NumRequiredLessons int    `json:"NumRequiredLessons"`
		NumOptionalLessons int    `json:"NumOptionalLessons"`
		CourseUUID         string `json:"CourseUUID"`
	} `json:"Chapters"`
}

// User's course progress. Courses can be referenced by UUID
type CourseProgress struct {
	LastViewedCourseUUID string                    `json:"LastViewedCourseUUID"`
	Progress             map[string]ProgressDetail `json:"Progress"`
}

type ProgressDetail struct {
	NumDone              int    `json:"NumDone"`
	NumMax               int    `json:"NumMax"`
	LastViewedLessonUUID string `json:"LastViewedLessonUUID"`
}

// List of potions user has in effect / unused
type PotionList struct {
	ActiveXPPotions   []XPPotion `json:"ActiveXPPotions"`
	NumUnusedXPPotion int        `json:"NumUnusedXPPotion"`
}

// XP Potion information
type XPPotion struct {
	ID        string    `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	UserUUID  string    `json:"UserUUID"`
	ExpiresAt time.Time `json:"ExpiresAt"`
	UsedAt    time.Time `json:"UsedAt"`
	SoldAt    any       `json:"SoldAt"`
}

// Frozen flame data. Included since it seems it will
// be used in the future
type FrozenFlame struct {
	UUID      string    `json:"UUID"`
	UserUUID  string    `json:"UserUUID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UsedAt    any       `json:"UsedAt"`
	SoldAt    any       `json:"SoldAt"`
}
