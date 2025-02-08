package reperrors

import "fmt"

var (
	ErrPostNotFound            = fmt.Errorf("post not found")
	ErrAuthorNotFound          = fmt.Errorf("author not found")
	ErrCommentNotFound         = fmt.Errorf("comment not found")
	ErrParentCommentNotFound   = fmt.Errorf("parent comment not found")
	ErrCommentChildrenNotFound = fmt.Errorf("comment children not found")
)
