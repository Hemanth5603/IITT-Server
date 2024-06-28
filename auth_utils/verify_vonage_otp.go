package auth_utils

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func VerifyVonageOTP(requestID string, otp string) (string, string, error) {
	auth := vonage.CreateAuthFromKeySecret(envVAPIKEY(), envVSECRETKEY())

	verifyClient := vonage.NewVerifyClient(auth)

	response, errResp, err := verifyClient.Check(requestID, otp)

	if err != nil {
		fmt.Printf("%#v\n", err)
	} else if response.Status != "0" {
		fmt.Println("Error status " + errResp.Status + ": " + errResp.ErrorText)
	} else {
		// all good
		fmt.Println("Request complete: " + response.RequestId)
	}
	return response.RequestId, response.Status, err
}
