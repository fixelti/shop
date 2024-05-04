package http

func (handler Handler) handlers() {
	v1 := handler.echo.Group("/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/signup", handler.user.Signup)
		}
	}
}
