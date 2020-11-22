package distributedTaskExecuter

// Task : Creating an Task interface
type Task interface {
	// Contains execute defintion
	Execute() interface{}
}
