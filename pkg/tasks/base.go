package dtasks

// Tasks supposed to go through stages. Number of stages is individual for each task

type BaseTask interface {
	String() string
	GetStage() int
	IsDone() bool
}
