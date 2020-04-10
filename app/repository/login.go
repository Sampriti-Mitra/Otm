package repository

import (
	"otm/app/interfaces"
)

type Login struct {
	Base
}

func GetLoginRepo() interfaces.ILoginRepo {
	return &Login{}
}
