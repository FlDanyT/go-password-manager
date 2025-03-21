package main 

import (

	"fmt"
	"strings"

	"demo/password/account" // Моя библиотеки
	"demo/password/output"

	"github.com/fatih/color" // Сторонняя библиотека для цветного текста с официального сайта: https://pkg.go.dev/

)


var menuVariants = []string{

	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по логину",
	"4. Удалить URL",
	"5. Выход",
	"Выберите вариант",

}


func main() {

	fmt.Println("__Менеджер паролей__")

	vault := account.NewVault()

Menu:		 

	for {

		variant := promptData(menuVariants...)

		switch variant {
		case "1":
			createAccount(vault)

		case "2":
			findAccountByUrl(vault)

		case "3":
			findAccountByLogin(vault)

		case "4":
			deleteAccount(vault)

		default:
			break Menu

		}

	}

}


func findAccountByUrl(vault *account.Vault) {

	url := promptData("Введите URL для поиска")

	accounts := vault.FindAccounts(url, func (acc account.Account, str string) bool { // Анонимная функция

		return strings.Contains(acc.Url, str)

	})

	outputResult(&accounts)

}


func findAccountByLogin(vault *account.Vault) {

	login := promptData("Введите Login для поиска")

	accounts := vault.FindAccounts(login, func (acc account.Account, str string) bool {

		return strings.Contains(acc.Login, str)

	})

	outputResult(&accounts)
	
}


func outputResult(accounts *[]account.Account) {

	if len(*accounts) == 0 {

		color.Red("Аккаунтов не найдено")

	}

	for _, account  := range *accounts {

		account.Output()

}

}

func deleteAccount(vault *account.Vault) {

	url := promptData("Введите URL для поиска")
	isDeleted := vault.DeleteAccountsByUrl(url)

	if isDeleted {

		color.Green("Удалено")

	} else {

		output.PrintError("Не найдено")


	}

}

func createAccount(vault *account.Vault) {

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	MyAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат URL или Login")
		return
	}

	vault.AddAccount(*MyAccount)
	
}


func promptData(prompt ...string) string{

	for i, line := range prompt {

		if i == len(prompt) - 1 {

			fmt.Printf("%v: ", line)

		} else {

			fmt.Println(line)

		}

	}

	var res string
	fmt.Scanln(&res)

	return res

}