package auth_utils

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func SendVonageOTP(phoneNumber string) (string, string, error) {
	auth := vonage.CreateAuthFromKeySecret(envVAPIKEY(), envVSECRETKEY())

	verifyClient := vonage.NewVerifyClient(auth)

	response, errResp, err := verifyClient.Request(phoneNumber, "IITTNiF", vonage.VerifyOpts{CodeLength: 6, WorkflowID: 4})

	if err != nil {
		fmt.Printf("%#v\n", err)
	} else if response.Status != "0" {
		fmt.Println("Error status " + errResp.Status + ": " + errResp.ErrorText)
	} else {
		fmt.Println("Request started: " + response.RequestId)
	}
	return response.RequestId, response.Status, err

}
