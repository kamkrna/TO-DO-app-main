package middleware

import (
	"log"
	"my-app-server/server/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        tokenstring := c.Request().Header.Get("Authorization")
        log.Printf("Token String: "+ tokenstring)
        if len(tokenstring) > 7 && tokenstring[:7] == "Bearer " {
            tokenstring = tokenstring[7:]
        }
        if tokenstring == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "Token is required")
        }
        claims, err := auth.VerifyToken(tokenstring)
        if err != nil {
            log.Printf("Error verifying token: %s", err)
            return c.JSON(http.StatusUnauthorized, "Invalid token!")
        }
        log.Printf("Claims: Username: %s", claims.Username)
        c.Set("username", claims.Username)
        return next(c)
    }
}