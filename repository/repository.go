package repository

import (
	"github.com/megaqstar/web-core/client"
	"github.com/megaqstar/web-core/repository/user"
)

type Repository struct {
	User user.Repository
}

func NewRepository(client *client.Client) *Repository {
	return &Repository{
		User: user.NewPG(client),
	}
}
