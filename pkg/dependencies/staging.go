package dependencies

import (
	"database/sql"

	"github.com/facup94/worldcup-predictor/pkg/environment"
	"github.com/facup94/worldcup-predictor/pkg/storage"
)

type StagingDependencyManager struct {
	db *sql.DB
}

func newStagingDependencyManager(env environment.Environment) (StagingDependencyManager, error) {
	db, err := storage.NewMySQL(env)
	if err != nil {
		return StagingDependencyManager{}, err
	}
	return StagingDependencyManager{db: db}, nil
}

func (depMng StagingDependencyManager) DB() *sql.DB {
	return depMng.db
}
