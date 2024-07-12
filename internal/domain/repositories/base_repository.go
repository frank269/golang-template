package repositories

type Operator string

const (
	EqualOperator              Operator = "="
	NotEqualOperator           Operator = "<>"
	GreaterThanOperator        Operator = ">"
	GreaterThanOrEqualOperator Operator = ">="
	LessThanOperator           Operator = "<"
	LessThanOrEqualOperator    Operator = "<="
	LikeOperator               Operator = "like"
	InOperator                 Operator = "IN"
	NotInOperator              Operator = "NOT IN"
)

type FilterQuery struct {
	Field    string
	Operator Operator
	Value    interface{}
}

type BaseRepository[T any] interface {
	FindAll(offset, limit int, sortBy, sortOrder string) ([]T, error)
	FilterByQueries(queries []FilterQuery, offset, limit int, sortBy, sortOrder string) ([]T, error)
	FindByID(id string) (*T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(id string) error
	Count(queries []FilterQuery) (int64, error)
}
