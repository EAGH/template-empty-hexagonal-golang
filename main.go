package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"templateApiRestGo/internal/bootstrap"
	"templateApiRestGo/internal/infrastructure/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

// 🧠 main.go: punto de entrada del programa
// Aquí se inicializa todo y se arranca el servidor.

func main() {
	app := fiber.New()

	container := bootstrap.NewContainer()
	http.SetupRoutes(app, container.UserHandler)

	go func() {
		if err := app.Listen(":3001"); err != nil {
			fmt.Println("❌ Error iniciando servidor:", err)
		}
	}()

	fmt.Println("\n🚀 Servidor corriendo en el puerto 3001")

	// Esperar señal del sistema
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	// Cerrar recursos de forma ordenada
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := container.CloseAll(ctx); err != nil {
		fmt.Println("⚠️ Error cerrando dependencias:", err)
	}

	if err := app.Shutdown(); err != nil {
		fmt.Println("⚠️ Error cerrando servidor:", err)
	} else {
		fmt.Println("✅ Servidor cerrado correctamente")
	}
}
