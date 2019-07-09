## Web Service Application to update values using CDI and IPCA to compare investment performance
`go get -u github.com/mmpereira2github/financial-go`

# Version 0

Version 0 has just a test that shows how the UpdateValueService works.
Next release it will be published as a REST API.

# UpdateValueService

This service receives as input the following data:

* value to be updated (float64 format)
* Date associated with the given value
* targetDate that is to when you want to have the value updated
* indexID that can be CDI or IPCA

The output is:

* updatedValue that is the given value updated using CDI or IPCA as reference from the given date until target date.

The test case UpdateValueService_test.go provides an example.