package simple

import "fmt"

type Connention struct {
	File *File
}

func (c *Connention) Close() {
	fmt.Println("Close Connection", c.File.Name)
}

// Cleanup Function
func NewConnection(file *File) (*Connention, func()) {
	connection := &Connention{File: file}
	return connection, func() {
		connection.Close()
	}
}
