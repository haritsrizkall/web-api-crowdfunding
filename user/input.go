package user

type RegisterUserInput struct {
	Name       string `name:"name" binding:"required"`
	Occupation string `name:"occupation" binding:"required"`
	Email      string `name:"email" binding:"required,email"`
	Password   string `name:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `name:"email" binding:"required,email"`
	Password string `name:"password" binding:"required"`
}

type EmailCheckInput struct {
	Email string `name:"email" binding:"required,email"`
}
