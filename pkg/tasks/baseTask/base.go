package dtasksbase

// Tasks supposed to go through stages. Number of stages is individual for each task
// At each stage task is supposed to interact with different entity

type ITask interface {
	String() string
	Stage() int
	NextStop() DTaskStop
}
