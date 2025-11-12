package domain

// 🧠 Dominio: contiene solo las entidades del negocio
// No sabe nada de base de datos ni HTTP

type User struct {
	ID    string `bson:"_id,omitempty" json:"id"`
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
}