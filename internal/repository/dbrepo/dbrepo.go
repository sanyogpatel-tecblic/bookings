package dbrepo

import (
	"database/sql"

	"github.com/sanyogpatel-tecblic/bookings/internal/config"
	"github.com/sanyogpatel-tecblic/bookings/internal/repository"
)

type postgressDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}
type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewpostgressRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgressDBRepo{
		App: a,
		DB:  conn,
	}
}

func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &postgressDBRepo{
		App: a,
	}
}
