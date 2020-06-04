package main

import (
	"testing"
)

func TestGetStatus(t *testing.T) {
	type args struct {
		scrapeResult []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"result up", args{[]string{"github.com is up"}}, "up"},
		{"result down", args{[]string{"github.com is down"}}, "down"},
		{"result invalid domain", args{[]string{"We need a valid domain to check! Try again"}}, "invalidDomain"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStatus(tt.args.scrapeResult); got != tt.want {
				t.Errorf("getStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
