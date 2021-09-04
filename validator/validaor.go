package validator

type Validator interface {
	Validate() error
}

type ValidatorWithStringFields interface {
	ValidateWithFields(fields ...string) error
}
