package vars

import (
	"os"
)

var (
	CallbackHost          = os.Getenv("OCBCCTL_CALLBACK_HOST")
	FirestoreProjectID    string
	FirestoreCollectionID string
)
