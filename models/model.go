package models

type Singers struct {
	SingerId  int64  `spanner:"SingerId`
	FirstName string `spanner:"FirstName"`
	LastName  string `spanner:"LastName"`
}
