package models

type Admin struct {
	ID       string `json:"admin_id" bson:"_id"`
	Email    string `json:"email" validate:"contains=@rahmana.com" bson:"email"`
	Password string `json:"password" validate:"required,min=7,endswith=rahp" bson:"password"`
	Secret   string `json:"secret" bson:"secret"`
}
