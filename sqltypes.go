package devisesession

type ServerType int

const (
	Postgresql ServerType = iota
	Sqlite
)