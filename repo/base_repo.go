package repo

type BaseRepositoryAggregation interface {
	Count(result interface{}, groupBy string) error
	GroupBy(result interface{}, selecedBy string, whereBy map[string]interface{}, groupBy string) error
}

type BaseRepositoryPaging interface {
	Paging(page int, itemPerPage int) (interface{}, error)
}