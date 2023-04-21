package app

import (
	"fmt"
	"os"

	"gopkg.in/launchdarkly/go-sdk-common.v2/lduser"
	ld "gopkg.in/launchdarkly/go-server-sdk.v5"
)

type Calculator struct {
	ldClient *ld.LDClient
	user     lduser.User
}

func NewCalculator(ldClient *ld.LDClient, user lduser.User) *Calculator {
	return &Calculator{ldClient: ldClient, user: user}
}

func (c *Calculator) Run() {

	if c.ldClient.Initialized() {
		fmt.Println("Client initialized successfully!")
	} else {
		fmt.Println("Client failed to initialize. Exiting...")
		os.Exit(1)
	}

	boolFlag, _ := c.ldClient.BoolVariation("arun-generic-app", c.user, true)
	if boolFlag {
		fmt.Println(" Numbers greater than 10 are allowed to be added")
	} else {
		fmt.Println("Numbers less than 10 are not allowed to be added")
	}

	moreFeatures, _ := c.ldClient.BoolVariation("show-more-features", c.user, true)

	if moreFeatures {
		fmt.Println("More features are enabled")
	} else {
		fmt.Println("More features are disabled")
	}

	decimalPlaces, _ := c.ldClient.IntVariation("decimal-places", c.user, 2)

	fmt.Printf("Decimal places: %d\n", decimalPlaces)

	fmt.Println("Menu:")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	if moreFeatures {
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
	}

	var inputChoice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&inputChoice)

	var num1 float64
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	var num2 float64
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	if !boolFlag {
		if num1 < 10 || num2 < 10 {
			fmt.Println("Numbers less than 10 are not allowed to perform operations")
			return
		}
	}

	switch inputChoice {
	case 1:
		fmt.Printf("Result: %f\n", num1+num2)
	case 2:
		fmt.Printf("Result: %f\n", num1-num2)
	case 3:
		if moreFeatures {
			fmt.Printf("Result: %f\n", num1*num2)
		} else {
			fmt.Println("More features are disabled")
		}
	case 4:
		if moreFeatures {
			fmt.Printf("Result: %f\n", num1/num2)
		} else {
			fmt.Println("More features are disabled")
		}
	default:
		fmt.Println("Invalid choice")
	}
}
