package main

import (
	"fmt"
	"os"
)

type Overview struct {
	Name   string
	UserId int
	Funds  bal
	AcType string
}

func (o *Overview) ChangeAcType(newAcType string) {
	o.AcType = newAcType
}

type bal struct {
	Balance    int
	Withdrawal int
	Deposit    int
}

func structs() {
	var UserNo int
	var TFunc string
	var Wamount int = 0
	var Damount int = 0

	fmt.Println("enter UserNo:")
	fmt.Scanln(&UserNo)

	if UserNo == 20202020 {
		var John = Overview{Name: "John Smith", UserId: 20202020, Funds: bal{Balance: 50000, Withdrawal: 00, Deposit: 0}, AcType: "savings"}
		fmt.Printf("%+v \n", John)

		Newbal := bal{
			Balance:    50000,
			Withdrawal: 0,
			Deposit:    0,
		}
		fmt.Println("balance is :", Newbal.Balance-Newbal.Withdrawal+Newbal.Deposit)
		John.ChangeAcType("checking")
		fmt.Printf("%+v \n", John)
		fmt.Println("would you like to perform a transaction : (W)Withdrawal, Deposit(D),(B) checkBalance:")
		fmt.Scanln(&TFunc)
		switch TFunc {
		case "W":
			fmt.Println("amount:")
			fmt.Scanln(&Wamount)
			fmt.Println("balance is :", Newbal.Balance-Wamount+Newbal.Deposit)
		case "D":
			fmt.Println("amount:")
			fmt.Scanln(&Damount)
			fmt.Println("balance is :", Newbal.Balance-Newbal.Withdrawal+Damount)

		case "B":
			fmt.Println("balance is :", Newbal.Balance-Wamount+Newbal.Deposit)

		default:
			fmt.Println("invalid option")
		}
	} else {
		fmt.Println("Wrong UserID")
		os.Exit(3)
	}
}
