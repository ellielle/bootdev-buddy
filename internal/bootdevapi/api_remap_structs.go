package bootdevapi

import "time"

type Archmage struct {
	StripeCustomerID        string    `json:"StripeCustomerID"`
	IsAdmin                 bool      `json:"IsAdmin"`
	ReferringUserUUID       any       `json:"ReferringUserUUID"`
	DiscordUserID           any       `json:"DiscordUserID"`
	DiscordUserHandle       any       `json:"DiscordUserHandle"`
	GithubAuthToken         any       `json:"GithubAuthToken"`
	SyncedGoogleID          any       `json:"SyncedGoogleID"`
	SyncedGithubID          any       `json:"SyncedGithubID"`
	ManualProSubExpiresAt   any       `json:"ManualProSubExpiresAt"`
	LifetimeProSubCreatedAt any       `json:"LifetimeProSubCreatedAt"`
	DiscordBotPrivate       bool      `json:"DiscordBotPrivate"`
	NotificationPreferences any       `json:"NotificationPreferences"`
	Email                   string    `json:"Email"`
	Currency                any       `json:"Currency"`
	Xp                      int       `json:"XP"`
	Level                   int       `json:"Level"`
	XPForLevel              int       `json:"XPForLevel"`
	XPTotalForLevel         int       `json:"XPTotalForLevel"`
	Gems                    int       `json:"Gems"`
	Armor                   int       `json:"Armor"`
	UUID                    string    `json:"UUID"`
	RecruitersCanContact    bool      `json:"RecruitersCanContact"`
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
	TwitterHandle           any       `json:"TwitterHandle"`
	LinkedinURL             string    `json:"LinkedinURL"`
	GithubHandle            string    `json:"GithubHandle"`
	WebsiteURL              any       `json:"WebsiteURL"`
	ProfileImageURL         string    `json:"ProfileImageURL"`
	ResumeURL               any       `json:"ResumeURL"`
	IsRecruiter             any       `json:"IsRecruiter"`
	ArchmageUnlockedAt      time.Time `json:"ArchmageUnlockedAt"`
}

type GlobalStats struct {
	LessonCompletions      int `json:"LessonCompletions"`
	XPEarned               int `json:"XPEarned"`
	BootsChats             int `json:"BootsChats"`
	SolutionsPeeked        int `json:"SolutionsPeeked"`
	DiscordMessagesSent    int `json:"DiscordMessagesSent"`
	RegisteredUsersAlltime int `json:"RegisteredUsersAlltime"`
}
