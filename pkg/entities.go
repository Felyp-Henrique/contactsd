package pkg

type Contact struct {
	Id        string `bson:"id" json:"id"`
	Name      string `bson:"name" json:"name"`
	Email     string `bson:"email" json:"email"`
	Cellphone string `bson:"cellphone" json:"cellphone"`
}
