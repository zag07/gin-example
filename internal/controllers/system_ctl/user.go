package system_ctl

import (
	"context"
	"fmt"
	"github.com/zs368/gin-example/internal/data/ent"
	"github.com/zs368/gin-example/internal/data/ent/user"
	"log"
)

type User struct{}

func NewUser() User { return User{} }

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetName("a8m").
		SetEmail("a8m@facebook.com").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name("a8m")).
		// `Only` 在 找不到用户 或 找到多于一个用户 时报错,
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
