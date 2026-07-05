package model

type ExternalUser struct {
	ID       int64           `json:"id"`
	Name     string          `json:"name"`
	Username string          `json:"username"`
	Email    string          `json:"email"`
	Address  ExternalAddress `json:"address"`
	Phone    string          `json:"phone"`
	Website  string          `json:"website"`
	Company  ExternalCompany `json:"company"`
}

type ExternalAddress struct {
	Street  string      `json:"street"`
	Suite   string      `json:"suite"`
	City    string      `json:"city"`
	Zipcode string      `json:"zipcode"`
	Geo     ExternalGeo `json:"geo"`
}

type ExternalGeo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type ExternalCompany struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	BS          string `json:"bs"`
}
type UserResponse struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	City     string `json:"city"`
	Company  string `json:"company"`
}
