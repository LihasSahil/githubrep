package main

import "fmt"

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
	var John = Overview{Name: "John Smith", UserId: 20202020, Funds: bal{Balance: 50000, Withdrawal: 2500, Deposit: 300}, AcType: "savings"}
	fmt.Printf("%+v \n", John)

	Newbal := bal{
		Balance:    50000,
		Withdrawal: 2500,
		Deposit:    300,
	}
	fmt.Println("balance is :", Newbal.Balance-Newbal.Withdrawal+Newbal.Deposit)
	John.ChangeAcType("checking")
	fmt.Printf("%+v \n", John)
}
