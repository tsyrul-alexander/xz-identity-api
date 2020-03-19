package condition

type QueryCondition interface {
	getType() Type
}