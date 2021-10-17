package login

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/golang/glog"
	"github.com/marcusljx/ocbcctl/lib/vars"
)

const (
	sessionFile = ".session"
)

func getSessionFile(sessionKey string) string {
	return path.Join(vars.ConfigDir, sessionKey, sessionFile)
}

// lookupLocalSession retrieves the local session, if available
func lookupLocalSession(sessionKey string) (*Session, error) {
	fd, err := os.Open(getSessionFile(sessionKey))
	if err != nil {
		return nil, ErrNoLocalSessionFile
	}

	defer func() {
		if closeErr := fd.Close(); closeErr != nil {
			glog.Errorf("error while closing file descriptor: %v", closeErr)
		}
	}()

	data, err2 := ioutil.ReadAll(fd)
	if err2 != nil {
		glog.Errorf("error reading login context : %v")
		return nil, ErrUnrecognizedTokenContext
	}

	result := &Session{}

	if err3 := json.Unmarshal(data, result); err3 != nil {
		glog.Errorf("error reading login context : %v")
		return nil, ErrUnrecognizedTokenContext
	}
	return result, nil
}

func writeSessionToDisk(sessionKey string, session *Session) error {
	jsonData, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("error reading JSON: %v", err)
	}

	if writeErr := os.WriteFile(getSessionFile(sessionKey), jsonData, os.ModePerm); writeErr != nil {
		return fmt.Errorf("error writing session file: %v", writeErr)
	}

	return nil
}
