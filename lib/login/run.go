package login

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os/user"
	"time"

	"github.com/golang/glog"

	"github.com/spf13/cobra"
)

func RunE(_ *cobra.Command, _ []string) error {
	sessionKey, err := getUserHash()
	if err != nil {
		glog.Errorf("getUserHash: %v", err)
		return fmt.Errorf("error while obtaining user hash")
	}

	session, err2 := GetSession(sessionKey)
	if err2 != nil {
		glog.Errorf("GetSession : %v", err2)
		return fmt.Errorf("error while logging in")
	}

	if err3 := writeSessionToDisk(sessionKey, session); err3 != nil {
		glog.Errorf("writeSessionToDisk : %v", err3)
		return fmt.Errorf("error while writing session")
	}

	fmt.Printf("session expiry: %v", session.Expiry.Sub(time.Now()))
	fmt.Printf("login successful")
	return nil
}

func getUserHash() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("unable to retrieve user details: %v", err)
	}
	data, err2 := json.Marshal(u)
	if err2 != nil {
		return "", fmt.Errorf("unrecognized user details format: %v", err)
	}

	hash := md5.Sum(data)
	return fmt.Sprintf("%x", hash), nil
}
