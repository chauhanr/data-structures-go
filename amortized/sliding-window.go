package amortized

import "log"

type Window struct {
	array []int
	Queue []int
	s     int
	e     int
}

/**
  returns true if the slide is done and false if a slide not further take place.
*/
func (w *Window) Slide() bool {
	if w.e == len(w.array)-1 {
		return false
	}
	w.s++
	w.e++
	element := w.array[w.e]
	if w.array[w.s] == w.Queue[0] {
		w.Queue = w.Queue[1:len(w.Queue)]
	}
	log.Printf("%v", w.Queue)
	for len(w.Queue) != 0 && w.Queue[len(w.Queue)-1] >= element {
		w.Queue = w.Queue[0 : len(w.Queue)-1]
	}
	w.Queue = append(w.Queue, element)
	return true
}

func (w *Window) Initialize(a []int, size int) Window {
	win := Window{array: a, s: 0, e: size - 1}
	queue := []int{}
	for i := 0; i < size; i++ {
		if len(queue) == 0 {
			queue = append(queue, a[i])
		} else {
			for j := len(queue) - 1; j >= 0; j-- {
				if queue[j] >= a[i] {
					if j == 0 {
						queue = []int{}
					} else {
						queue = queue[0:j]
					}
				} else {
					break
				}
			}
			queue = append(queue, a[i])
		}
	}
	win.Queue = queue
	return win
}

func Display(a []int, size int) {
	win := Window{}
	win = win.Initialize(a, size)
	log.Printf("Start : %d, End: %d Min: %d, %v", win.s, win.e, win.Queue[0], win.Queue)
	for win.Slide() {
		log.Printf("Start : %d, End: %d Min: %d, %v", win.s, win.e, win.Queue[0], win.Queue)
	}
}
