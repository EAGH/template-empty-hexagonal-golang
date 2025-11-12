package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 🧠 Infrastructure: acceso a la base de datos y conexiones externas

type MongoClient struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// Crea y retorna una conexión a MongoDB
func NewMongoClient(uri, dbName string) (*MongoClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	// Verificar la conexión con un Ping al servidor primario.
	if err := client.Database(dbName).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		// intentar desconectar para liberar recursos
		_ = client.Disconnect(ctx)
		return nil, err
	}

	return &MongoClient{
		Client:   client,
		Database: client.Database(dbName),
	}, nil
}

// Close desconecta el cliente de MongoDB y libera recursos.
func (m *MongoClient) Close(ctx context.Context) error {
	if m == nil || m.Client == nil {
		return nil
	}
	return m.Client.Disconnect(ctx)
}
