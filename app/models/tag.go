package models

// Tag Tag
type Tag struct {
	// Model

	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

// GetTags a
func GetTags(page int, pageSize int) (tags []Tag) {
	// db.C("user").Find(nil).All(&tags)
	db := NewSessionStore()
	defer db.Close()
	db.C("user").Find(nil).All(&tags)
	return
}
