package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

func Dao() error  {
	err := sql.ErrNoRows
	return errors.Wrap(err,"this error is happen in dao")
}