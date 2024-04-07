# Ringy

A simple, generic ring data structure.

## Using

The API is very simple, create a new ring, then add and pop values
from it, which happens in FIFO order. It is also possible to get
the capacity (`Cap()`) and current length (`Len()`) of the ring.

```go
package main

import (
	"errors"

	"github.com/glesica/ringy"
)

func main() {
	r := ringy.New[string](5) // Create a ring of strings with capacity 5
	
	err := r.Add("hello")
	if errors.Is(err, ringy.QueueFull) {
		// Ohnoes!
	}
	
	value, err := r.Pop()
	if errors.Is(err, ringy.QueueEmpty) {
		// Disaster!
	}
	
	println(value) // prints "hello"
}
```

The entire capacity is allocated immediately and no further allocations
will occur.

The ring is NOT thread-safe, so you have to do your own
synchronization if you are using multiple goroutines.
