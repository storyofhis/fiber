# fiber documentation

## this is open for public 🤙 🚀

installation: 
```shellsession
go get github.com/gofiber/fiber/v2
```
before you want to running. you should find into your main file. 
running file: 
```shellsession
go run main.go
```

simple routing on port :8000
```go
package main
func main() {
    app := fiber.New()
    
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })
    app.Listen(":8080")
}
```

routing with your own params
```go
// GET http://localhost:8080/hello%20world

app.Get("/:value", func(c *fiber.Ctx) error {
  return c.SendString("value: " + c.Params("value"))
  // => Get request with value: hello world
})
```

optional params : 
```go
// GET http://localhost:3000/john

app.Get("/:name?", func(c *fiber.Ctx) error {
  if c.Params("name") != "" {
    return c.SendString("Hello " + c.Params("name"))
    // => Hello john
  }
  return c.SendString("Where is john?")
})
```

wildscard: 
```go
// GET http://localhost:3000/api/user/john

app.Get("/api/*", func(c *fiber.Ctx) error {
  return c.SendString("API path: " + c.Params("*"))
  // => API path: user/john
})
```
