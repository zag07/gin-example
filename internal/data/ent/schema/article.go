package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"time"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("title"),
		field.String("desc"),
		field.String("cover_image_url"),
		field.Text("content"),
		field.Int("status").Default(1),
		field.String("created_by"),
		field.String("updated_by"),
		field.Time("created_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("deleted_at"),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return nil
}
