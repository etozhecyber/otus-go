package errors

/*EventError bla*/
type EventError string

func (ee EventError) Error() string {
	return string(ee)
}

var (
	/*ErrOverlaping is another event exists for this date*/
	ErrOverlaping = EventError("another event exists for this date")
	/*ErrDateBusy date is busy*/
	ErrDateBusy = EventError("date is busy")
)
