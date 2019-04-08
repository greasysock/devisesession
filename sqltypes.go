package devisesession

type ServerType int

const (
	postgresql ServerType = iota
	sqlite
)