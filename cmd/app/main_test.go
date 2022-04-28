package main

import "testing"

func Test_run(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "not implemented yet",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := run(); (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
