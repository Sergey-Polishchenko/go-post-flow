package utils

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
