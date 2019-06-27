package indicator

import (
	"financial/indicator"
	"testing"
)

func Test_PerfIndicatorType_GetById(t *testing.T) {
	if ref, error := indicator.GetPerfIndicatorTypeDao().GetByID(1); error != nil {
		t.Error(error)
	} else {
		t.Log(ref)
	}
}

func Test_PerfIndicatorType_GetById_NotFound(t *testing.T) {
	id := 500
	if ref, error := indicator.GetPerfIndicatorTypeDao().GetByID(id); error != nil {
		t.Log(error)
	} else {
		t.Errorf("%v should have not been found with id=%d\n", ref, id)
	}
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
	} else {
		if ref.ID() != 2 {
			t.Errorf("CDI should have ID=2 and not=%d", ref.ID())
		}
	}
	if ref := indicator.GetPerfIndicatorTypeDao().FindByName("IPCA"); ref == nil {
		t.Errorf("It should have found IPCA")
	} else {
		if ref.ID() != 1 {
			t.Errorf("IPCA should have ID=1and not=%d", ref.ID())
		}
	}
}
