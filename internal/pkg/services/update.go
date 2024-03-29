package services

import (
	"log"
	"time"

	"github.com/mmpereira2github/financial-go/internal/pkg/index"
	"github.com/mmpereira2github/financial-go/internal/pkg/infra/dao"
)

// UpdateValueServiceInput is the input for the UpdatePriceService service
type UpdateValueServiceInput struct {
	Value      float64   `form:"value" json:"value" binding:"required"`
	Date       time.Time `json:"date"`
	TargetDate time.Time `json:"targetDate"`
	IndexID    string    `form:"index" json:"index" binding:"required"`
}

// UpdateValueServiceOutput is the output of the UpdatePriceService service
type UpdateValueServiceOutput struct {
	UpdatedValue float64
}

func init() {
	inputFactory := func() interface{} { return &UpdateValueServiceInput{} }
	invoke := func(input interface{}) (interface{}, *Status) {
		if inputCasted, ok := input.(*UpdateValueServiceInput); ok {
			return UpdateValue(*inputCasted)
		}
		return Status{0, nil}, nil
	}
	Manager.Register(&ServiceEntry{
		ID:           "UpdateValue",
		InputFactory: inputFactory,
		Invoke:       invoke,
	})
}

// UpdateValue update a given value collected at given date to target date using the given update index
func UpdateValue(input UpdateValueServiceInput) (UpdateValueServiceOutput, *Status) {
	log.Println("input", input)
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
			return UpdateValueServiceOutput{resultValue}, nil
		}
		return UpdateValueServiceOutput{}, &Status{Code: DateNotFound, Error: err}
	}
	return UpdateValueServiceOutput{}, &Status{Code: IndexIDNotFound, Error: err}
}
