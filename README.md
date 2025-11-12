# API REST con Arquitectura Hexagonal en Go

Template de proyecto para construir APIs REST escalables y mantenibles usando arquitectura hexagonal (puertos y adaptadores) en Go.

## 📋 Descripción

Este proyecto implementa una API REST siguiendo los principios de **arquitectura hexagonal** (también conocida como arquitectura de puertos y adaptadores). Esta arquitectura permite desacoplar la lógica de negocio de los detalles técnicos, facilitando testing, mantenibilidad y escalabilidad.

## 🏗️ Estructura del Proyecto

```
templateApiRestGo/
├── main.go                          # Punto de entrada del programa
├── go.mod                           # Definición del módulo Go
├── go.sum                           # Checksums de dependencias
├── README.md                        # Este archivo
│
└── internal/
    ├── application/                 # Capa de aplicación (lógica de negocio)
    │   └── user_service.go         # Servicios y puertos (interfaces)
    │
    ├── domain/                      # Capa de dominio (modelos)
    │   └── user_model.go           # Entidades del negocio
    │
    ├── infrastructure/              # Adaptadores técnicos (detalles)
    │   ├── db/                     # Adaptador de base de datos
    │   │   ├── conecction_mongo_client.go
    │   │   └── mongo_user_repository.go
    │   │
    │   ├── http/                   # Adaptador HTTP (API REST)
    │   │   ├── router.go           # Definición de rutas
    │   │   └── user_handler.go     # Handlers HTTP
    │   │
    │   └── repository/             # Implementaciones de repositorios
    │       └── mongo_user_repository.go
    │
    └── bootstrap/                   # Contenedor de dependencias
        └── container.go            # DI manual (inyección de dependencias)
```

## 🧠 Arquitectura Hexagonal

### Capas

#### 1. **Domain Layer** (`internal/domain/`)
- Contiene los **modelos de negocio** (entidades)
- No tiene dependencias externas
- Define la esencia del negocio
- Ejemplo: `User` model

```go
// internal/domain/user_model.go
type User struct {
    ID   string
    Name string
    Email string
}
```

#### 2. **Application Layer** (`internal/application/`)
- Contiene la **lógica de negocio** y los **puertos** (interfaces)
- No conoce cómo se implementan las interfaces (independiente de tecnología)
- Define qué operaciones puede hacer el sistema
- Ejemplo: `UserService` y `UserRepository` interface

```go
// internal/application/user_service.go
type UserRepository interface {
    Create(user *User) error
    GetByID(id string) (*User, error)
}

type UserService struct {
    repo UserRepository
}

func (s *UserService) CreateUser(user *User) error {
    // Lógica de negocio aquí
    return s.repo.Create(user)
}
```

#### 3. **Infrastructure Layer** (`internal/infrastructure/`)
- Contiene los **adaptadores** (implementaciones concretas)
- Código específico de tecnología (MongoDB, HTTP, etc.)
- Se conecta con sistemas externos
- Subcapas:
  - **`db/`**: Adaptador de base de datos
  - **`http/`**: Adaptador HTTP (REST API)
  - **`repository/`**: Implementaciones de repositorios

```go
// internal/infrastructure/repository/mongo_user_repository.go
type MongoUserRepository struct {
    client *db.MongoClient
}

func (r *MongoUserRepository) Create(user *domain.User) error {
    collection := r.client.Database.Collection("users")
    _, err := collection.InsertOne(context.Background(), user)
    return err
}
```

#### 4. **Bootstrap Layer** (`internal/bootstrap/`)
- Contenedor de dependencias (DI manual)
- Crea todas las instancias y conecta las dependencias
- Configura toda la aplicación

```go
// internal/bootstrap/container.go
func NewContainer() *Container {
    mongoClient, err := db.NewMongoClient(uri, "testdb")
    userRepo := repository.NewMongoUserRepository(mongoClient)
    userService := application.NewUserService(userRepo)
    userHandler := http.NewUserHandler(userService)
    
    return &Container{...}
}
```

## 🔌 Puertos y Adaptadores

### Puertos (Interfaces)
Los puertos son **contratos** que definen cómo interactúa el sistema:

```go
// Puerto de salida: cómo persiste datos
type UserRepository interface {
    Create(user *User) error
    GetByID(id string) (*User, error)
}

// Puerto de entrada: cómo recibe solicitudes HTTP
// Implementado por UserHandler
```

### Adaptadores (Implementaciones)
Los adaptadores **conectan** los puertos con tecnologías concretas:

- **Adaptador HTTP**: Convierte requests HTTP → dominio
- **Adaptador MongoDB**: Convierte dominio → documentos MongoDB
- **Adaptador PostgreSQL** (futuro): Alternativa a MongoDB

## 🚀 Inicio Rápido

### Requisitos
- Go 1.24 o superior
- MongoDB local o remoto
- Git

### Instalación

1. **Clonar o usar como template**
```bash
git clone <repo>
cd templateApiRestGo
```

2. **Instalar dependencias**
```bash
go mod download
```

3. **Configurar MongoDB**
Actualiza la URI de conexión en `internal/bootstrap/container.go`:
```go
mongoClient, err := db.NewMongoClient(
    "mongodb://localhost:27017",  // ← Cambia aquí
    "testdb",
)
```

4. **Ejecutar la aplicación**
```bash
go run main.go
```

La API estará disponible en `http://localhost:3001`

## 📡 Endpoints

### Crear Usuario
```bash
POST /users
Content-Type: application/json

{
  "id": "user123",
  "name": "John Doe",
  "email": "john@example.com"
}
```

**Respuesta (201)**
```json
{
  "message": "Usuario creado correctamente",
  "user": {
    "id": "user123",
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

## 🔄 Flujo de una Solicitud

```
HTTP Request
    ↓
UserHandler (Adaptador HTTP)
    ↓
UserService (Lógica de negocio)
    ↓
UserRepository (Puerto - interfaz)
    ↓
MongoUserRepository (Adaptador MongoDB)
    ↓
MongoDB Database
```

## ✅ Ventajas de esta Arquitectura

✨ **Independencia de Frameworks**: Cambiar Fiber por otro framework es fácil  
✨ **Independencia de Base de Datos**: Cambiar MongoDB por PostgreSQL sin tocar la lógica  
✨ **Testing Facilitado**: Mock fácil de interfaces  
✨ **Escalabilidad**: Fácil agregar nuevas funcionalidades  
✨ **Mantenibilidad**: Código organizado y desacoplado  
✨ **Reutilización**: Lógica no depende de detalles técnicos  

## 🧪 Testing

### Test unitario del servicio (sin BD)
```go
// test: Mock del repositorio
type MockUserRepository struct{}

func (m *MockUserRepository) Create(user *domain.User) error {
    return nil // Simular éxito
}

func TestCreateUser(t *testing.T) {
    mockRepo := &MockUserRepository{}
    service := application.NewUserService(mockRepo)
    
    user := &domain.User{ID: "1", Name: "Test"}
    err := service.CreateUser(user)
    
    if err != nil {
        t.Fail()
    }
}
```

### Test de integración (con BD real)
```bash
go test ./... -v
```

## 📦 Dependencias

```
github.com/gofiber/fiber/v2     → Framework HTTP (Adaptador)
go.mongodb.org/mongo-driver      → Driver MongoDB (Adaptador)
```

Ver `go.mod` para versiones exactas.

## 🔧 Extensiones Futuras

### Agregar nuevo adaptador (ej: PostgreSQL)

1. Crear `internal/infrastructure/postgres/postgres_user_repository.go`
```go
type PostgresUserRepository struct {
    db *sql.DB
}

func (r *PostgresUserRepository) Create(user *domain.User) error {
    // Implementación PostgreSQL
}
```

2. Actualizar `container.go` para usar PostgreSQL en lugar de MongoDB

**No requiere cambios en**: Domain, Application, HTTP layers ✨

### Agregar nuevo endpoint

1. Crear método en `UserHandler`
2. Registrar ruta en `router.go`
3. Listo (sin cambios en lógica)

## 📚 Referencias

- [Arquitectura Hexagonal - Alistair Cockburn](https://alistair.cockburn.us/hexagonal-architecture/)
- [Clean Architecture - Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Go Project Layout](https://github.com/golang-standards/project-layout)

## 🤝 Contribuir

Para agregar mejoras a este template:

1. Mantén la estructura de carpetas
2. Respeta las capas (no hagas imports circulares)
3. Documenta cambios importantes

## 📄 Licencia

Este template es de código abierto y disponible bajo licencia MIT.

---

**Creado como template para proyectos Go con arquitectura hexagonal**
