package main

import (
	"context"
	"log"
	"os"

	"github.com/machinebox/graphql"
)

type Response struct {
    License struct {
        Name string `json:"name"`
        Description string `json:"description"`
    }`json:"license"`
}

func main()  {
    client := graphql.NewClient("https://api.github.com/graphql")

    req := graphql.NewRequest(`
        query {
            license(key: "MIT") {
                name
                description
            }
        }
    `)
    var GH_Token = os.Getenv("GITHUB_TOKEN")
    req.Header.Add("Authorization", "bearer "+GH_Token)
    ctx := context.Background()
    var response Response
    err := client.Run(ctx, req, &response)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(response.License.Description)
}
