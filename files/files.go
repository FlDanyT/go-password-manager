package files

import (

	"os"
	
	"demo/password/output"
	"github.com/fatih/color"

)


type JsonDb struct {

	filename string

}


func NewJsonDb(name string) *JsonDb {

	return &JsonDb {

		filename: name,

	}

}


func (db *JsonDb) Read() ([]byte, error){

	data, err := os.ReadFile(db.filename) // Читаем файл
	if err != nil {

		return nil, err

	}

	return data, nil

}


func (db *JsonDb) Write(content []byte) {

	file, err := os.Create(db.filename) // Создаем файл
	if err != nil {
		
		output.PrintError(err)

	}

	_, err  = file.Write(content) // Делаем запись
	defer file.Close() // defer позволяет закрыть файл после завершения функции

	if err != nil {
		
		output.PrintError(err)
		return

	}

	color.Green("Запись успешна!")

}