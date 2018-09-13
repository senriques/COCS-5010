// Shawn Enriques
// senriqu1@uwyo.edu
// Assignment 1.6 Test a simple library package

package test1

import "testing"

func TestDouble(t *testing.T)	{

  tests := []struct {
    in              int
    expected        int
    }{
      {
        in:         23,
        expected:   46,
      },
      {
        in:	        1,
        expected:	  2,
      },
    }

  for ii, test := range tests {
    rr := DoubleValue(test.in)
    if rr != test.expected  {
      t.Errorf("DoubleValue Test %d, expected %d got %d\n", ii, test.expected, rr)
    }
  }
}

func TestTriple(t *testing.T)	{

  tests := []struct {
    in              int
    expected        int
    }{
      {
        in:         23,
        expected:   69,
      },
      {
        in:	        1,
        expected:	  3,
      },
    }

  for ii, test := range tests {
    rr := TripleValue(test.in)
    if rr != test.expected  {
      t.Errorf("TripleValue Test %d, expected %d got %d\n", ii, test.expected, rr)
    }
  }
}
