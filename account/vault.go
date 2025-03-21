package account

import (

	"encoding/json"
	"time"
	"strings" // Для поиска url из json

	"demo/password/files"
	"demo/password/output"

)

type Vault struct { // Делаем новую структуру где Account в массиве

	Accounts []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`

}


func NewVault() *Vault {

	db := files.NewJsonDb("data.json")

	file, err := db.Read()
	if err != nil{

		return &Vault {

			Accounts: []Account{},
			UpdatedAt: time.Now(),
	
		}

	}

	var vault Vault

	err = json.Unmarshal(file, &vault) // Преобразования JSON-формата в объект или структуру данных внутри программы.

	if err != nil {

		output.PrintError("Не удалось разобрать файл data.json")

		return &Vault {

			Accounts: []Account{},
			UpdatedAt: time.Now(),
	
		}

	}

	return &vault // Возвращаем файл JSON в нужном формате

}


func (vault *Vault) DeleteAccountsByUrl( url string ) bool {

	var accounts [] Account // Список найденных аккаунтов
	isDeleted := false

	for _, account := range vault.Accounts {

		isMatched := strings.Contains(account.Url, url) // Поиск аккаунта

		if !isMatched { // Ищем url которые не входят в список для удаления

			accounts = append(accounts, account) // Если нашли добавляем в список
			continue

		}

		isDeleted = true

	}

	vault.Accounts = accounts
	vault.save()

	return isDeleted

}


func (vault *Vault) FindAccounts( str string, checker func(Account, string) bool) [] Account {
 
	var accounts []Account // Список найденных аккаунтов

	for _, account := range vault.Accounts {

		isMatched := checker(account, str) // Поиск аккаунта

		if isMatched {

			accounts = append(accounts, account) // Если нашли добавляем в список

		}

	}

	return accounts

}


func (vault *Vault) AddAccount(acc Account) { // Добавляем значения в объект

	vault.Accounts = append(vault.Accounts, acc) // Добавляем данные
	vault.save()
}


func (vault *Vault) ToBytes() ([]byte, error) {

	file, err :=  json.Marshal(vault) // Преобразование в байты для записи в json
	if err != nil {
			return nil, err
	}
	return file, nil

}

func (vault *Vault) save() { // Метод записи данных

	vault.UpdatedAt = time.Now()

	data, err := vault.ToBytes()
	if err != nil {

		output.PrintError("Не удалось преобразовать")

	}

	db := files.NewJsonDb("data.json")
	db.Write(data) // Записываем данные

}