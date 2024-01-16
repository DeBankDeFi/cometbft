package types

var disableABCIQueryMutex = true

const (
	FlagDisableCheckTx = "disable-checktx"
)

func GetDisableABCIQueryMutex() bool {
	return disableABCIQueryMutex
}

func SetDisableABCIQueryMutex(isClose bool) {
	disableABCIQueryMutex = isClose
}
