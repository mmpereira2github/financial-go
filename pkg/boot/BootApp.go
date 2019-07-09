package boot

import (
	"financial/pkg/config"
	"financial/pkg/infra/dao"
	"financial/pkg/infra/dao/daoimpl"
	"log"
)

// BootApp boots the app initializing its components
func BootApp(appHome string) {
	var indexDaoImpl dao.IndexDao
	indexDaoImplName := config.Config.DaoConfig.IndexDaoImplName
	log.Printf("config.Config.DaoConfig.IndexDaoImplName=%s", indexDaoImplName)
	switch indexDaoImplName {
	case "JSON":
		indexDaoImpl = daoimpl.NewJSONFileIndexDao()
	default:
		indexDaoImpl = daoimpl.NewJSONFileIndexDao()
	}

	dao.SetIndexDao(indexDaoImpl)
}
