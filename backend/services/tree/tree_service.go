package tree

import "sovereignty/backend/models"

type TreeService struct {
	Trees map[uint]*models.Tree
}

func NewTreeService() *TreeService {
	return &TreeService{
		Trees: make(map[uint]*models.Tree),
	}
}
