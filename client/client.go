package client

import (
	"gorm.io/gorm"
)

type Client struct {
	mysqlDB *gorm.DB
}

func (c *Client) MySQL() *gorm.DB {
	return c.mysqlDB
}

func NewClient(
	mysqlDB *gorm.DB,
) *Client {
	return &Client{
		mysqlDB: mysqlDB,
	}
}
