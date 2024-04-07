package ringy

import "errors"

var QueueEmpty = errors.New("queue empty")

var QueueFull = errors.New("queue full")
