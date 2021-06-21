package internal

import (
	"reflect"
	"testing"
)

func TestCalcAddressByToken(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		want    string
		wantErr bool
	}{
		{"kumaon", "7307fb5d7c33a3a51968eb124ce2004ce2d293f22c26682661bed4011b74322c", "cosmos1t2xr07qvvfrywxzcw47akx44z745yhqz2lc52y", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calcAddressByToken(tt.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("calcAddressByToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcAddressByToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
