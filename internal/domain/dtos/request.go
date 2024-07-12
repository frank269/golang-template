package dtos

type Pagination struct {
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
}

type GetUserRequest struct {
	Pagination
	UserName  string `json:"username"`
	SortBy    string `json:"sortBy"`
	SortOrder string `json:"sortOrder"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
