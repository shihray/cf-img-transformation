package transform

import (
	"cloud.google.com/go/storage"
	"context"
	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	log "github.com/sirupsen/logrus"
)

var (
	ctx        = context.Background()
	storageCli *storage.Client
)

func init() {
	var err error
	gcs, err := storage.NewClient(ctx)
	if err != nil {
		panic(err)
	}
	storageCli = gcs
}

func Entry(ctx context.Context, e GCSEvent) (err error) {
	defer func() {
		panicErr := recover()
		if panicErr != nil {
			log.Errorf("recovery: %v", panicErr)
			return
		}
	}()

	log.Infof("File: %v\n", e.Name)
	log.Infof("SelfLink: %v\n", e.SelfLink)
	log.Infof("ContentType: %v\n", e.ContentType)

	err = ResizeAvatar(ctx, storageCli, e)
	if err != nil {
		return err
	}
	return
}
