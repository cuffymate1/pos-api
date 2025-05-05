package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/golang-jwt/jwt/v5"
)

// SecurityHeaders middleware adds security headers to all responses
func SecurityHeaders() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Add security headers
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Frame-Options", "DENY")
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("Content-Security-Policy", "default-src 'self'")
		c.Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Set("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Set("Permissions-Policy", "geolocation=(), microphone=(), camera=()")

		return c.Next()
	}
}

// RateLimiter middleware limits repeated requests
func RateLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        20,              // Maximum number of requests
		Expiration: 1 * time.Minute, // Expiration time
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Use IP address as key
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Too many requests, please try again later",
			})
		},
	})
}

// CORSConfig middleware configures CORS
func CORSConfig() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // เปลี่ยนเป็น domain ที่ต้องการ
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
}

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
