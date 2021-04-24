package dependencies

import (
	"database/sql"

	"github.com/facup94/worldcup-predictor/pkg/environment"
	"github.com/facup94/worldcup-predictor/pkg/storage"
)

type LocalDependencyManager struct {
	db *sql.DB
}

func newLocalDependencyManager(env environment.Environment) (LocalDependencyManager, error) {
	db, err := storage.NewMySQL(env)
	if err != nil {
		return LocalDependencyManager{}, err
	}
	return LocalDependencyManager{db: db}, nil
}

func (depMng LocalDependencyManager) DB() *sql.DB {
	return depMng.db
}
