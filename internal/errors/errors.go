package errors

import "fmt"

var (
	ErrUserNotFound            = fmt.Errorf("user not found")
	ErrAuthorNotFound          = fmt.Errorf("author not found")
	ErrPostNotFound            = fmt.Errorf("post not found")
	ErrCommentNotFound         = fmt.Errorf("comment not found")
	ErrParentCommentNotFound   = fmt.Errorf("parent comment not found")
	ErrCommentChildrenNotFound = fmt.Errorf("comment children not found")
	ErrCommentTooLong          = fmt.Errorf("comment too long")
	ErrCommentsNotAllowed      = fmt.Errorf("post not allows comments")

	ErrPingDatabase = fmt.Errorf("failed to ping database")

	ErrEnvParsing = fmt.Errorf("failed to parse environment variables")
)

type SQLQueryLoadingError struct {
	Value error
}

func (e *SQLQueryLoadingError) Error() string {
	return fmt.Sprintf("sql query not loading: %s", e.Value)
}

type SQLCreatingError struct {
	Value error
}

func (e *SQLCreatingError) Error() string {
	return fmt.Sprintf("sql not created: %s", e.Value)
}

type SQLScaningError struct {
	Value error
}

func (e *SQLScaningError) Error() string {
	return fmt.Sprintf("sql row not scaned: %s", e.Value)
}

type DatabaseConnectingError struct {
	Value error
}

func (e *DatabaseConnectingError) Error() string {
	return fmt.Sprintf("failed to connect to database: %s", e.Value)
}

type EnvLoadingError struct {
	Value error
}

func (e *EnvLoadingError) Error() string {
	return fmt.Sprintf("failed to load .env file: %s", e.Value)
}

// fmt.Errorf("failed to load .env file: %v", err)
