package login

import (
	"context"
	"fmt"

	"github.com/denisbrodbeck/machineid"

	"github.com/marcusljx/ocbcctl/lib/vars"

	"cloud.google.com/go/firestore"
)

func getSessionDocListener(sessionKey string) (*firestore.DocumentSnapshotIterator, error) {
	ctx := context.Background()
	firestoreClient, err := firestore.NewClient(ctx, vars.FirestoreProjectID)
	if err != nil {
		return nil, fmt.Errorf("error initializing firestore client: %v", err)
	}

	collectionRef := firestoreClient.Collection(vars.FirestoreCollectionID)

	pid, hashErr := machineid.ProtectedID(sessionKey)
	if hashErr != nil {
		return nil, fmt.Errorf("internal error: %v", hashErr)
	}

	return collectionRef.Doc(pid).Snapshots(ctx), nil
}
