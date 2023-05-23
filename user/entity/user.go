package entity

type User struct {
	ID string `json:"id"`
	//姓名
	Name string `json:"name"`
	//性别
	Gender bool `json:"gender"`
	//年龄
	Age string `json:"age"`
	//创建时间
	CreatedAt string `json:"createdAt"`
	//更新时间
	UpdatedAt string `json:"updatedAt"`
}
