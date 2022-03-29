package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"errors"
	"regexp"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int32("group_id").Default(0),
		field.String("username"),
		field.String("password"),
		field.String("mobile").Default(""),
		field.String("email").Match(regexp.MustCompile(
			"^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$")).
			Validate(func(s string) error {
				return errors.New("incorrect email format")
			}).Default(""),
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

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("id", UserLog.Type),
	}
}
