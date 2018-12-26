package turn

import (
	"testing"
	"time"
)

func Test_TurnGet(t *testing.T) {
	turn := NewTurn(1)
	go func() {
		turn.Get()
		time.Sleep(time.Second * 3)
		turn.Free()
	}()
	now := time.Now()
	time.Sleep(time.Second)
	println("wait...", time.Since(now).String())
	turn.Get()
	println("in...", time.Since(now).String())
	turn.Free()
}

func Test_TurnWait(t *testing.T) {
	turn := NewTurn(1)
	go func() {
		turn.Get()
		time.Sleep(time.Second * 2)
		turn.Free()
	}()

	time.Sleep(time.Millisecond * 10)

	if err := turn.Wait(6 * time.Second); err != nil {
		println(err.Error())
	} else {
		println("ok")
		turn.Free()
	}
}
