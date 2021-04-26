package mocking

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

// 记录次数
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls += 1
}

// 真正的sleeper
type ConfigurableSleeper struct {
	duration time.Duration
}

func (s *ConfigurableSleeper) Sleep() {
	time.Sleep(s.duration)
}

// 记录顺序
const (
	sleep = "sleep"
	write = "write"
)

type CounterDownSleeper struct {
	Calls []string
}

func (c *CounterDownSleeper) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CounterDownSleeper) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

func CounterDown(writer io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}
	sleeper.Sleep()
	fmt.Fprint(writer, "Go!")
}
