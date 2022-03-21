package main

import (
	"context"
	"flag"
	"github.com/custom-app/protobson"
	"github.com/custom-app/protobson/test"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/proto"
	"log"
	"reflect"
)

func main() {
	var url, db, collection string
	flag.StringVar(&url, "url", "", "url to connect")
	flag.StringVar(&db, "db", "", "database name")
	flag.StringVar(&collection, "collection", "", "collection to use")
	flag.Parse()
	if url == "" || db == "" || collection == "" {
		log.Panicln("cannot connect to database (empty url, db or collection)")
	}

	regBuilder := bson.NewRegistryBuilder()
	codec := protobson.NewCodec()
	msgType := reflect.TypeOf((*proto.Message)(nil)).Elem()
	registry := regBuilder.RegisterHookDecoder(msgType, codec).RegisterHookEncoder(msgType, codec).Build()

	opts := options.Client().SetRegistry(registry).ApplyURI(url)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Panicln(err)
	}
	defer client.Disconnect(context.Background())

	msg := &test.SimpleMessage{
		StringField: "test",
		Int64Field:  13,
		Int32Field:  37,
	}
	if _, err := client.Database(db).Collection(collection).InsertOne(context.Background(), msg); err != nil {
		log.Panicln(err)
	}

	res := client.Database(db).Collection(collection).FindOne(context.Background(),
		bson.D{{"string_field", bson.D{{"$eq", "test"}}}})
	if res.Err() != nil {
		log.Panicln(err)
	}
	// Note: you have to take pointer to proto generated struct, because mongo driver passes dereferenced value
	// to decode function, and if type of passed value will be test.SimpleMessage (not *test.SimpleMessage),
	// decoding func will fail (test.SimpleMessage doesn't implement proto.Message interface,
	// all required methods are pointer methods)
	if err := res.Decode(&msg); err != nil {
		log.Panicln(err)
	}

	var list []*test.SimpleMessage
	findRes, err := client.Database(db).Collection(collection).Find(context.Background(), bson.D{})
	if err != nil {
		log.Panicln(err)
	}
	if err := findRes.All(context.Background(), &list); err != nil {
		log.Panicln(err)
	}
}
