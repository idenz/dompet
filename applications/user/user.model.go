package users

type UserModel struct {
	username string `json:"username" bson:"user_name"`
	email    string `json:"email" bson:"user_email"`
	password string `json:"password" bson:"user_password"`
}
