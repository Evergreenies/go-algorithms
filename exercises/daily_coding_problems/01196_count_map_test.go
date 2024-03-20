package dailycondingproblems

import (
	"math"
	"testing"
)

/**
* Create a data structure that performs all the following operations in O(1) time:
*
* plus: Add a key with value 1. If the key already exists, increment its value by one.
* minus: Decrement the value of a key. If the key's value is currently 1, remove it.
* get_max: Return a key with the highest value.
* get_min: Return a key with the lowest value.
**/

type CountMap struct {
	countMap map[interface{}]int
	maxKey   interface{}
	minKey   interface{}
}

func (cm *CountMap) plus(key interface{}) {
	if _, ok := cm.countMap[key]; !ok {
		cm.countMap[key] = 1
	} else {
		cm.countMap[key]++
	}

	if cm.minKey == nil || cm.countMap[key] < cm.countMap[cm.minKey] {
		cm.minKey = key
	}

	if cm.maxKey == nil || cm.countMap[key] > cm.countMap[cm.maxKey] {
		cm.maxKey = key
	}
}

func (cm *CountMap) minus(key interface{}) {
	if _, ok := cm.countMap[key]; !ok {
		return
	}

	cm.countMap[key]--
	if cm.countMap[key] == 0 {
		delete(cm.countMap, key)
	}

	if key == cm.minKey {
		var minKey interface{}
		var minValue int

		minValue, ok := cm.countMap[cm.minKey]
		if !ok {
			minValue = int(math.Inf(1))
		}

		for key, value := range cm.countMap {
			if value <= minValue {
				minKey = key
				break
			}
		}

		cm.minKey = minKey
	}

	if key == cm.maxKey {
		var maxKey interface{}
		var maxValue int

		maxValue, ok := cm.countMap[key]
		if !ok {
			maxValue = int(math.Inf(-1))
		}

		for key, value := range cm.countMap {
			if value > maxValue {
				maxKey, maxValue = key, value
			}
		}

		cm.maxKey = maxKey
	}
}

func (cm *CountMap) getMax() interface{} {
	return cm.maxKey
}

func (cm *CountMap) getMin() interface{} {
	return cm.minKey
}

func assert(t *testing.T, actual interface{}, expected interface{}) {
	if actual == expected {
		t.Logf("PASSED: \n\tActual: %v\n\tExpected: %v\n", actual, expected)
	} else {
		t.Logf("FAILED: \n\tActual: %v\n\tExpected: %v\n", actual, expected)
	}
}

func TestCountMap(t *testing.T) {
	var countMap CountMap
	countMap.countMap = make(map[interface{}]int)

	countMap.minus("a")
	assert(t, countMap.getMax(), nil)
	assert(t, countMap.getMin(), nil)

	countMap.plus("a")
	assert(t, countMap.getMax(), "a")
	assert(t, countMap.getMin(), "a")

	countMap.plus("a")
	assert(t, countMap.getMax(), "a")
	assert(t, countMap.getMin(), "a")

	countMap.plus("b")
	assert(t, countMap.getMax(), "a")
	assert(t, countMap.getMin(), "b")

	countMap.minus("b")
	assert(t, countMap.getMax(), "a")
	assert(t, countMap.getMin(), "a")

	countMap.minus("b")
	assert(t, countMap.getMax(), "a")
	assert(t, countMap.getMin(), "a")

	countMap.minus("a")
	assert(t, countMap.getMax(), "a")
	assert(t, countMap.getMin(), "a")

	countMap.minus("a")
	assert(t, countMap.getMax(), nil)
	assert(t, countMap.getMin(), nil)
}
