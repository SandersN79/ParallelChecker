package core

import (
	"context"
	"fmt"
	"sync"
)

// DuplicateCheck for duplicate codes in inputted csv files
func DuplicateCheck(files []string) error {
	// initialize process wait group and context with cancel function
	var wg sync.WaitGroup //wait for all processes to be done.
	ctx, cancel := context.WithCancel(context.Background())
	b := board{
		values: map[string]value{},
		cancel: cancel, //cancels when match is found or all csvs are checked
	}
	c := b.start()
	defer cancel() // defer cancel to release resources even if no duplicates are found
	for _, f := range files {
		wg.Add(1)
		//f string below allows the go func to use the most recent f import
		go func(f string) {
			defer wg.Done() //has a -1 built in to close the wg.add
			rf, err := newCsvFile(f)
			if err != nil {
				return
			}
			for i, cd := range rf.csvData {
				select { // Check if any duplicate found by other goroutines:
				case <-ctx.Done():
					return // duplicate found somewhere, terminate
				default: // Default avoids blocking
				}
				c <- value{cd.code, i + 2, f} // Send csv row data to shared values board
			}
		}(f)
	}
	wg.Wait()        //waits for all processes to be finished above
	close(c)         //refers to the shared memory which is the board
	return ctx.Err() //if duplicate is found, this will return an error due to duplicate map indexes
}

type value struct {
	key  string
	row  int
	file string
}

type board struct {
	values map[string]value
	cancel func()
}

// opens up a channel of values to be streamed to the board
func (b *board) start() chan<- value {
	c := make(chan value)
	go func() {
		for v := range c {
			b.notify(v)
		}
	}()
	return c
}

func (b *board) notify(v value) {
	if _, ok := b.values[v.key]; !ok { // if not ok, then no duplicate
		b.values[v.key] = v
		return
	}
	b.cancel() // if duplicate, cancel all goroutines
	d := b.values[v.key]
	fmt.Printf("Duplicate found... Code: %s\nA) File: (%s); Row: %v\nB) File: (%s); Row: %v\n",
		v.key, v.file, v.row, d.file, d.row)
}
