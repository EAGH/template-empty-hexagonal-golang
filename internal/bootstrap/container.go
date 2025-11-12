package bootstrap

import (
	"context"
	"fmt"
	"templateApiRestGo/internal/application"
	"templateApiRestGo/internal/infrastructure/db"
	"templateApiRestGo/internal/infrastructure/http"
	"templateApiRestGo/internal/infrastructure/repository"
)

// 🧠 Bootstrap: donde se crean todas las instancias e inyectan dependencias.
// Actúa como un contenedor de dependencias simple (manual DI container).

type Container struct {
	MongoClient    *db.MongoClient
	UserRepository *repository.MongoUserRepository
	UserService    *application.UserService
	UserHandler    *http.UserHandler
}

// Crea todas las dependencias del proyecto y las conecta entre sí.
func NewContainer() *Container {
	// 1️⃣ Conexión a MongoDB
	mongoClient, err := db.NewMongoClient("mongodb+srv://eduardogomezsk8:xxxxxxxxx@dragonball.pygof.mongodb.net/?appName=dragonball", "testdb")
	if err != nil {
		panic(err)
	}

	// 2️⃣ Crear el repositorio (adaptador Mongo)
	userRepo := repository.NewMongoUserRepository(mongoClient)

	// 3️⃣ Crear el servicio e inyectar la interfaz
	userService := application.NewUserService(userRepo)

	// 4️⃣ Crear el handler HTTP e inyectar el servicio
	userHandler := http.NewUserHandler(userService)

	return &Container{
		MongoClient:    mongoClient,
		UserRepository: userRepo,
		UserService:    userService,
		UserHandler:    userHandler,
	}
}

// CloseAll cierra todas las conexiones del contenedor
func (c *Container) CloseAll(ctx context.Context) error {
	fmt.Println("\n🧹 Cerrando recursos...")

	if c.MongoClient != nil {
		if err := c.MongoClient.Close(ctx); err != nil {
			return fmt.Errorf("error al cerrar MongoDB: %w", err)
		}
		fmt.Println("✅ MongoDB cerrado correctamente")
	}

	return nil
}
