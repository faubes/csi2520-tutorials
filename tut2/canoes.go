package main

import (
	"fmt"
	"time"
)

type canoe struct {
	name     string
	capacity int
}

const delay = 25 * time.Millisecond

func load(empty_boats <-chan canoe,
	loaded_boats chan<- canoe,
	remaining_load chan int,
	done_loading chan<- bool) {
	load := <-remaining_load
	if load <= 0 {
		done_loading <- true
		remaining_load <- load
		return
	}
	c := <-empty_boats
	next_remaining_load := load - c.capacity
	var current_load int
	var remaining int
	if next_remaining_load <= 0 {
		current_load = load
		remaining = 0
	} else {
		current_load = c.capacity
		remaining = next_remaining_load
	}

	fmt.Printf("Loading %s with %d\n", c.name, current_load)
	fmt.Printf("%d remaining\n", remaining)
	time.Sleep(10 * delay)
	remaining_load <- remaining
	loaded_boats <- c
}

func unload(loaded_boats <-chan canoe,
	empty_boats chan<- canoe,
	done_loading <-chan bool,
	all_done chan bool) {
	c := <-loaded_boats
	fmt.Printf("%s traveling...\n", c.name)
	time.Sleep(15 * delay)
	fmt.Printf("%s unloading...\n", c.name)
	time.Sleep(10 * delay)
	fmt.Printf("%s returning...\n", c.name)
	time.Sleep(5 * delay)
	empty_boats <- c
	select {
	case <-done_loading:
		all_done <- true
	default:
		return
	}
}

func start(units int,
	empty_boats chan canoe,
	loaded_boats chan canoe,
	remaining_load chan int,
	done_loading chan bool,
	all_done chan bool) {

	fmt.Printf("Delivering %d units.\n", units)
	remaining_load <- units

	done_canoes := 0
	for {
		select {
		case <-all_done:
			done_canoes++
			if done_canoes == 2 {
				return
			} // break
		default:
			go load(empty_boats, loaded_boats,
				remaining_load, done_loading)
			go unload(loaded_boats, empty_boats,
				done_loading, all_done)
		}
	}
}

func main() {
	empty_boats := make(chan canoe, 2)
	empty_boats <- canoe{"Small canoe", 10}
	empty_boats <- canoe{"Big canoe", 15}

	loaded_boats := make(chan canoe, 2)

	remaining_load := make(chan int, 1)

	done_loading := make(chan bool, 2)
	all_done := make(chan bool, 2)

  units := 65

	start(units, empty_boats, loaded_boats,
		remaining_load, done_loading, all_done)

	fmt.Println("All units delivered.")
}
