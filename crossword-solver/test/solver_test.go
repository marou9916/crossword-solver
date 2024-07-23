package test

import "testing"

func TestExample(t *testing.T) {
    // Example test
    if 1+1 != 2 {
        t.Errorf("1+1 should equal 2")
    }
}

