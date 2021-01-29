package main

import (
	"fmt"

	//"github.com/kimmk4369/learn-go/accounts"
	"github.com/kimmk4369/learn-go/mydict"
)

func main() {
	//account := accounts.NewAccount("mk")
	//account.Deposit(10)
	//fmt.Println(account)

	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
