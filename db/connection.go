package db

import (
	"database/sql"
	"fmt"

	"github.com/italopatrick/api-todo/configs"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()
	sc := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	fmt.Println(conf)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()

	return conn, err
}
