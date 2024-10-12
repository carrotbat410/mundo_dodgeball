package controllers

import (
	"fiber_prac/models"
	"fiber_prac/services"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// 회원가입 핸들러 함수
func RegisterUser(c *fiber.Ctx) error {
	// 요청 바디에서 유저 정보 파싱
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// 유저 등록 서비스 호출
	err := services.RegisterUser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// 성공적으로 회원가입 완료
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

// api := api.Group("/v1", middleware) 필요시 middleware + c.Next() 사용 가능

// app.Get("/", func(c *fiber.Ctx) error {
// 	return c.SendString("hello world")
// })

// app.Get("/sample/error", func(c *fiber.Ctx) error {
// 	return fiber.NewError(782, "Custom Error Message")
// })

func Login(c *fiber.Ctx) error {

	var input models.LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	// validation
	if input.Id == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	userData, loginErr := services.Login(input.Id, input.Password)

	if loginErr != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Not Found",
			"error":   loginErr.Error(),
		})
	}

	token, err := generateJWT(userData.Id, userData.Username)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate token",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"token":   token,
	})
}

func generateJWT(id, username string) (string, error) {
	// JWT 클레임 설정 (서명된 데이터)
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // 유효 시간: 24시간 //TODO 개발단계니 72시간으로
	}

	// 비밀 키로 서명
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}
