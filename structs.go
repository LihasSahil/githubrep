package main

import (
	"fmt"
)

var IE string
var wmt int = 0
var dmt int = 0

type AccDetail struct {
	Name   string
	Accno  int
	Funds  FDT
	AcType string
}

type FDT struct {
	Balance  int
	Withdraw int
	Deposit  int
}

type Bank interface {
	AddFunds()
	GetAccountInfo()
	Withdrawfunds()
}

func (a *AccDetail) GetAccountInfo() {
	fmt.Println("name:", a.Name, "Type: ", a.AcType, "Balance:", a.Funds, "A/c no:", a.Accno)
}
func (f *FDT) WithdrawFunds() {
	fmt.Println(" withdraw amount: ")
	fmt.Scanln(&wmt)
	fmt.Println("your remaining balance", f.Balance-wmt)
}
func (f *FDT) AddFunds() {
	fmt.Println(" deposit amount: ")
	fmt.Scanln(&dmt)
	fmt.Println("your remaining balance", f.Balance+dmt)
}
func structs() {

	var John = AccDetail{Name: "John Smith", Accno: 20202020, Funds: FDT{Balance: 50000, Withdraw: 0, Deposit: 0}, AcType: "savings"}
	fmt.Printf("%+v \n", John)
	fmt.Printf("Enter Interface (WithdrawFunds,AddFunds,GetAccountInfo): ")
	fmt.Scanln(&IE)
	switch IE {
	case "GetAccountInfo":
		John.GetAccountInfo()
	case "WithdrawFunds":
		John.Funds.WithdrawFunds()
	case "AddFunds":
		John.Funds.AddFunds()
	default:
		fmt.Println("invalid input")
	}
}
