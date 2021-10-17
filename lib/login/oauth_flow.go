package login

import (
	"fmt"
	"net/url"
	"path"
	"time"

	"cloud.google.com/go/firestore"

	"github.com/golang/glog"

	"github.com/denisbrodbeck/machineid"

	"github.com/marcusljx/ocbcctl/lib/vars"
)

const (
	ocbcAPIPath   = "ocbcapi"
	authURLFormat = "https://api.ocbc.com/ocbcauthentication/api/oauth2/authorize?client_id=%s&redirect_uri=%s&scope=transactional"
)

func oauthFlow(sessionKey string) (*Session, error) {
	iter, err := getSessionDocListener(sessionKey)
	if err != nil {
		return nil, fmt.Errorf("error while starting session listener: %v", err)
	}
	defer iter.Stop()

	webviewURL, err2 := getOCBCAuthWebViewURL(sessionKey)
	if err2 != nil {
		return nil, fmt.Errorf("error computing OCBC Auth URL: %v", err2)
	}
	fmt.Printf("Visit this URL to login: %s", webviewURL)

	return waitForValidAccessSession(iter), nil
}

func waitForValidAccessSession(iter *firestore.DocumentSnapshotIterator) *Session {
	c := make(chan *Session)

	go func() {
		for {
			snapshot, snapshotErr := iter.Next()
			if snapshotErr != nil {
				glog.Errorf("snapshotIterator.Next() returned error: %v", snapshotErr)
				break
			}

			if !snapshot.Exists() {
				continue
			}

			session, parseErr := ParseSessionDoc(snapshot.Data())
			if parseErr != nil {
				glog.Errorf("login.ParseSessionDoc returned error: %v", parseErr)
				break
			}

			if session.Expiry.Before(time.Now()) {
				glog.Errorf("current token is expired")
				continue
			}

			c <- session
			close(c)
			break
		}
	}()

	return <-c
}

func getOCBCAuthWebViewURL(sessionKey string) (string, error) {
	callbackURL, err := getCallbackURL(sessionKey)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(authURLFormat, vars.OCBCAPIClientID, callbackURL), nil
}

func getCallbackURL(sessionKey string) (string, error) {
	if vars.CallbackHost == "" {
		return "", vars.ErrCallbackHostEnvNotSet
	}

	pid, hashErr := machineid.ProtectedID(sessionKey)
	if hashErr != nil {
		return "", fmt.Errorf("internal error: %v", hashErr)
	}

	u, err := url.Parse(path.Join(vars.CallbackHost, ocbcAPIPath, pid))
	if err != nil {
		return "", vars.ErrCallbackHostInvalidURL
	}

	return u.String(), nil
}
