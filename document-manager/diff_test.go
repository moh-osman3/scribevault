package diff

import (
   "testing"
   "github.com/stretchr/testify/assert"
)

func TestDiffLines(t *testing.T) {
   s1 := "I\nam\na\nman\n"
   s2 := "The\ndog\nis\nman\n"
   
   plus, minus := diffLines(s1, s2)

   expectedMinus := []Line{
       {linenumber: 0, text: "I"},
       {linenumber: 1, text: "am"},
       {linenumber: 2, text: "a"},
   }

   expectedPlus := []Line{
       {linenumber: 0, text: "The"},
       {linenumber: 1, text: "dog"},
       {linenumber: 2, text: "is"},
   }

   assert.Equal(t, len(expectedMinus), len(minus))
   assert.Equal(t, len(expectedPlus), len(plus))

   for i := range expectedMinus {
       assert.Equal(t, expectedMinus[i].linenumber, minus[i].linenumber)
       assert.Equal(t, expectedMinus[i].text, minus[i].text)
   }

   for i := range expectedPlus {
       assert.Equal(t, expectedPlus[i].linenumber, plus[i].linenumber)
       assert.Equal(t, expectedPlus[i].text, plus[i].text)
   }
}