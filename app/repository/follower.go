package repository

import "otm/app/interfaces"

type follower struct {
	Base
}

func GetFollowerRepo() interfaces.IFollowerRepo {
	return &follower{}
}
