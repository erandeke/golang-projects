package main

import (
	"fmt"
	"sync"

	"github.com/labstack/echo/v4"
)

// cappedQueue stores oreder in FIFO way
type CappedQueue[T any] struct {
	items    []T
	lock     *sync.RWMutex
	capacity int
}

type SendMessage struct {
	Message string `json:"message"`
}

func NewCappedQueue[T any](capacity int) *CappedQueue[T] {
	return &CappedQueue[T]{
		items:    make([]T, 0, capacity),
		lock:     new(sync.RWMutex),
		capacity: capacity,
	}
}

func (q *CappedQueue[T]) Append(item T) {

	q.lock.Lock()

	defer q.lock.Unlock()

	if l := len(q.items); l == 0 {
		q.items = append(q.items, item)
	} else {
		to := q.capacity - 1
		if l < q.capacity {
			to = l
		}
		q.items = append([]T{item}, q.items[:to]...)
	}

}

func (q *CappedQueue[T]) Copy() []T {

	q.lock.RLock()

	defer q.lock.RUnlock()

	copiedItems := make([]T, len(q.items))

	for index, eachItem := range q.items {
		copiedItems[index] = eachItem
	}

	return copiedItems

}

func main_1() {

	q := NewCappedQueue[string](10)
	e := echo.New()

	e.GET("updates", func(c echo.Context) error {
		return c.JSON(200, q.Copy())
	})

	e.POST("send", func(c echo.Context) error {
		var request SendMessage
		if err := c.Bind(&request); err != nil {
			return c.String(400, fmt.Sprintf("Bad request: %v", err))
		}
		q.Append(request.Message)
		return c.JSON(201, "I've sent your request.")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
