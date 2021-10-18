package login

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang/glog"
)

type Session struct {
	Expiry      time.Time `json:"expiry"`
	AccessToken string    `json:"access_token"`
}

// GetSession returns the access token of the current session.
// If no session exists, or the session has expired,
// the caller operation is paused and the user will be presented with
// the login flow again
func GetSession(sessionKey string) (*Session, error) {
	localSession, err := lookupLocalSession(sessionKey)
	if err == nil && localSession.Expiry.After(time.Now()) {
		fmt.Printf("Expiry: %v\n", localSession.Expiry.Sub(time.Now()))
		return localSession, nil
	}

	if err != nil {
		switch err {
		case ErrNoLocalSessionFile:
			fmt.Println("no session found, starting login flow")
		case ErrUnrecognizedTokenContext:
			fmt.Println("unable to recognize session, reattempting login")
		}
	}

	// at this point, do login flow
	session, loginErr := oauthFlow(sessionKey)
	if loginErr != nil {
		return nil, fmt.Errorf("error during login: %v", loginErr)
	}

	return session, nil
}

func ParseSessionDoc(data map[string]interface{}) (*Session, error) {
	b, err := json.Marshal(data)
	if err != nil {
		glog.Errorf("json.Marshal: %v", err)
		return nil, ErrUnrecognizedTokenContext
	}

	s := &Session{}
	if err := json.Unmarshal(b, s); err != nil {
		glog.Errorf("json.Unmarshal: %v", err)
		return nil, ErrUnrecognizedTokenContext
	}

	return s, nil
}
