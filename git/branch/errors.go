package branch

type Error struct {
	error
	ErrorCode   string
	ErrorString string
}

var (
	NilRepoError         = Error{ErrorCode: "REPO_NIL_ERROR", ErrorString: "Repo is nil"}
	EmptyBranchNameError = Error{ErrorCode: "BRANCH_EMPTY_ERROR", ErrorString: "Branch name(s) is empty"}
)

func (e Error) Error() string {
	return e.ErrorString
}
