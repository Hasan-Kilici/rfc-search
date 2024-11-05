package middlewares

import "github.com/gofiber/fiber/v2/middleware/cors"

var Cors = cors.New(cors.Config{
        AllowOrigins		:     	"http://localhost:3000",
        AllowHeaders		:     	"Origin, Content-Type, Accept",
        AllowCredentials	: 	true,
})