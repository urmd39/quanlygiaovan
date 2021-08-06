package control

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongo_uri_atlas = `mongodb+srv://urmd39:model.kn2412@cluster0.vovio.mongodb.net/test?
	authSource=admin&replicaSet=atlas-t75pz3-shard-0&readPreference=primary&
	appname=MongoDB%20Compass&ssl=true`

func Connected() (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_uri_atlas))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx
}
