package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/segmentio/ksuid"
	"time"
)

// UserDemo holds the schema definition for the UserDemo entity.
type UserDemo struct {
	ent.Schema
}

// Fields of the UserDemo.
func (UserDemo) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(func() string {
			return ksuid.New().String()
		}),
		field.String("name"),                       //姓名
		field.String("age"),                        //年龄
		field.Bool("gender"),                       //性别
		field.Time("created_at").Default(time.Now), //创建时间
		field.Time("updated_at").Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the UserDemo.
func (UserDemo) Edges() []ent.Edge {
	return nil
}
