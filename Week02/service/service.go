package service

import "studygolang/demo/dao"

func Service() error {
	return dao.Dao()
}