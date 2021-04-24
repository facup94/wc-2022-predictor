// Package environment is used to determine the current environment in which the app is running
package environment

import "os"

type Environment int

const (
	Production Environment = iota
	Staging
	Local
)

func GetEnvironment() Environment {
	switch os.Getenv("ENVIRONMENT") {
	case "prod":
		return Production
	case "staging":
		return Staging
	case "local":
		fallthrough
	default:
		return Local
	}
}

func (env Environment) String() string {
	switch env {
	case Production:
		return "local"
	case Staging:
		return "staging"
	case Local:
		return "local"
	}
	return ""
}
