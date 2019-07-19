package config

import (
	"financial/internal/pkg/config"
	"testing"
)

/**
-> Teste 01

Dado: nada

Pergunta:
	Qual a configuração default?

Esperado:
	ler a configuração padrão do arquivo configs/financial.json

*/
func Test_ReadConfigDefault(t *testing.T) {
	if err := config.LoadConfig("../../../configs/financial.json"); err != nil {
		t.Error(err)
	}
	indexDaoImplName := config.Config.DaoConfig.IndexDaoImplName
	if indexDaoImplName != "JSON" {
		t.Errorf("FinancialConfig.DaoConfig.PerfIndicatorTypeDaoName=%s not expected", indexDaoImplName)
	}
}
