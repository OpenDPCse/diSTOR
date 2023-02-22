package err

import "errors"

var (
	DirExists                   = errors.New("dir already exists")
	ErrFileNotFound             = errors.New("file not found")
	NothingToStore              = errors.New("there is nothing to be stored")
	IllegalSeed                 = errors.New("the RandomSeed is not a positive number")
	IllFormedIP                 = errors.New("the format of ipAddr is wrong")
	ConfFailed                  = errors.New("configure failed")
	JoinNetworkFailed           = errors.New("join network failed")
	NodeNotStarted              = errors.New("node not started")
	KeyTableDistributeException = errors.New("key table distributed exception")
	KeyTableUpdateErr           = errors.New("key table update failed")
)
