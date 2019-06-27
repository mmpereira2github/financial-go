package index

import (
	"financial/index"
	"fmt"
	"testing"
)

func Test_RefIndexDao_GetById(t *testing.T) {
	if ref, error := index.GetRefIndexDao().GetById(1); error != nil {
		panic(error)
	} else {
		fmt.Println(ref)
	}
}

func Test_RefIndexDao_FindByName(t *testing.T) {
	if ref := index.GetRefIndexDao().FindByName("IPCA"); ref == nil {
		panic(fmt.Errorf("It should have found IPCA"))
	} else {
		fmt.Println(ref)
	}
}

func Test_RefIndexDao_FindByName_NotFound(t *testing.T) {
	if index.GetRefIndexDao().FindByName("IP") != nil {
		panic(fmt.Errorf("Expected not found"))
	}
}
