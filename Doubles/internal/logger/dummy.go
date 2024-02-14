package logger

// Dummy is a dummy logger.
func NewDummy() *Dummy {
	return &Dummy{}
}

type Dummy struct{}

// this implementation is empty because it is a dummy logger
func (d *Dummy) Logf(format string, args ...interface{}) {
}

func (d *Dummy) Warnf(format string, args ...interface{}) {
}
