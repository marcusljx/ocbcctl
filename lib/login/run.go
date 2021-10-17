package login

import (
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"

	"github.com/spf13/cobra"
)

var (
	sessionKey = os.Getenv("USER")
)

func RunE(_ *cobra.Command, _ []string) error {
	session, err := GetSession(sessionKey)
	if err != nil {
		glog.Errorf("GetSession : %v", err)
		return fmt.Errorf("error while logging in")
	}

	if err2 := writeSessionToDisk(sessionKey, session); err2 != nil {
		glog.Errorf("writeSessionToDisk : %v", err2)
		return fmt.Errorf("error while writing session")
	}

	fmt.Printf("session expiry: %v", session.Expiry.Sub(time.Now()))
	fmt.Printf("login successful")
	return nil
}
