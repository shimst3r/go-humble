package gohumble

import (
	"sync"
	"testing"
)

func Test_updateCategory(t *testing.T) {
	type args struct {
		wg       *sync.WaitGroup
		category string
	}
	tests := []struct {
		name string
		args args
	}{
		{"games", args{wg: &sync.WaitGroup{}, category: "games"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.wg.Add(1)
			updateCategory(tt.args.wg, tt.args.category)
		})
	}
}
