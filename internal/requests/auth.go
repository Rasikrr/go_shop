package requests

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUp struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password1 string `json:"password1"`
	Password2 string `json:"password2"`
}

type EditProfile struct {
	FirstName string `bson:"firstName" form:"firstName"`
	LastName  string `bson:"lastName" form:"lastName"`
	MobileNum string `bson:"telNumber" form:"telNum"`
	Country   string `bson:"country" form:"country"`
	State     string `bson:"state" form:"state"`
	City      string `bson:"city" form:"city"`
	Address   string `bson:"address" form:"address"`
	PostCode  string `bson:"postCode" form:"postCode"`
	PhotoPath string `bson:"photoPath,omitempty" json:"PhotoPath,omitempty"`
}
