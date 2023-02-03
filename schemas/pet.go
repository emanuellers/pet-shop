package schemas

type Pet struct {
	Id    string `json:"id" xml:"id"`
	Name  string `json:"name" xml:"name"`
	Owner string `json:"owner" xml:"owner"`
}
