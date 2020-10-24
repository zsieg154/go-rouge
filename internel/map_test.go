package internel

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("Generation of map", func(t *testing.T) {
		testMap := Map{}

		testMap.Generate(10,10)
		got := testMap.coordinates

		want := [10][10]string {
			{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
			{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
			{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
			{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
			{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
			{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
			{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
			{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
			{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
			{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
		}
		assertionHelper(t, got, want)
	})
}

func assertionHelper(t *testing.T, got [10][10]string, want [10][10]string) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v but wanted %v", got, want)
	}
}
