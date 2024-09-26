package repository

import "github.com/kimxuanhong/user-manager-go/pkg/infra/config"

type Repository struct {
	DB *config.Datasource `inject:""`
}
