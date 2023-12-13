package logging

import (
	"fmt"
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

// LoggingMiddleware logs the details of each request and any errors.
func ErrorLoggingMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		start := time.Now()

		// Recover from any panics and log the error
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[ERROR] [%s] %s %s Panic: %v", ctx.RemoteAddr(), ctx.Method(), ctx.Path(), err)
			}
		}()

		// Call the handler and check for errors
		next(ctx)

		// Calculate request duration
		duration := time.Since(start)

		// Check if the response status code indicates an error
		statusCode := ctx.Response.StatusCode()
		if statusCode >= 400 {
			// Log client or server errors with details
			log.Printf("[ERROR] [%s] %s %s %d %v", ctx.RemoteAddr(), ctx.Method(), ctx.Path(), statusCode, duration)
		} else {
			// Log normal requests
			log.Printf("[%s] %s %s %d %v", ctx.RemoteAddr(), ctx.Method(), ctx.Path(), statusCode, duration)
		}
	}
}

func LoggingMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		start := time.Now()

		// Log request details
		logRequest(ctx)

		// Call the handler
		next(ctx)

		// Log response details
		logResponse(ctx, start)
	}
}

// logRequest logs the details of the incoming request.
func logRequest(ctx *fasthttp.RequestCtx) {
	headers := ""
	ctx.Request.Header.VisitAll(func(key, value []byte) {
		headers += fmt.Sprintf("%s: %s\n", string(key), string(value))
	})

	log.Printf("Incoming request: %s %s\nHeaders:\n%s\n",
		ctx.Method(), ctx.RequestURI(), headers)
}

// logResponse logs the details of the response.
func logResponse(ctx *fasthttp.RequestCtx, startTime time.Time) {
	duration := time.Since(startTime)
	log.Printf("Response: status code: %d, duration: %v\n", ctx.Response.StatusCode(), duration)
}
