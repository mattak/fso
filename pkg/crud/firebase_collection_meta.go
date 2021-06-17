package crud

import (
	"cloud.google.com/go/firestore"
)

type FirebaseCollectionMeta struct {
	ID          string `json:"id"`
	Path        string `json:"path"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
	ReadTime    string `json:"read_time"`
}

func CreateFirebaseCollectionMeta(snapshot *firestore.DocumentSnapshot) FirebaseCollectionMeta {
	return FirebaseCollectionMeta{
		ID:          snapshot.Ref.ID,
		Path:        snapshot.Ref.Path,
		CreatedTime: snapshot.CreateTime.String(),
		UpdatedTime: snapshot.UpdateTime.String(),
		ReadTime:    snapshot.ReadTime.String(),
	}
}
