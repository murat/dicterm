package main

import (
	"fmt"
	"os"
	"testing"
)

func deleteTestConf(f string) {
	err := os.Remove(f)
	if err != nil {
		fmt.Printf("could not delete file, %v\n", err)
	}
}

func Test_run(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "returns err with missing arguments",
			args: args{
				args: []string{},
			},
			wantErr: true,
		},
		{
			name: "returns err if could not write to file",
			args: args{
				args: []string{"", "-key", "xxx", "-config", "test.conf", "-word", "hello"},
			},
			wantErr: true,
		},
		{
			name: "returns err if could not read from file",
			args: args{
				args: []string{"", "-config", "test.conf", "-word", "hello"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := run(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
			}

			// delete not_exists.conf after test
			for i := 0; i < len(tt.args.args); i++ {
				if tt.args.args[i] == "-config" {
					deleteTestConf(tt.args.args[i+1])
					break
				}
			}
		})
	}
}
