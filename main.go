package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/launchdarkly/go-sdk-common.v2/lduser"
	ld "gopkg.in/launchdarkly/go-server-sdk.v5"

	"LaunchDarkly/app"
)

const (
	sdkKey         = "sdk-4bb7b64a-906b-4307-8d24-5381bf8124fd"
	featureFlagKey = "arun-generic-app"
)

func main() {
	ldClient, _ := ld.MakeClient(sdkKey, 5*time.Second)
	if ldClient.Initialized() {
		fmt.Println("Client initialized successfully!")
	} else {
		fmt.Println("Client failed to initialize. Exiting...")
		os.Exit(1)
	}

	user := lduser.NewUser("arunprasath42")

	calc := app.NewCalculator(ldClient, user)
	calc.Run()
}
