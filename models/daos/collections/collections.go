package collections

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/tranHieuDev23/IdeTwo/utils/configs"
)

type DbCollections struct {
	SourceCodeCollection *mongo.Collection
	cancel               context.CancelFunc
}

func (collections *DbCollections) Cancel() {
	collections.cancel()
}

var instance *DbCollections = nil
var once sync.Once

func GetInstance() DbCollections {
	once.Do(func() {
		c := configs.GetInstance()
		client, err := mongo.NewClient(options.Client().ApplyURI(c.MongoDbUri))
		if err != nil {
			panic(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			panic(err)
		}

		db := client.Database(c.MongoDbDb)
		instance = &DbCollections{
			SourceCodeCollection: db.Collection("source_codes"),
			cancel:               cancel,
		}
	})
	return *instance
}
