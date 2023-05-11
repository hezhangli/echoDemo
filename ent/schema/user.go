package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/segmentio/ksuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(func() string {
			return ksuid.New().String()
		}),
		field.String("hash"),                       //文件md5 hash，唯一
		field.String("name"),                       //文件名
		field.String("qiniu_file_path"),            //七牛云文件链接
		field.String("qiniu_tx_path"),              //七牛云文件处理成文本后的链接
		field.Strings("tags"),                      //文件标签，可用于筛选文件, 可以通过接口单独list
		field.String("user_id"),                    //用户id
		field.String("wx_id"),                      //用户企业微信id
		field.String("corp_id"),                    //用户企业id
		field.String("dept_id"),                    //用户部门id
		field.Time("created_at").Default(time.Now), //创建时间
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now), //更新时间
	}

}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
