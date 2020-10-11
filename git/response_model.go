package git

type GitResponseModel struct{
	Status    string
	Message   string
	HasFailed bool
}