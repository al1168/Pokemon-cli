package main
import(
	"fmt"
)

func MapCommandB(c *config) error{
	client := c.pokemonapiClient
	responseLoc, err := client.ListLocations(c.prevUrl)
	if err != nil{
		return fmt.Errorf("failed to get loc structd from ListLocations, error: %v", err)
	}
	
	for _, result := range responseLoc.Results{
		fmt.Printf("%v \n", result.Name)
	}
	c.nextURL = &responseLoc.Next
	c.prevUrl = &responseLoc.Previous
	return nil
}