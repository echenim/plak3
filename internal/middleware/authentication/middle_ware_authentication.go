package authentication

import (
	"strings"

	"github.com/valyala/fasthttp"
)

var jwtKey = []byte("aa3d118b-d505-482c-85c4-3cfa47d1ef45")

func AuthenticationMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		tokenString, err := extractBearerToken(ctx)
		if err != nil {
			sendError(ctx, err.Error(), fasthttp.StatusUnauthorized)
			return
		}

		claims, err := VerifyToken(tokenString)
		if err != nil {
			sendError(ctx, "Invalid or expired token", fasthttp.StatusUnauthorized)
			return
		}

		ctx.SetUserValue(userIDKey, claims.UserId)
		next(ctx)
	}
}

func AuthorizationMiddleware(allowedRole string, next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		tokenString, err := extractBearerToken(ctx)
		if err != nil {
			sendError(ctx, err.Error(), fasthttp.StatusUnauthorized)
			return
		}

		claims, err := tokenVarifications(tokenString)
		if err != nil {
			ctx.Error("Unauthorized", fasthttp.StatusUnauthorized)
			return
		}

		userRole := claims.AuthorizationTo
		if !isAllowedRole(userRole, allowedRole) {
			sendError(ctx, "Forbidden - User role not authorized", fasthttp.StatusForbidden)
			return
		}

		ctx.SetUserValue(userIDKey, claims.UserId)
		next(ctx)
	}
}

func sendError(ctx *fasthttp.RequestCtx, message string, statusCode int) {
	ctx.SetStatusCode(statusCode)
	ctx.SetContentType(contentTypeJSON)
	ctx.Error(message, statusCode)
}

func isAllowedRole(userRole []string, allowedRoles string) bool {
	for _, role := range userRole {
		if strings.TrimSpace(role) == strings.TrimSpace(allowedRoles) {
			return true
		}
	}
	return false
}
