package services_test

import (
	"financial/pkg/boot"
	"financial/pkg/infra"
	"financial/pkg/services"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	boot.BootApp("../../..")
	code := m.Run()
	os.Exit(code)
}

func TestUpdate892_40From201212Til201302UsingCDI(t *testing.T) {
	value := float64(892.40)
	when := time.Date(2012, 12, 20, 0, 0, 0, 0, infra.PlatForm.Location)
	targetDate := time.Date(2013, 03, 20, 0, 0, 0, 0, infra.PlatForm.Location)
	input := services.UpdateValueServiceInput{Value: value, Date: when, TargetDate: targetDate, IndexID: "CDI"}
	status, output := services.UpdateValue(input)
	if status.Code != 0 {
		t.Error(status.Error)
	} else {
		expectedValue := 897.634818
		if (output.UpdatedValue - expectedValue) > 0.0000009 {
			t.Errorf("Expected updated value=%f, but received=%f", expectedValue, output.UpdatedValue)
		}
	}
}
