package devisesession

import (
	"errors"
)

type SqlUser struct {
	username string
	password string
}

// Structure for keeping track of sql session for all sql queries for the session
type SqlConnection struct {
	address string
	port int16
	sql_user SqlUser
	db_name string
	auth_table_name string
	db_type ServerType
}

func NewSqlUser(username string, password string) (user SqlUser, err error){
	var u SqlUser
	// TODO: implement some validation
	u.username = username
	u.password = password
	return u, nil
}

func NewSqlConnection(address string, port int16, db_name string, sql_user SqlUser, auth_table_name string, db_type ServerType) (con SqlConnection, err error){
	var c SqlConnection
	switch db_type{
	case Sqlite:
		return c, errors.New("sqlite not supported")
	}
	c.address = address
	c.port = port
	c.db_name = db_name
	c.db_type = db_type
	c.auth_table_name = auth_table_name

	return c, nil
}