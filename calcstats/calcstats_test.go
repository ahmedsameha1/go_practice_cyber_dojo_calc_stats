package calcstats

import "testing"

func TestAll(t *testing.T) {
	a := []int{1}
	name_to_test_function := map[string]func(*testing.T){"test1":
		func(t *testing.T) {
			result := Min(a)
			if result != 1 {
				t.Errorf("Tesing Min(). expect: %d, got: %d", 1, result)
			}
		},
	}
	
	for key, function := range name_to_test_function {
		t.Run(key, function)
	}
}