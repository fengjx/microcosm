package utils

// GetOffsetLimit 根据分页参数获得offset和limit
func GetOffsetLimit(page int, size int) (offset int, limit int) {
	offset = (page - 1) * size
	limit = size
	return
}
