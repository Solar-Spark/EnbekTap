// Для начала сделайте новый db и назовите first_db
// пароль свой поставьте



// Потом создайте таблицу 
// Create table employee(
// 	Company varchar(30),
// 	Applicant varchar(50),
// 	id int Primary key
// )








package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Импорт драйвера для PostgreSQL
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "aldiyar"
	dbname   = "first_db"
	//
)

func main() {
	var CompanyName string
	var InputId int
	var Applicant string
	var choice int

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckErrror(err)
	defer db.Close()

	for {
		fmt.Println("Что вы хотите сделать?")
		fmt.Println("1) Добавить данные")
		fmt.Println("2) Обновить данные")
		fmt.Println("3) Показать все записи")
		fmt.Println("4) Удаление")
		fmt.Println("5) Выход")

		fmt.Scan(&choice)
		switch choice{
		case 1:
			fmt.Println("Введите имя компании:")
			fmt.Scan(&CompanyName)
			fmt.Println("Введите имя Апликанта:")
			fmt.Scan(&Applicant)
			fmt.Println("Введите ID:")
			fmt.Scan(&InputId)

			InsertAldik := `INSERT INTO "employee" ("company", "applicant", "id") VALUES ($1, $2, $3)`
			_, err = db.Exec(InsertAldik, CompanyName, Applicant,InputId)
			CheckErrror(err)
			fmt.Println("Данные добавлены успешно!")
			break
		case 2:
			fmt.Println("Введите имя новой компании:")
			fmt.Scan(&CompanyName)
			fmt.Println("Введите имя нового Апликанта:")
			fmt.Scan(&Applicant)
			fmt.Println("Введите ID: чье имя и компанию хотите поменять")
			fmt.Scan(&InputId)


			UpdateAldik := fmt.Sprintf(`UPDATE employee SET company = '%s', applicant = '%s' WHERE id = %d;`, CompanyName, Applicant, InputId)
			_, err := db.Exec(UpdateAldik)
			CheckErrror(err)
			fmt.Println("Данные таблицы под номером", InputId, "обновлины")
			break
		case 3:
			selectAll(db)
			break
		case 4:
			DeleteRow(db)
			break
		case 5:
			fmt.Println("Вы вышли из Crud системы")
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}

func CheckErrror(err error) {
	if err != nil {
		panic(err)
	}
}

func selectAll(db *sql.DB) {
	rows, err := db.Query(`SELECT * FROM "employee"`)
	CheckErrror(err)
	defer rows.Close()

	fmt.Println("Результаты:")
	for rows.Next() {
		var CompanyName string
		var Applicant string
		var InputId int
		if err := rows.Scan(&CompanyName, &Applicant, &InputId); err != nil {
			CheckErrror(err)
		}
		fmt.Println("Имя компании :", CompanyName, "Имя аппликанта:", Applicant,  "ID:", InputId)
	}

	if err := rows.Err(); err != nil {
		CheckErrror(err)
	}
}
func DeleteRow(db *sql.DB){
	var InputId int
	fmt.Println("Введите Id строи которую хотите удалить")
	fmt.Scan(&InputId)
	command := fmt.Sprintf(`Delete from employee where id = '%d'`, InputId)
	_, err := db.Exec(command)
	CheckErrror(err)
}



