package tests

import (
	"github.com/neel1996/gitconvex-server/api"
	"testing"
)

func TestUpdateRepoName(t *testing.T) {
	type args struct {
		repoId   string
		repoName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := api.UpdateRepoName(tt.args.repoId, tt.args.repoName)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateRepoName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateRepoName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
