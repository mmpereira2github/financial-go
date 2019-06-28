package indicator

import (
	"financial/indicator"
	boot "financial/init"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	boot.RepoLoader()
	code := m.Run()
	os.Exit(code)
}

func Test_PerfIndicatorType_FindByName(t *testing.T) {
	if ref := indicator.GetPerfIndicatorTypeDao().FindByName("IPCA"); ref == nil {
		t.Errorf("It should have found IPCA")
	} else {
		t.Log(ref)
	}
}

func Test_PerfIndicatorType_FindByName_NotFound(t *testing.T) {
	if indicator.GetPerfIndicatorTypeDao().FindByName("IP") != nil {
		t.Errorf("Expected not found")
	}
}

func Test_Preloaded_PerfIndicatorType(t *testing.T) {
	if ref := indicator.GetPerfIndicatorTypeDao().FindByName("CDI"); ref == nil {
		t.Errorf("It should have found CDI")
	}
	if ref := indicator.GetPerfIndicatorTypeDao().FindByName("IPCA"); ref == nil {
		t.Errorf("It should have found IPCA")
	}
}
