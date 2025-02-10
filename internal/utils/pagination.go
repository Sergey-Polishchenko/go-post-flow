package utils

import (
	"context"
	"errors"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	flowerrors "github.com/Sergey-Polishchenko/go-post-flow/internal/errors"
)

func ApplyPagination[T any](data []*T, limit, offset *int) []*T {
	var start int
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
	ctx context.Context,
	comments []*model.Comment,
	maxDepth *int,
	expand bool,
	getChildrenIDs func(string) ([]string, error),
	getChildrenByIDs func(context.Context, []string) ([]*model.Comment, error),
) ([]*model.Comment, error) {
	result := make([]*model.Comment, 0, len(comments))

	var depth int
	if maxDepth != nil {
		depth = *maxDepth
	}

	for _, c := range comments {
		cloned := *c
		if depth > 0 && expand {
			childrenIDs, err := getChildrenIDs(c.ID)
			if err != nil {
				if errors.Is(err, flowerrors.ErrCommentChildrenNotFound) ||
					errors.Is(err, flowerrors.ErrCommentNotFound) {
					continue
				}
				return nil, err
			}

			children, err := getChildrenByIDs(ctx, childrenIDs)
			if err != nil {
				if errors.Is(err, flowerrors.ErrCommentChildrenNotFound) ||
					errors.Is(err, flowerrors.ErrCommentNotFound) {
					continue
				}
				return nil, err
			}

			childDepth := depth - 1
			clonedChildren, err := ProcessCommentsWithDepth(
				ctx,
				children,
				&childDepth,
				expand,
				getChildrenIDs,
				getChildrenByIDs,
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
