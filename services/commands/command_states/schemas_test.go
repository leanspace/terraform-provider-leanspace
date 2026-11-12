package command_states

import "testing"

func TestValidateStateName(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "TERRAFORM_STATE"},
		{name: "A"},
		{name: "A_B"},
		{name: "DRAFT", wantErr: true},
		{name: "STAGED", wantErr: true},
		{name: "QUEUED", wantErr: true},
		{name: "POPPED", wantErr: true},
		{name: "ABORTED", wantErr: true},
		{name: "lowercase", wantErr: true},
		{name: "A_", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, errs := validateStateName(tt.name, "name")
			if gotErr := len(errs) > 0; gotErr != tt.wantErr {
				t.Errorf("validateStateName(%q) returned errors %v; want error = %t", tt.name, errs, tt.wantErr)
			}
		})
	}
}
