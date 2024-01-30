package dtasks

type ITask interface {
	GetTaskId() DTaskId
	String() string
}
