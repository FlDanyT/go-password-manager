package account

import (

	"fmt"
	"errors"
	"time"

	"math/rand/v2" // Генирация рандомных значений
	"net/url" // Проверка на url

)


var letterRunes = []rune("1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz") // Символы для генирации пороля


type Account struct { //  Указываем тип данных для полей

	Login string `json:"login"` // Meta информация в `` описание тега
	Password string `json:"password"`
	Url string `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

}


func (acc *Account) Output() { // Метод, для вывода значений

	fmt.Println(acc.Login, acc.Password, acc.Url)

}


func (acc *Account) generatePassword(n int){

	res := make([]rune, n) // Итоговый массив длинны n

	for i := range res {
		
		res[i] = letterRunes[rand.IntN(len(letterRunes))] // Берем рандомные значение и подставляем для res[i]

	}

	acc.Password = string(res) // Сохраняем в аккаунт

}


func NewAccount(login, password, urlString string) (*Account, error) {

	if login == "" {
		return nil,errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString) // Валидация на url
	if err != nil {
		return nil, errors.New("INVALID_URL") // Возвращаем ошибку
	}

	newAcc := &Account{ // Добавляем получаемые данные от пользователя

		Url: urlString,
		Login: login,
		Password: password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

	}

	if password == "" {
		
		newAcc.generatePassword(12)

	}

	return newAcc, nil

}