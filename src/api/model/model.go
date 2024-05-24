package model

type GetData struct {
	Name     string `json:"name" bson:"name" binding:"required"`
	Email    string `json:"email" bson:"email" binding:"required,email"`
	Password string `json:"password" bson:"password" binding:"required"`
	Age      int    `json:"age" bson:"age" binding:"required,gte=0,lte=100"`
	MobileNo int    `json:"mobileNo" bson:"mobileNo" binding:"required"`
}
