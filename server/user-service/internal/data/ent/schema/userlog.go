package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"time"
)

// UserLog holds the schema definition for the UserLog entity.
type UserLog struct {
	ent.Schema
}

// Fields of the UserLog.
func (UserLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id"),
		field.String("ip").Default(""),
		field.String("extra").Default(""),
		field.Time("create_time").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).UpdateDefault(time.Now),
	}
}

// Edges of the UserLog.
func (UserLog) Edges() []ent.Edge {
	return nil
}
