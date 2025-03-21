package main
import (
	"github.com/gofiber/fiber/v2"
	

)
func handler(c * fiber.Ctx)error {
	result := c.Params("foo")
	buffer := make([]byte , len(result))
	copy(buffer,result)
	resultCopy := string(buffer)
	return c.SendString(resultCopy)
}
func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx)  error {
		return c.SendString("hello world my first app with fiber")
	}) 
	app.Listen(":3000")
		
	}
