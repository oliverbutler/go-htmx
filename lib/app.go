package lib

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Pool     *pgxpool.Pool
}

func NewApp(db *pgxpool.Pool) *App {
	return &App{
		Pool:     db,
	}
}
