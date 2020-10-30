package model

type ResponseModel struct {
	Status    string
	Message   string
	HasFailed bool
}

type NewRepoInputs struct {
	RepoName    string
	RepoPath    string
	CloneSwitch bool
	RepoURL     *string
	InitSwitch  bool
	AuthOption  string
	UserName    *string
	Password    *string
}
