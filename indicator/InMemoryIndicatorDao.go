package indicator

var inMemoryIndicatorRepo = map[string][]MonthlyPerfIndicatorValue{}

type inMemoryIndicatorDao struct{}

func (d *inMemoryIndicatorDao) FindByRange(cat PerfIndicatorCategory, startYear int, startMonth int, endYear int, endMonth int) []MonthlyPerfIndicatorValue {
	list := inMemoryIndicatorRepo[cat.Name()]
	if list != nil {
		result := make([]MonthlyPerfIndicatorValue, len(list))
		for _, value := range list {
			if value.Year() >= startYear && value.Month() <= startMonth {
				if value.Year() <= endYear && value.Month() <= endMonth {
					result = append(result, value)
				}
			}
		}
		return result
	}
	return []MonthlyPerfIndicatorValue{}
}

func (d *inMemoryIndicatorDao) Save(cat PerfIndicatorCategory, value MonthlyPerfIndicatorValue) error {
	return nil
}
