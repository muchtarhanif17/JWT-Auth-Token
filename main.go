package main

import (
	// "belajargolang/middleware"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID   int
	Name string
}

type (
	LoginPayload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	JwtTokenPayload struct {
		Username string `json:"username"`
		jwt.RegisteredClaims
	}

	Response struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}
)

var SECRET_KEY = []byte("secretkey")

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtTokenPayload)
		},
		SigningKey: SECRET_KEY,
	})
}

func main() {

	router := echo.New()
	router.POST("api/token", func(c echo.Context) error {
		LoginPayload := new(LoginPayload)
		fmt.Print(LoginPayload.Username)
		if err := c.Bind(LoginPayload); err != nil {
			return err
		}

		if LoginPayload.Username == "admin" || LoginPayload.Password == "admin" {
			var token JwtTokenPayload
			lifeTimeToken := 1

			now := time.Now()
			token.RegisteredClaims = jwt.RegisteredClaims{
				Issuer:    "my-App",
				IssuedAt:  jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * time.Duration(lifeTimeToken))),
			}

			token.Username = LoginPayload.Username
			_token := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
			tokenResult, _ := _token.SignedString(SECRET_KEY)

			return c.JSON(http.StatusOK, Response{
				Status: 200,
				Data: struct {
					AccessToken string `json:"access_token"`
				}{
					AccessToken: tokenResult,
				},
			})
		}
		return nil
	})

	router.GET("api/profile", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Response{
			Status:  200,
			Message: "Anda telah mengakses halaman profile dengan API Token yang Valid",
		})
	}, JWTMiddleware())

	router.Start(":8080")
	// err := utils.InitDB()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Database connected:", utils.Conn)

	// GetDataDonatur()

	// r := gin.Default()

	// GET endpoint

	// r.Run(":8080") // start server

	// mux := http.NewServeMux()

	// mux.HandleFunc("/login", controller.Login)
	// mux.Handle("/profile", middleware.JWTMiddleware(http.HandlerFunc(controller.Profile)))

	// http.ListenAndServe(":8080", mux)
}
