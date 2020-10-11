package api

import (
	"github.com/neel1996/gitconvex-server/graph/model"
	"reflect"
	"testing"
)

func TestAddRepo(t *testing.T) {
	type args struct {
		repoName    string
		repoPath    string
		cloneSwitch bool
		repoURL     *string
		initSwitch  bool
	}
	tests := []struct {
		name string
		args args
		want *model.AddRepoParams
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddRepo(tt.args.repoName, tt.args.repoPath, tt.args.cloneSwitch, tt.args.repoURL, tt.args.initSwitch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cloneHandler(t *testing.T) {
	type args struct {
		repoPath string
		repoURL  string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_initRepoHandler(t *testing.T) {
	type args struct {
		repoPath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_repoDataFileWriter(t *testing.T) {
	type args struct {
		repoId   string
		repoName string
		repoPath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_repoIdGenerator(t *testing.T) {
	type args struct {
		c chan string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
