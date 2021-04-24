package storage

import (
	"database/sql"
	"fmt"

	"github.com/facup94/worldcup-predictor/pkg/config"
	"github.com/facup94/worldcup-predictor/pkg/environment"

	_ "github.com/go-sql-driver/mysql"
)

const connectionStringFormat = "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true"

// NewMySQL initializes a DB connection pool to a MySQL database
func NewMySQL(env environment.Environment) (*sql.DB, error) {
	cfg, err := config.GetDBConfig(env)

	connectionString := fmt.Sprintf(connectionStringFormat, cfg.Username, cfg.Password, cfg.Host, cfg.Name)

	db, err := sql.Open("mysql", connectionString)
	if err == nil {
		err = db.Ping()
	}
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConnections)

	return db, nil
}
