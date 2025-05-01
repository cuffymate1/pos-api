package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("Jwt")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: No token provided",
			})
		}
		// Parse and validate the JWT token
		claims := jwt.MapClaims{}
		parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			// Verify signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Replace this with your actual secret key
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Invalid token",
			})
		}

		if !parsedToken.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Invalid token",
			})
		}

		// Store the claims in the context for use in subsequent handlers
		c.Locals("jwt", claims)
		fmt.Print(claims)
		// Continue to the next middleware/handler
		return c.Next()
	}
}

func OnlyAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// ดึง claims ที่ได้จาก VerifyToken
		claimsRaw := c.Locals("jwt")
		if claimsRaw == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: No claims found",
			})
		}

		claims, ok := claimsRaw.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Invalid claims",
			})
		}

		role, ok := claims["role"].(string)
		if !ok || role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden: Admins only",
			})
		}

		return c.Next()
	}
}
