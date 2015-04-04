package wordsearch

import (
	"reflect"
	"testing"
)

// Define a function Solve(words []string, puzzle []string) (map[string][2][2]int, error).
//
// Also define an exported TestVersion with a value that matches
// the internal testVersion here.

const testVersion = 1

var words = []string{
	"clojure", "ecmascript", "elixir", "go", "java", "lisp",
	"ocaml", "ruby", "rust", "scheme",
}

var puzzle = []string{
	"gefblpepre",
	"cbmdcimguc",
	"oikoknrjsm",
	"pbwjrqrota",
	"rixilelhgs",
	"woncqlispc",
	"schemekmgr",
	"alxhprubyi",
	"javaocamlp",
	"clojurermt",
}

// Top left corner is (0, 0)
// Entries are {{firstX, firstY}, {lastX, lastY}}.
var positions = map[string][2][2]int{
	"clojure":    {{0, 9}, {6, 9}},
	"ecmascript": {{9, 0}, {9, 9}},
	"elixir":     {{5, 4}, {0, 4}},
	"go":         {{8, 4}, {7, 3}},
	"java":       {{0, 8}, {3, 8}},
	"lisp":       {{5, 5}, {8, 5}},
	"ocaml":      {{4, 8}, {8, 8}},
	"ruby":       {{5, 7}, {8, 7}},
	"rust":       {{8, 0}, {8, 3}},
	"scheme":     {{0, 6}, {5, 6}},
}

func TestSolve(t *testing.T) {
	if TestVersion != testVersion {
		t.Fatalf("Found TestVersion = %v, want %v", TestVersion, testVersion)
	}
	actual, err := Solve(words, puzzle)
	if err != nil {
		t.Fatalf("Didn't expect error but got %v", err)
	}
	if !reflect.DeepEqual(actual, positions) {
		t.Fatalf("Got %v, want %v", actual, positions)
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve(words, puzzle)
	}
}
