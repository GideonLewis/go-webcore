package user

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/megaqstar/web-core/payload"
	"github.com/megaqstar/web-core/presenter"
)

func (u *UseCase) GetList(ctx echo.Context, req *payload.UserGetListRequest) (*presenter.ListUserResponseWrapper, error) {
	var (
		conditions = map[string]interface{}{}
		orderBy    string
	)

	if req.OrderBy.Field != "" {
		orderBy = fmt.Sprintf("%s %s", req.OrderBy.Field, req.OrderBy.Sort)
	}

	if len(req.Conditions) != 0 {
		for _, cond := range req.Conditions {
			conditions[cond.Field] = cond.Value
		}
	}

	users, err := u.UserRepo.GetList(ctx, req.Paginator, conditions, orderBy)
	if err != nil {
		return nil, err
	}

	return &presenter.ListUserResponseWrapper{
		Users: nil,
		Meta: map[string]interface{}{
			"page":  req.Paginator.Page,
			"limit": req.Paginator.Limit,
			"total": len(users),
		},
	}, nil
}
