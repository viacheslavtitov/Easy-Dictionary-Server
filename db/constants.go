package db

type OrderByType struct {
	VALUE string
}

var (
	OrderByASC  = OrderByType{VALUE: `ASC`}
	OrderByDESC = OrderByType{VALUE: `DESC`}
)
