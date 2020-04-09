package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Chopstick struct {
	sync.Mutex
}

type Philo struct {
	id            int
	leftC, rightC *Chopstick
}

func (p Philo) eat(wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		fmt.Println("Starting to eat ", p.id+1)
		p.leftC.Lock()
		p.rightC.Lock()
		time.Sleep(10 * time.Millisecond)
		p.leftC.Unlock()
		p.rightC.Unlock()
		fmt.Println("Finishing eating ", p.id+1)

	}
	wg.Done()
}
func host(p chan bool) {
	p <- true
}
func main() {
	ChopS := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		ChopS[i] = new(Chopstick)
	}
	philos := make([]*Philo, 5)
	nums := []int{0, 1, 2, 3, 4}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
	var wg sync.WaitGroup
	for k := 0; k < 5; k++ {
		j := nums[k]
		philos[j] = &Philo{j, ChopS[j], ChopS[(j+1)%5]}
	}
	for i := 0; i < 5; i += 2 {
		c1 := make(chan bool)
		c2 := make(chan bool)
		go host(c1)
		go host(c2)
		philo1 := <-c1
		philo2 := <-c2
		if philo1 && philo2 {
			wg.Add(1)
			go philos[i].eat(&wg)
			if i+1 < 5 {
				go philos[i+1].eat(&wg)
			}
		}
		wg.Wait()
	}
}
