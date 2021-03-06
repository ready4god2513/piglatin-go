package piglatin

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

var translateTests = []struct {
  in string
  out string
}{
  {
    in: "dog",
    out: "ogday",
  },
  {
    in: "cat",
    out: "atcay",
  },
  {
    in: "hat",
    out: "athay",
  },
  {
    in: "egg",
    out: "eggday",
  },
}

func TestTranslate(t *testing.T){
  for i, test := range translateTests {
    actual := Translate(test.in)
    assert.Equal(t, actual, test.out, "Test %d", i)
  }
}