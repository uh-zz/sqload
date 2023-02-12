package sqload

type Dialector interface {
	Name() string
	Parse(string, *[]string) error
}
