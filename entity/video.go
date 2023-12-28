package entity

type Persons struct {
	FisrtName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}

type Video struct {
	Tittle      string  `json:"tittle" binding:"min=2,max=10"`
	Description string  `json:"description" binding:"max=20"`
	URL         string  `json:"url" binding:"required,url"`
	Author      Persons `json:"author" binding:"required"`
}
