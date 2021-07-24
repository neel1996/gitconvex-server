package commit

type Error struct {
	error
	ErrorCode   string
	ErrorString string
}

var (
	CommitChangeError              = Error{ErrorCode: "COMMIT_FAILED", ErrorString: "Failed to delete add new remote"}
	CommitLogsError                = Error{ErrorCode: "LIST_COMMIT_LOGS_FAILED", ErrorString: "Failed to delete add new remote"}
	CommitFileHistoryError         = Error{ErrorCode: "LISTING_COMMIT_FILE_HISTORY_FAILED", ErrorString: "Failed to get the list of file changes for the commit"}
	CommitFileHistoryNoParentError = Error{ErrorCode: "LISTING_COMMIT_FILE_HISTORY_NO_PARENT", ErrorString: "The HEAD commit is the only commit in the repo and has no previous histories"}
	CommitFileHistoryTreeError     = Error{ErrorCode: "LISTING_COMMIT_FILE_HISTORY_INVALID_TREE", ErrorString: "The commit tree is invalid"}
)

func (e Error) Error() string {
	return e.ErrorString
}
