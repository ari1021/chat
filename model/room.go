package model

type Group struct {
	ID      int
	Name    string
	Members []*User
}

var Groups = map[int]*Group{}
