package manager

import (
	"enigmacamp.com/goacc/config"
	"enigmacamp.com/goacc/utils"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Infra interface {
	SqlDb() *sqlx.DB
	IntraClient() *utils.IntraClient
}
type infra struct {
	db          *sqlx.DB
	intraClient *utils.IntraClient
	cfg         config.Config
}

func (i *infra) SqlDb() *sqlx.DB {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", i.cfg.User, i.cfg.Password, i.cfg.Host, i.cfg.Port, i.cfg.Name)
	conn, err := sqlx.Connect("pgx", dataSourceName)
	if err != nil {
		panic(err)
	}
	i.db = conn
	return i.db
}

func (i *infra) IntraClient() *utils.IntraClient {
	i.intraClient = new(utils.IntraClient)
	return i.intraClient
}

func NewInfra(config config.Config) Infra {
	return &infra{
		cfg: config,
	}
}
