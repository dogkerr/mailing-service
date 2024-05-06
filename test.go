package main

import (
	"github.com/dogkerr/mailing-service/m/v2/structs"
	"github.com/dogkerr/mailing-service/m/v2/utils"
)

func main() {
	var dummyData = structs.VerificationData{Name: "David", VerificationLink: "http://localhost:8080/verify"}

	utils.SendGomail(structs.Verification, dummyData, "davidlou0810@gmail.com", "Verification")
}
