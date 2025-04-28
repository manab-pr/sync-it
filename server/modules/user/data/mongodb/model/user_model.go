package model

type UserModel struct {
	Id        string `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	Email     string `json:"email unique" bson:"email"`
	Phone	 string `json:"phone unique" bson:"phone"`
	PasswordHash string `json:"password_hash" bson:"password_hash"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
}