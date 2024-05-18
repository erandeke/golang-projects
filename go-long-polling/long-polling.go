package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type CappedQueueBuffer[T any] struct {
	items    []T
	capacity int
}

type SendMessageEx struct {
	Message string `json:"message"`
}

func PopulateCappedQueueBuffer[T any](capacity int) *CappedQueueBuffer[T] {
	return &CappedQueueBuffer[T]{
		items:    make([]T, 0, capacity),
		capacity: capacity,
	}
}

func (q *CappedQueueBuffer[T]) Append(item T) {

	if l := len(q.items); l == 0 {
		q.items = append(q.items, item)
	}

}

func (q *CappedQueueBuffer[T]) copy() []T {

	copied := make([]T, len(q.items))

	for i, elements := range q.items {
		copied[i] = elements
	}

	return copied
}

func main() {

	q := PopulateCappedQueueBuffer[string](10)
	e := echo.New()

	e.GET("updates", func(c echo.Context) error {
		return c.JSON(200, q.copy())
	})

	e.POST("send", func(c echo.Context) error {
		var request SendMessageEx
		if err := c.Bind(&request); err != nil {
			return c.String(400, fmt.Sprintf("Bad request: %v", err))
		}
		q.Append(request.Message)
		return c.JSON(201, "I've sent your request.")
	})

}
