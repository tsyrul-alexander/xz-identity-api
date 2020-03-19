package condition

type LogicalOperation int

const (
	And = LogicalOperation(0)
	Or = LogicalOperation(1)
)