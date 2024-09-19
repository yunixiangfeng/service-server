package utils

// import "git.internal.attains.cn/attains-cloud/go-api/attains/common/api"

func ParsePaginationLimit(pagination *api.Pagination) int {
	pageSize := int(pagination.PageSize)
	if pageSize < 1 {
		pageSize = 10
	}
	return pageSize
}

func ParsePaginationOffset(pagination *api.Pagination) int {
	_pageIndex, _pageSize := int(pagination.PageIndex), ParsePaginationLimit(pagination)
	if _pageIndex < 1 {
		_pageIndex = 1
	}
	return (_pageIndex - 1) * _pageSize
}
