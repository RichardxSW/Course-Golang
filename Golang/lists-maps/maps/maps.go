package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "www.google.com",
		"Amazon Web Services": "aws.amazon.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])

	websites["Facebook"] = "www.facebook.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
