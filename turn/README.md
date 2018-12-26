# Turn

Turn a concurrent controller, it limit the maximum number of concurrent accesses.

## Installation
Install Turn with go tool:
```
    go get github.com/SongLiangChen/util/turn
```

## Usage

### Get
```go
package turn

import (
	"time"
	"github.com/SongLiangChen/util/turn"
)

func main() {
	t := turn.NewTurn(1)
	
	go func() {
		t.Get()
		defer t.Free()
	
		time.Sleep(time.Second)
		// ...
	}()
	
	time.Sleep(time.Millisecond*10)
	
	t.Get()
	// Arrive here after 1.01 second 
	defer t.Free()
}
```

### Wait
```go
package turn

import (
	"time"
	"github.com/SongLiangChen/util/turn"
)

func main() {
	t := turn.NewTurn(1)
	
	go func() {
		t.Get()
		defer t.Free()
	
		time.Sleep(time.Second)
		// ...
	}()
	
	time.Sleep(time.Millisecond*10)
	
	if err := t.Wait(time.Second*3); err != nil {
		println(err.Error())
		return
	}
	// Arrive here after 1.01 second 
	t.Free()
	
	go func() {
		t.Get()
		defer t.Free()
        	
		time.Sleep(time.Second*5)
		// ...
	}()
	
	if err := t.Wait(time.Second*3); err != nil {
	    // Will get error here after 3 second
		println(err.Error())
		return
    }
	t.Free()
}
```