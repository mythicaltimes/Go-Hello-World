package main

import (
	"fmt"

	"github.com/jegfish/goroyale"
)

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjI4OSwiaWRlbiI6IjI0NjY2OTgwNjg5ODUxMTg3MiIsIm1kIjp7fSwidHMiOjE1NDgyNjM2MzA2MzB9.AsefPX4hNnnQpCkG6w3S47PwgsXXOU6y8zt_Ngp7Fr4"

func main() {
	c, err := goroyale.New(token, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	ver, err := c.APIVersion()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("API Version:", ver)
	}

	params := map[string][]string{
		"exclude": {"name"},
	}
	
	p, err := c.Clan("9R2QG2J2", params)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Name:", p.Name)
		fmt.Println("Tag", p.Tag)
	}
}
