package easy

import (
	"time"
)

/*
	As per notes from https://www.cs.rutgers.edu/~pxk/417/notes/clocks/index.html
*/

type vector map[string]int

type event struct {
	process   string
	value     interface{}
	timestamp time.Time
	vectorMap vector
}

func (c *event) setVal(val interface{}) {
	c.value = val
}

func (c *event) receive(d *event) {
	for k, v := range d.vectorMap {
		val, found := c.vectorMap[k]
		if found {
			if v > val {
				c.vectorMap[k] = v
			}
		} else {
			c.vectorMap[k] = v
		}
	}
	c.value = d.value
	c.vectorMap[c.process]++
}

func newEvent(name string) *event {
	vs := make(vector, 10)
	vs[name] = 0

	return &event{
		process:   name,
		vectorMap: vs,
	}
}

func send(from, to *event) {
	to.receive(from)
}
