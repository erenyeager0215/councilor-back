package models

type NoRecordError struct {
	Massage string
}

func (e *NoRecordError) Error() string {
	return e.Massage
}

func RaiseError() error {
	return &NoRecordError{Massage: "支持する議員を登録していません"}
}