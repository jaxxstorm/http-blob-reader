package cloudblob

import (
	"context"
	"gocloud.dev/blob"
	// Import the blob driver packages we want to be able to open.
	_ "gocloud.dev/blob/azureblob"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
)

func Read(bucketUrl string, blobKey string) (*blob.Reader, error) {

	bucket, err := blob.OpenBucket(context.Background(), bucketUrl)

	if err != nil {
		return nil, err
	}
	defer bucket.Close()

	reader, err := bucket.NewReader(context.Background(), blobKey, nil)
	if err != nil {
		return nil, err
	}

	return reader, nil

}
