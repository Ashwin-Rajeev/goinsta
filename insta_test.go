package goinsta

import (
	"testing"

	"github.com/Ashwin-Rajeev/goerror"
)

func TestInfo_GetImage(t *testing.T) {
	tests := []struct {
		name      string
		get       func() error
		parseHTML func() error
		want      bool
		wantErr   bool
	}{
		{
			name: "normal case",
			get: func() error {
				return nil
			},
			parseHTML: func() error {
				return nil
			},
			want: true,
		},
		{
			name: "failure case(Error in get function)",
			get: func() error {
				return goerror.New("new error")
			},
			wantErr: true,
		},
		{
			name: "failure case(Error in parseHTML function)",
			get: func() error {
				return nil
			},
			parseHTML: func() error {
				return goerror.New("new error")
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Info{Crawler: &mockInsta{
				MockGet:       tt.get,
				MockParseHTML: tt.parseHTML,
			}}
			got, err := i.GetImage()
			if (err != nil) != tt.wantErr {
				t.Errorf("Info.GetImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Info.GetImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
