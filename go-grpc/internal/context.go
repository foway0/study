package internal

import (
	"database/sql"
	config2 "github.com/foway0/study/go-grpc/internal/config"
	"github.com/foway0/study/go-grpc/internal/infra"
	"log"
)

type driver struct {
	mysql *sql.DB
}

type ApplicationContext struct {
	config config2.Config
	driver driver
}

func NewApplicationContext() *ApplicationContext {
	config := config2.GetConfig()

	return &ApplicationContext{
		config: config,
		driver: driver{
			mysql: infra.InitMySQL(config),
		},
	}
}

func (c *ApplicationContext) Config() config2.Config {
	return c.config
}

func (c *ApplicationContext) Mysql() *sql.DB {
	if c.driver.mysql == nil {
		log.Fatal("MySQL driver is not initialized")
	}
	return c.driver.mysql
}
