package calcstats

import (
	"errors"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestMin(t *testing.T) {
	name_to_test_function := map[string]func(*testing.T){
		`Given:
		 When: Calling Min() with a slice of one element
		 Then: Min() should return that element as a result and nil as an error
		 `:
		func(t *testing.T) {
			e := 1
			result, err := Min([]int{e})
			if result != 1 || err != nil {
				t.Errorf("Tesing Min(). expect: %d, %v, got: %d, %v", e, nil, result, err)
			}
			rand.Seed(time.Now().UnixNano())
			e = rand.Int()
			result, err = Min([]int{e})
			if result != e || err != nil {
				t.Errorf("Tesing Min(). expect: %d, %v, got: %d, %v", e, nil, result, err)
			}
			e = rand.Int()
			result, err = Min([]int{e})
			if result != e || err != nil {
				t.Errorf("Tesing Min(). expect: %d, %v, got: %d, %v", e, nil, result, err)
			}
		},

		`Given:
		 When: Calling Min() with an empty slice
		 Then: Min() should return the minimum int value as a result and an error as an error
		 `:
		 func(t *testing.T) {
			 result, err := Min([]int{})
			 //expectederror := errors.New("This is an empty slice!")
			 if result != math.MinInt || !errors.Is(err, emptysliceerror) {
				t.Errorf("Tesing Min(). expect: %d, %v, got: %d, %v", math.MinInt, emptysliceerror, result, err)
			 }
		 },
	}
	
	for key, function := range name_to_test_function {
		t.Run(key, function)
	}
}