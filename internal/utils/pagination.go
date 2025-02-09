package utils

import (
	"errors"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	reperrors "github.com/Sergey-Polishchenko/go-post-flow/internal/repository/errors"
)

func ApplyPagination[T any](data []*T, limit, offset *int) []*T {
	start := 0
	if offset != nil {
		start = *offset
	}

	end := len(data)
	if limit != nil {
		end = start + *limit
		if end > len(data) {
			end = len(data)
		}
	}

	if start > len(data) {
		return []*T{}
	}

	return data[start:end]
}

func ProcessCommentsWithDepth(
	comments []*model.Comment,
	maxDepth *int,
	getChildren func(string) ([]*model.Comment, error),
	expand bool,
) ([]*model.Comment, error) {
	result := make([]*model.Comment, 0, len(comments))

	currentDepth := 0
	if maxDepth != nil {
		currentDepth = *maxDepth
	}

	for _, c := range comments {
		cloned := *c
		if currentDepth > 0 && expand {
			children, err := getChildren(c.ID)
			if err != nil {
				if errors.Is(err, reperrors.ErrCommentChildrenNotFound) ||
					errors.Is(err, reperrors.ErrCommentNotFound) {
					continue
				}
				return nil, err
			}

			childDepth := currentDepth - 1
			clonedChildren, err := ProcessCommentsWithDepth(
				children,
				&childDepth,
				getChildren,
				expand,
			)
			if err != nil {
				return nil, err
			}
			cloned.Children = clonedChildren
		} else {
			cloned.Children = []*model.Comment{}
		}
		result = append(result, &cloned)
	}

	return result, nil
}
