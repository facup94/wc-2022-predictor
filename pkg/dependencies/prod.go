package dependencies

import (
	"database/sql"

	"github.com/facup94/worldcup-predictor/pkg/environment"
	"github.com/facup94/worldcup-predictor/pkg/storage"
)

type ProdDependencyManager struct {
	db *sql.DB
}

func newProdDependencyManager(env environment.Environment) (ProdDependencyManager, error) {
	db, err := storage.NewMySQL(env)
	if err != nil {
		return ProdDependencyManager{}, err
	}
	return ProdDependencyManager{db: db}, nil
}

func (depMng ProdDependencyManager) DB() *sql.DB {
	return depMng.db
}
