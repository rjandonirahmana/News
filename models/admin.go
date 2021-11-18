package models

type Admin struct {
	ID       string `bson:"id" json:"admin_id"`
	Email    string `bson:"email" json:"email" validate:"contains=@rahmana.com"`
	Password string `bson:"password" json:"password" validate:"required,min=7,endswith=rahp"`
	Secret   string `bson:"secret" json:"secret"`
}
