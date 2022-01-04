package form

type AddProductForm struct {

}


type AddUserAreaForm struct {
	AreaCode int64       `json:"areaCode" binding:"required"`
	AppId    int64       `json:"appId" binding:"required"`
	UserId   string      `json:"userId" binding:"required"`
	Address  string      `json:"address" binding:"required"`
	IsDefault int 		 `json:"isDefault" binding:"lte=1"`
	MergerName string
}