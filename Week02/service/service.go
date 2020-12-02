package service

import "dao"

func Service() error {
	return dao.Dao()
}
