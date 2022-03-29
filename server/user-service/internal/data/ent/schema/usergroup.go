package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"time"
)

// UserGroup holds the schema definition for the UserGroup entity.
type UserGroup struct {
	ent.Schema
}

// Fields of the UserGroup.
func (UserGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.String("name"),
		field.Int("access_level"),
		field.Time("create_time").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("update_time").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).UpdateDefault(time.Now),
	}
}

// Edges of the UserGroup
func (UserGroup) Edges() []ent.Edge {
	return nil
}
