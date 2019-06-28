package boot

import (
	"financial/indicator"
	"fmt"
)

var repoInitialized = false

// RepoLoader preloads some data into repository
func RepoLoader() {
	if !repoInitialized {
		cdi := indicator.NewPerfIndicatorType("CDI")
		indicator.GetPerfIndicatorTypeDao().Save(cdi)
		fmt.Println("Loading CDI")
		ipca := indicator.NewPerfIndicatorType("IPCA")
		indicator.GetPerfIndicatorTypeDao().Save(ipca)
		fmt.Println("Loading IPCA")
		repoInitialized = true
	}
}
