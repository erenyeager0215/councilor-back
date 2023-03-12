package models

type NoRecordError struct {
	Massage string
}

func (e *NoRecordError) Error() string {
	return e.Massage
}

func RaiseError() error {
	return &NoRecordError{Massage: "議員・カテゴリのお気に入り登録がしていません"}
}
