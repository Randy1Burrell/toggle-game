package application

import (
	"sync"

	"github.com/randy1burrell/toggle-game/pkg/config"
	"github.com/randy1burrell/toggle-game/pkg/db"
)

type App struct {
	Cfg *config.Config
	DB  *db.Config
}

var cfg *config.Config
var dbConf *db.Config
var once sync.Once

func Get() *App {
	once.Do(func() {
		cfg = config.Get()
		dbConf = db.Get()
	})

	return &App{
		DB:  dbConf,
		Cfg: cfg,
	}
}
