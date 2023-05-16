package helper

import (
	"fmt"
)

func GeneratePaginationQuery(page int, pageSize int) (query string) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	offset := pageSize * (page - 1)

	return fmt.Sprintf(`
		LIMIT %d 
		OFFSET %d
	`, pageSize, offset)
}
