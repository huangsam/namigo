package main

import (
	"context"
	"testing"

	"github.com/urfave/cli/v3"
)

func TestCheckSizeFlag(t *testing.T) {
	tests := []struct {
		name    string
		value   int
		wantErr bool
	}{
		{
			name:    "valid size",
			value:   5,
			wantErr: false,
		},
		{
			name:    "invalid size zero",
			value:   0,
			wantErr: true,
		},
		{
			name:    "invalid size negative",
			value:   -1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkSizeFlag(context.Background(), &cli.Command{}, tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkSizeFlag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckLengthFlag(t *testing.T) {
	tests := []struct {
		name    string
		value   int
		wantErr bool
	}{
		{
			name:    "valid length",
			value:   10,
			wantErr: false,
		},
		{
			name:    "invalid length zero",
			value:   0,
			wantErr: true,
		},
		{
			name:    "invalid length negative",
			value:   -5,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkLengthFlag(context.Background(), &cli.Command{}, tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkLengthFlag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
