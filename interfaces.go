package sqload

type Dialector interface {
	Name() string
	Load() error
	Parse() error
}
