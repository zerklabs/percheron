package percheron

import (
	"github.com/nu7hatch/gouuid"
	"time"
)

type UserInfo struct {
	Email   string
	Created time.Time
	ID      *uuid.UUID
}

type Bucket struct {
	Name    string
	Created time.Time
	Owner   *uuid.UUID
	ID      *uuid.UUID
}

// TODO(rch): add directory creation
func (user *UserInfo) NewBucket(name string) (Bucket, error) {
	id, err := uuid.NewV4()

	if err != nil {
		return Bucket{}, err
	}

	bucket := make(Bucket)

	bucket.Name = name
	bucket.Created = time.Now()
	bucket.Owner = user.ID
	bucket.ID = id

	return bucket, nil
}
