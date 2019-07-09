package services

import (
	"financial/pkg/index"
	"financial/pkg/infra/dao"
	"time"
)

// UpdateValueServiceInput is the input for the UpdatePriceService service
type UpdateValueServiceInput struct {
	Value      float64
	Date       time.Time
	TargetDate time.Time
	IndexID    string
}

// UpdateValueServiceOutput is the output of the UpdatePriceService service
type UpdateValueServiceOutput struct {
	UpdatedValue float64
}

// UpdateValue update a given value collected at given date to target date using the given update index
func UpdateValue(input UpdateValueServiceInput) (ServiceExecutionStatus, UpdateValueServiceOutput) {
	var err error
	var timeInterator index.TimeIterator
	var indexInstance *index.Index
	dao := dao.GetIndexDao()
	indexInstance, err = (*dao).FindByID(input.IndexID)
	if err == nil {
		timeInterator, err = index.NewTimeIterator(input.Date, indexInstance.IntervalType)
		if err == nil {
			valueCalculator := index.NewUpdateFactorCalculator(indexInstance)
			resultValue := input.Value
			nextTime := timeInterator.Next()
			endTimeReached := nextTime.After(input.TargetDate)
			for !endTimeReached {
				factor := valueCalculator.GetUpdateFactor(nextTime)
				resultValue = resultValue * factor
				nextTime = timeInterator.Next()
				endTimeReached = nextTime.After(input.TargetDate)
			}
			return ServiceExecutionStatus{Code: 0}, UpdateValueServiceOutput{resultValue}
		}
	}
	return ServiceExecutionStatus{Code: -1, Error: err}, UpdateValueServiceOutput{}
}
