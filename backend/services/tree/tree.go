package tree

import (
	"sovereignty/backend/models"
	"time"
)

func getNextID() int {
	return 1 // TO Implement
}

func createTemplateUsers() []models.User {
	names := []string{"GrandPa1", "GrandMa1", "GrandPa2", "GrandMa2", "Parent1", "Parent2", "Child"}
	spousesID := []int{2, 1, 4, 3, 6, 5, 0}
	childrenID := [][]int{{5}, {5}, {6}, {6}, {7}, {7}, {0}}
	users := make([]models.User, 0, len(names))

	for i := 0; i < 5; i++ {
		users = append(users, models.User{ID: i, Name: names[i], Birthday: time.Now(),
			SpouseID: spousesID[i], ChildrenIDs: childrenID[i]})
	}

	return users
}

func initData(name, description string) models.Tree {
	// Check ID
	return models.Tree{
		ID:          getNextID(),
		Name:        name,
		Description: description,
		Nodes:       createTemplateUsers(),
	}
}

func NewTree(name, description string) models.Tree {
	tree := initData(name, description)
	tree.PrintUsers()
	return tree
}
