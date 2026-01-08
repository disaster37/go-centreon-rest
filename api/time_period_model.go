package api

// TimePeriodCreateOrUpdateRequest represents the payload to create or update a time periode in Centreon.
type TimePeriodCreateOrUpdateRequest struct {
	Name       string                `json:"name" validate:"required,max=200"`
	Alias      string                `json:"alias" validate:"required,max=200"`
	Days       []TimePeriodDays      `json:"days" validate:"required,min=1"`
	Templates  []int64               `json:"templates"`
	Exceptions []TimePeriodException `json:"exceptions"`
}

// TimePeriodResponse represents a time periode in Centreon.
type TimePeriodDays struct {
	Day       TimeDay `json:"day" validate:"required"`
	TimeRange string  `json:"time_range" validate:"required"`
}

// TimePeriodException represents an exception period within a time periode.
type TimePeriodException struct {
	Id        *int64 `json:"id,omitempty"`
	DayRange  string `json:"day_range" validate:"required"`
	TimeRange string `json:"time_range,omitempty"`
}

// TimeDay represents a day of the week.
type TimeDay int

const (
	TimeDayMonday    TimeDay = 1
	TimeDayTuesday   TimeDay = 2
	TimeDayWednesday TimeDay = 3
	TimeDayThursday  TimeDay = 4
	TimeDayFriday    TimeDay = 5
	TimeDaySaturday  TimeDay = 6
	TimeDaySunday    TimeDay = 7
)

// TimePeriodResponse represents a time periode in Centreon.
type TimePeriodResponse struct {
	Id         int64                 `json:"id"`
	Name       string                `json:"name"`
	Alias      string                `json:"alias"`
	Days       []TimePeriodDays      `json:"days"`
	Templates  []IdName              `json:"templates,omitempty"`
	Exceptions []TimePeriodException `json:"exceptions,omitempty"`
}
