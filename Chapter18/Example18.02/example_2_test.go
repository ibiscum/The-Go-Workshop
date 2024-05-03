package command_injection

import (
	"reflect"
	"testing"
)

func TestListFiles(t *testing.T) {
	out, err := listFiles(" .; cat /etc/hosts")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(out)
	}
}

func Test_listFiles2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := listFiles2(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("listFiles2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("listFiles2() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
