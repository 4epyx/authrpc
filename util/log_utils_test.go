package util_test

import (
	"os"
	"testing"

	"github.com/4epyx/authrpc/util"
)

func TestGetTextFileLogger(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "smoke test",
			args: args{
				filepath: "../authrpc.log",
			},
			wantErr: false,
		},
		{
			name: "file already exist",
			args: args{
				filepath: "../authrpc.log",
			},
			wantErr: false,
		},
		{
			name: "permission must denied",
			args: args{
				filepath: "/root/logs/authrpc/test.log",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := util.GetTextFileLogger(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTextFileLogger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				os.Remove(tt.args.filepath)
			}
		})
	}
}
