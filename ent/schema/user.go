package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Stars").Min(0).Max(999),
		field.String("Nickname"),
		// Remove all useless fields
		//field.Int("Other"),
		//field.String("AnotherField"),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		// No index for now
		//index.Fields("Nickname").Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
