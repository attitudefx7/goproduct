package form

type AddAttributeForm struct {
	Name string `json:"name" binding:"required"`
}

type EditAttributeForm struct {
	ID   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type DeleteAttributeForm struct {
	ID int64 `form:"id" binding:"required"`
}

type GetAttributeForm struct {
	ID int64 `form:"id" binding:"required"`
}

type ListAttributeForm struct {
	ID   int64  `form:"id"`
	Name string `form:"name"`
}