package utils

func ApplyPagination[T any](data []*T, limit, offset *int) []*T {
	lenD := len(data)
	var start int
	if offset != nil {
		start = *offset
	}

	if start > lenD {
		return []*T{}
	}

	end := lenD
	if limit != nil {
		end = start + *limit
		if end > lenD {
			end = lenD
		}
	}

	return data[start:end]
}
