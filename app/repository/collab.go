package repository

import "otm/app/interfaces"

type Collab struct {
	Base
}

func GetCollabRepo() interfaces.ICollabRepo {
	return &Collab{}
}
