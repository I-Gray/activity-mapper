package calendar

import "fmt"

func AddEvents(strategy CalendarStrategy, eventSource EventSourceAdapter) error {
	fmt.Println("Adding events to calendar")
	return strategy.AddEvents(eventSource)
}
