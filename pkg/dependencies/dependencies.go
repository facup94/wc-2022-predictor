package dependencies

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/facup94/worldcup-predictor/pkg/environment"
)

// DependencyManager defines an entity that provides storages (or external dependencies)
// for the app, depending on the current environment it is running
type DependencyManager interface {
	DB() *sql.DB
}

func GetDependencyManager(env environment.Environment) (DependencyManager, error) {
	switch env {
	case environment.Production:
		return newProdDependencyManager(env)
	case environment.Staging:
		return newStagingDependencyManager(env)
	case environment.Local:
		return newLocalDependencyManager(env)

	}
	return nil, errors.New(fmt.Sprintf("no dependency manager found for env: '%s'", env.String()))
}
