package commit

type Error struct {
	error
	ErrorCode   string
	ErrorString string
}

var (
	ChangeError              = Error{ErrorCode: "COMMIT_FAILED", ErrorString: "Failed to commit the new changes"}
	LogsError                = Error{ErrorCode: "LIST_COMMIT_LOGS_FAILED", ErrorString: "Failed to list the commit logs for the repo"}
	FileHistoryError         = Error{ErrorCode: "LISTING_COMMIT_FILE_HISTORY_FAILED", ErrorString: "Failed to get the list of file changes for the commit"}
	FileHistoryNoParentError = Error{ErrorCode: "LISTING_COMMIT_FILE_HISTORY_NO_PARENT", ErrorString: "The HEAD commit is the only commit in the repo and has no previous histories"}
	FileHistoryTreeError     = Error{ErrorCode: "LISTING_COMMIT_FILE_HISTORY_INVALID_TREE", ErrorString: "The commit tree is invalid"}
)

func (e Error) Error() string {
	return e.ErrorString
}
