package models

import (
	"fmt"
	"time"
)

type User struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Birthday    time.Time `json:"birthday"`
	SpouseID    int       `json:"spouse_id"`
	ChildrenIDs []int     `json:"children_ids"`
}

type Tree struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Nodes       []User `json:"nodes"`
}

func (t *Tree) PrintUsers() []string {
	stringUsers := []string{}
	if len(t.Nodes) == 0 {
		fmt.Println("No users in the tree.")
		return nil
	}

	fmt.Println("Users in the tree:")
	for _, user := range t.Nodes {
		formattedUser := fmt.Sprintf(
			"ID: %d, Name: %s, SpouseID: %d, Children: %v",
			user.ID, user.Name, user.SpouseID, user.ChildrenIDs,
		)
		stringUsers = append(stringUsers, formattedUser)
	}
	return stringUsers
}
