package repository

import "echoDemo/ent"

type Repository interface {
	User
}

type repository struct {
	//自动建表
	db *ent.Client
}

func NewRepository(db *ent.Client) *repository {
	return &repository{db: db}
}
