package main

import "testing"

func Test_hello(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Index 0",
			args:    args{0},
			want:    "Hello, world",
			wantErr: false,
		},
		{
			name:    "Index 1",
			args:    args{1},
			want:    "Καλημέρα κόσμε",
			wantErr: false,
		},
		{
			name:    "Index out of range",
			args:    args{5},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Index out of range",
			args:    args{-1},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hello(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("hello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %v, wanted %v", got, tt.want)
			}
		})
	}
}
