package convert

type CategoryModel struct {
	Result []CategoryBaseModel `json:"result"`
}

type CategoryBaseModel struct {
	Name     string              `json:"title"`
	Id       int                 `json:"category_id"`
	Children []CategoryBaseModel `json:"children"`
}

type Categories []Category

type Category struct {
	Name    string
	Parents []string
	Id      interface{}
}
