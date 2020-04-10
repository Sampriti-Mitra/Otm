package repository

import "otm/app/interfaces"

type Profile struct {
	Base
}

func GetProfileRepo() interfaces.IProfileRepo {
	return &Profile{}
}
