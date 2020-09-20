package enum

//ErrorType specifies possible error types
type ErrorType int

const (
	ErrorEasy   ErrorType = 1 // Easy (user error)
	ErrorMedium ErrorType = 2 // Medium (file, code error)
	ErrorHard   ErrorType = 3 // Critical (code/system/runtime error)
)
