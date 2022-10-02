package product

type SomethingWithID struct {
	ID int `uri:"id" binding:"required"`
}