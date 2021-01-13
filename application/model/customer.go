package model

type Customer struct {
	Document string  `bson:"document"`
	Name     string  `bson:"name"`
	Role     string  `bson:"role"`
	Salary   float64 `bson:"salary"`
}
