package repository

import (
	"context"
)

const PRODUCT_TABLE_NAME = "products"

type ProductRepository interface {
	Save(name string) error
}

func (repository Repository) Save(id string, name string) error {
	_, err := repository.DB.Query(context.Background(), "INSERT INTO "+PRODUCT_TABLE_NAME+" (id, name) VALUES ($1, $2)", id, name)
	if err != nil {
		return err
	}
	return nil
}

// func FindByID(ctx context.Context, id int) (*model.User, error) {
// 	var user model.User
// 	err := r.db.QueryRow(ctx, "SELECT id, username, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func Update(ctx context.Context, user *model.User) error {
// 	_, err := r.db.Exec(ctx, "UPDATE users SET username = $1, email = $2 WHERE id = $3", user.Username, user.Email, user.ID)
// 	return err
// }

// func Delete(ctx context.Context, id int) error {
// 	_, err := r.db.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
// 	return err
// }
