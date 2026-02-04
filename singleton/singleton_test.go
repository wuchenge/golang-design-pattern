package singleton

import (
	"log"
	"sync"
	"testing"
)

const parCount = 100

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()

	if ins1 != ins2 {
		t.Error("Expected both instances to be the same")
	}
}

func TestParallelSingleton(t *testing.T) {
	start := make(chan struct{})
	var wg sync.WaitGroup
	instances := make([]*Singleton, parCount)

	for i := 0; i < parCount; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			<-start
			instances[index] = GetInstance()
		}(i)
	}

	close(start)
	wg.Wait()

	firstInstance := instances[0]
	for i, instance := range instances {
		if instance != firstInstance {
			t.Errorf("Instance at index %d is not the same as the first instance", i)
		}
	}

	log.Printf("All %d instances are the same", parCount)
}
