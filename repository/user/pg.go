package user

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/megaqstar/web-core/client"
	"github.com/megaqstar/web-core/meg-pkg/common"
	"github.com/megaqstar/web-core/model"
	"gorm.io/gorm"
)

type pgRepository struct {
	mysql *gorm.DB
}

func NewPG(client *client.Client) Repository {
	return &pgRepository{
		mysql: client.MySQL(),
	}
}

func (p pgRepository) Create(ctx echo.Context, data *model.User) error {
	return p.mysql.Create(data).Error
}

func (p pgRepository) GetByID(ctx echo.Context, id int64) (*model.User, error) {
	var user *model.User
	if err := p.mysql.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (p pgRepository) Update(ctx echo.Context, data *model.User) error {
	return p.mysql.Save(data).Error
}

func (p pgRepository) Delete(ctx echo.Context, data *model.User, unscoped bool) error {
	db := p.mysql
	if unscoped {
		db = db.Unscoped()
	}
	return db.Delete(data).Error
}

func (p pgRepository) GetList(
	ctx echo.Context,
	paginator common.Paginator,
	conditions interface{},
	order string) ([]model.User, error) {
	var (
		db     = p.mysql.Model(&model.User{})
		data   = make([]model.User, 0)
		offset int
	)

	if conditions != nil {
		db = db.Where(conditions)
	}

	if order != "" {
		db = db.Order(order)
	}

	paginator.Format()
	if paginator.Page != 1 {
		offset = paginator.Limit * (paginator.Page - 1)
	}
	fmt.Println("--------------------------------")
	err := db.Limit(paginator.Limit).Offset(offset).Find(&data).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(len(data))
	return data, nil
}
