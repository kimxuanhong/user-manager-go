package dependencies

import (
	"github.com/kimxuanhong/user-manager-go/internal/config"
)

type Dependency struct {
	Db *config.Database
}
