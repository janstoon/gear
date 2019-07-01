package telemetry

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	mcTest := mockConn(1)
	result , err := mcTest.Get("1")

	if result == "ok" {
		fmt.Println("Everything is OK")
	}else{
		t.Errorf("Something is wrong with error: %v", err)
	}
}
