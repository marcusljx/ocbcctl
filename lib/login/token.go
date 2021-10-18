package login

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/glog"

	"cloud.google.com/go/firestore"
	"github.com/marcusljx/ocbcctl/lib/vars"
)

const (
	docTimeout = 90 * time.Second
)

func getSessionDocListener(sessionKey string) (*firestore.DocumentSnapshotIterator, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), docTimeout)
	glog.Infof("using project=%s", vars.DefaultConfig.FirebaseProjectID)
	firestoreClient, err := firestore.NewClient(ctx, vars.DefaultConfig.FirebaseProjectID)
	if err != nil {
		return nil, cancel, fmt.Errorf("error initializing firestore client: %v", err)
	}

	closeFunc := func() {
		cancel()
		_ = firestoreClient.Close()
	}

	glog.Infof("using collection=%s", vars.DefaultConfig.FirestoreCollectionID)
	collectionRef := firestoreClient.Collection(vars.DefaultConfig.FirestoreCollectionID)

	return collectionRef.Doc("userHash12345").Snapshots(ctx), closeFunc, nil
}
