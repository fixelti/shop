package http

func (handler Handler) handlers() {
	v1 := handler.echo.Group("/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/signup", handler.user.Signup)
			user.POST("/login", handler.user.Login)
			user.POST("/refresh-access-token", handler.user.RefreshAccessToken, handler.VerifyRefreshToken)
		}
		product := v1.Group("/product")
		{
			product.POST("", handler.product.Create, handler.AdministratorCheck)
			product.GET("", handler.product.GetByID)
			product.GET("/list", handler.product.GetList)
			product.PUT("", handler.product.Update, handler.AdministratorCheck)
		}
	}
}
