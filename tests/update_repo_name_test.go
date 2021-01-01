package tests

import (
	"github.com/neel1996/gitconvex-server/api"
	"github.com/neel1996/gitconvex-server/utils"
	"testing"
)

func TestUpdateRepoName(t *testing.T) {
	utils.DataFileWriter([]utils.RepoData{{RepoId: "test", RepoName: "test", RepoPath: "", TimeStamp: ""}})

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
		{name: "test script for rename repo API", args: struct {
			repoId   string
			repoName string
		}{repoId: "test", repoName: "NewTest"}, want: "Repo name updated successfully", wantErr: false},
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
