package financial

import (
	"financial/internal/pkg/config"
	"financial/internal/pkg/infra/dao"
	"financial/internal/pkg/infra/dao/json"
	"log"
)

// Boot boots the app initializing its components
func Boot(appHome string) {
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
