package schema

import (
	"entgo.io/ent"
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
		field.String("title").
			Default(""),
		field.String("desc").
			Default(""),
		field.String("cover_image_url").
			Default(""),
		field.Text("content").
			Default(""),
		field.Int8("status").
			Default(1).
			Range(1, 9).
			Comment("1：正常使用 2：删除"),
		field.String("created_by").
			Default(""),
		field.String("updated_by").
			Default(""),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return nil
}
