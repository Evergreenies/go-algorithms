/**
* Implement the singleton pattern with a twist. First, instead of storing one instance,
* store two instances. And in every even call of getInstance(), return the first instance
* and in every odd call of getInstance(), return the second instance.
* */

package dailycodingproblems

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type singletonWithTwist struct {
	idx int
}

var (
	instances    [2]*singletonWithTwist
	counter      int
	instanceOnce [2]sync.Once

	lock sync.Mutex
)

func getInstance() *singletonWithTwist {
	lock.Lock()
	defer lock.Unlock()

	idx := counter % 2
	counter++

	instanceOnce[idx].Do(func() {
		fmt.Printf("initialized: %d\n", idx)
		instances[idx] = &singletonWithTwist{idx: idx}
	})

	return instances[idx]
}

func TestSingletonWithTwist(t *testing.T) {
	instance1 := getInstance()
	instance2 := getInstance()
	instance3 := getInstance()
	instance4 := getInstance()

	assert := assert.New(t)

	assert.False(instance1 == instance2)
	assert.True(instance1 == instance3)
	assert.True(instance2 == instance4)
	assert.False(instance3 == instance2)
}
