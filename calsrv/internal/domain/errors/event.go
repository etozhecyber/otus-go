package errors

/*EventError bla*/
type EventError string

func (ee EventError) Error() string {
	return string(ee)
}

var (
	/*ErrIncorectEndDate start_date after end_date*/
	ErrIncorectEndDate = EventError("start_date after end_date")
	/*ErrIncorectEndDate start_date on weekend*/
	ErrWeekendStartDate = EventError("Start_date on weekend")
	/*ErrIncorectEndDate End_date on weekend*/
	ErrWeekendEndDate = EventError("End_date on weekend")
)
