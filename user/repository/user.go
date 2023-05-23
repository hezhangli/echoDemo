package repository

import (
	"context"
)

type User interface {

	// CreateUser 创建用户信息
	CreateUser(ctx context.Context, name, age string, gender bool) error
	// UpdateUser 更新用户信息
	/*UpdateUser(ctx context.Context, user entity.User) error
	// ListUser 查询用户信息
	ListUser(ctx context.Context) ([]entity.User, error)
	// SearchUserByName 根据用户name查询用户信息
	SearchUserByName(ctx context.Context, name string) (entity.User, error)
	// DeleteUser 删除用户信息
	DeleteUser(ctx context.Context, name string) error*/
}

func (r *repository) CreateUser(ctx context.Context, name, age string, gender bool) error {

	_, err := r.db.UserDemo.Create().
		SetName(name).
		SetAge(age).
		SetGender(gender).
		Save(ctx)

	if err != nil {
		return err
	}
	return nil
}

/*func (r *repository) UpdateUser(ctx context.Context, userId string, user entity.User) error {

	return nil
}

func (r *repository) ListUser(ctx context.Context) ([]entity.User, error) {

	users, err := r.db.User.Query().All(ctx)

	if err != nil {
		return nil, err
	}

	eusers := make([]entity.User, 0, len(users))

	for _, v := range users {

		eusers = append(eusers, entity.User{
			CreatedAt: v.CreatedAt.Format("2006-01-02 13:04:05"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 13:04:05"),
		})
	}
	return eusers, nil
}

func (r *repository) SearchUserByName(ctx context.Context, name string) ([]entity.User, error) {

	users, err := r.db.User.Query().Where(user.NameEQ(name)).All(ctx)

	if err != nil {
		return nil, err
	}

	eusers := make([]entity.User, 0, len(users))

	for _, v := range users {

		eusers = append(eusers, entity.User{
			CreatedAt: v.CreatedAt.Format("2006-01-02 13:04:05"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 13:04:05"),
		})
	}

	return eusers, nil
}

func (r *repository) DeleteUser(ctx context.Context, name string) error {

	_, err := r.db.User.Delete().Where(user.NameEQ(name)).Exec(ctx)

	if err != nil {
		return err
	}
	return nil
}
*/
