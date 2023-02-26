package model

import "time"

type GithubToken struct {
	AccessToken string `json:"accessToken"`
	Scope       string `json:"scope"`
	TokenType   string `json:"tokenType"`
}

type GithubUser struct {
	Login             string    `json:"login"`
	ID                int       `json:"id"`
	NodeID            string    `json:"nodeID"`
	AvatarUrl         string    `json:"avatarUrl"`
	GravatarID        string    `json:"gravatarID"`
	URL               string    `json:"URL"`
	HtmlUrl           string    `json:"htmlUrl"`
	FollowersURL      string    `json:"followersURL"`
	FollowingsURL     string    `json:"followingsUrl"`
	GistsUrl          string    `json:"gistsURL"`
	StarredURL        string    `json:"starredURL"`
	SubscriptionURL   string    `json:"subscriptionURL"`
	OrganizationsURL  string    `json:"organizationsURL"`
	ReposURL          string    `json:"reposURL"`
	EventsURL         string    `json:"eventsURL"`
	ReceivedEventsURL string    `json:"receivedEventsURL"`
	Type              string    `json:"type"`
	SiteAdmin         bool      `json:"siteAdmin"`
	Name              string    `json:"name"`
	Blog              string    `json:"blog"`
	Location          string    `json:"location"`
	Email             *string   `json:"email"`
	Hireable          bool      `json:"hireable"`
	Bio               string    `json:"bio"`
	PublicRepos       int       `json:"publicRepos"`
	PublicGists       int       `json:"publicGists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	Token             string    `json:"-"`
}
