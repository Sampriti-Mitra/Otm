package repository

import "otm/app/interfaces"

type About struct {
	Base
}

func GetAboutRepo() interfaces.IProfileRepo {
	return &About{}
}
