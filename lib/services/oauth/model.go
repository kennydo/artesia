package oauth

import "time"

// Client represents an OAuth client's model in the db
type Client struct {
	ID          int    `db:"id"`
	ExternalID  string `db:"externalId"`
	Secret      string `db:"secret"`
	RedirectURI string `db:"redirect_uri"`
	UserData    string `db:"user_data"`
}

// Authorization represents an OAuth client's authorization (who they are)
type Authorization struct {
	ID          int       `db:"id"`
	ClientID    int       `db:"clientId"`
	Code        string    `db:"code"`
	Expiration  int32     `db:"expiration"`
	Scope       string    `db:"scope"`
	RedirectURI string    `db:"redirect_uri"`
	StateData   string    `db:"state_data"`
	CreatedAt   time.Time `db:"created_at"`
	UserData    string    `db:"user_data"`
}

// Access is a DB representation of a an OAuth Client's access rights
type Access struct {
	ID           int    `db:"id"`
	ClientID     int    `db:"client_id"`
	AuthorizeID  int    `db:"authorize_id"`
	AccessToken  string `db:"access_token"`
	Refreshtoken string `db:"refresh_token"`
	Expiration   int32  `db:"expiration"`
	Scope        string `db:"scope"`
	RedirectURI  string `db:"redirect_uri"`
	CreatedAt    string `db:"created_at"`
}
