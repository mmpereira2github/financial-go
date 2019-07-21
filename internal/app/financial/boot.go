package financial

import (
	"log"

	"github.com/mmpereira2github/financial-go/internal/pkg/config"
	"github.com/mmpereira2github/financial-go/internal/pkg/infra/dao"
	"github.com/mmpereira2github/financial-go/internal/pkg/infra/dao/json"
)

// Boot boots the app initializing its components
func Boot(appHome string) {
	config.LoadConfigFromAppHome((appHome))
	var indexDaoImpl dao.IndexDao
	indexDaoImplName := config.Config.DaoConfig.IndexDaoImplName
	log.Printf("config.Config.DaoConfig.IndexDaoImplName=%s", indexDaoImplName)
	switch indexDaoImplName {
	case "JSON":
		indexDaoImpl = json.NewJSONFileIndexDao()
	default:
		indexDaoImpl = json.NewJSONFileIndexDao()
	}

	dao.SetIndexDao(indexDaoImpl)
}
