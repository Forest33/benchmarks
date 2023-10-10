package entity

type User struct {
	ID        uint64  `json:"a"`
	FirstName string  `json:"b"`
	LastName  string  `json:"c"`
	Active    bool    `json:"d"`
	Flags     []int32 `json:"e"`
}
