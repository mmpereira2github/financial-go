package infra

import "time"

var PlatForm platform = newPlatform()

type platform struct {
	Location *time.Location
}

func newPlatform() platform {
	location, _ := time.LoadLocation("Local")
	return platform{location}
}
