/*
* Write a function to flatten a nested dictionary. Namespace the keys with a period.
* For example, given the following dictionary:
{
	    "key": 3,
	    "foo": {
	        "a": 5,
	        "bar": {
	            "baz": 8
	        }
	    }
	}

* it should become:

	{
	    "key": 3,
	    "foo.a": 5,
	    "foo.bar.baz": 8
	}

* You can assume keys do not contain dots in them, i.e. no clobbering will occur.
*/

package dailycodingproblems

import (
	"testing"
)

type (
	flattenMapStrct struct{}
	flattenMapInf   interface {
		flattenMap(dict map[string]interface{}, parentKey string, sep string, result map[string]interface{})
	}
)

func (f flattenMapStrct) flattenMap(dict map[string]interface{}, parentKey string, sep string, result map[string]interface{}) {
	for key, value := range dict {
		var newKey string
		if parentKey == "" {
			newKey = key
		} else {
			newKey = parentKey + sep + key
		}

		if nestedMap, ok := value.(map[string]interface{}); ok {
			f.flattenMap(nestedMap, newKey, sep, result)
		} else {
			result[newKey] = value
		}
	}
}

func TestFlattenDict(t *testing.T) {
	// Example input
	nestedMap := map[string]interface{}{
		"key": 3,
		"foo": map[string]interface{}{
			"a": 5,
			"bar": map[string]interface{}{
				"baz": 8,
			},
		},
	}

	// Prepare result map and flatten
	var fmp flattenMapInf
	fmpS := flattenMapStrct{}
	fmp = fmpS

	flatMap := make(map[string]interface{})
	fmp.flattenMap(nestedMap, "", ".", flatMap)

	// Print the flattened map
	for k, v := range flatMap {
		t.Logf("%s: %v\n", k, v)
	}
}
