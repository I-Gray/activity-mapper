package calendar

type EventSourceAdapter interface {
	GetData() (interface{}, error)
}

type CalendarStrategy interface {
	AddEvents(eventSource EventSourceAdapter) error
}
