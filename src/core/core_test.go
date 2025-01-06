package core_test

import (
	"redir/src/core"
	"testing"

	"github.com/google/uuid"
)

func TestAddToRamMemory(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		params  map[string]interface{}
		memory  *map[uuid.UUID]core.CoreModel
		model   core.CoreModel
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := core.AddToRamMemory(tt.params, tt.memory, tt.model)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("AddToRamMemory() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("AddToRamMemory() succeeded unexpectedly")
			}
		})
	}
}
