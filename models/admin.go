package models

type Admin struct {
	ID       string `json:"admin_id" bson:"_id"`
	Secret   string `json:"secret" bson:"secret"`
}
