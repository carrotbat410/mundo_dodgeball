package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AuthRequired(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	// JWT 토큰 검증
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 토큰 서명 메소드가 올바른지 확인
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	// 토큰 유효성 검사
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//AuthRequired 미들웨어에는 토큰을 복호화하여 id, username을 요청 컨텍스트에 저장한다.
		c.Locals("id", claims["id"])
		c.Locals("username", claims["username"])
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}

	// 토큰이 유효하다면 다음 핸들러로 이동
	return c.Next()
}
