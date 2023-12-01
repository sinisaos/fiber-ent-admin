package model

type NewUserInput struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserInput struct {
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Superuser bool   `json:"superuser"`
}

type LoginUserInput struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type NewQuestionInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  int    `json:"author"`
	Tags    []int  `json:"tags"`
}

type UpdateQuestionInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  int    `json:"author"`
	Tags    []int  `json:"tags"`
}

type NewAnswerInput struct {
	Content  string `json:"content"`
	Author   int    `json:"author"`
	Question int    `json:"question"`
}

type UpdateAnswerInput struct {
	Content  string `json:"content"`
	Author   int    `json:"author"`
	Question int    `json:"question"`
}

type NewTagInput struct {
	Name string `json:"name"`
}

type UpdateTagInput struct {
	Name string `json:"name"`
}
