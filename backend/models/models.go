package models

import (
	"database/sql"
	"enbektap/database"
	"fmt"

	_ "github.com/lib/pq" // Импорт драйвера для PostgreSQL
)

func Models() {
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for {
		var choice int
		fmt.Println("Что вы хотите сделать?")
		fmt.Println("1) Добавить данные")
		fmt.Println("2) Обновить данные")
		fmt.Println("3) Показать все записи")
		fmt.Println("4) Удаление")
		fmt.Println("5) Выход")

		fmt.Scan(&choice)
		switch choice {
		case 1:
			if err := createRecord(db); err != nil {
				fmt.Println("Ошибка при добавлении данных:", err)
			}
		case 2:
			if err := updateRecord(db); err != nil {
				fmt.Println("Ошибка при обновлении данных:", err)
			}
		case 3:
			if err := readRecords(db); err != nil {
				fmt.Println("Ошибка при чтении данных:", err)
			}
		case 4:
			if err := deleteRecord(db); err != nil {
				fmt.Println("Ошибка при удалении данных:", err)
			}
		case 5:
			fmt.Println("Вы вышли из Crud системы")
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}

func createRecord(db *sql.DB) error {
	var CompanyName, Applicant string
	var InputId int

	fmt.Println("Введите имя компании:")
	fmt.Scan(&CompanyName)
	fmt.Println("Введите имя Апликанта:")
	fmt.Scan(&Applicant)
	fmt.Println("Введите ID:")
	fmt.Scan(&InputId)

	query := `INSERT INTO "employee" ("company", "applicant", "id") VALUES ($1, $2, $3)`
	_, err := db.Exec(query, CompanyName, Applicant, InputId)
	if err != nil {
		return err
	}
	fmt.Println("Данные добавлены успешно!")
	return nil
}

func updateRecord(db *sql.DB) error {
	var CompanyName, Applicant string
	var InputId int

	fmt.Println("Введите имя новой компании:")
	fmt.Scan(&CompanyName)
	fmt.Println("Введите имя нового Апликанта:")
	fmt.Scan(&Applicant)
	fmt.Println("Введите ID: чье имя и компанию хотите поменять")
	fmt.Scan(&InputId)

	query := `UPDATE employee SET company = $1, applicant = $2 WHERE id = $3`
	_, err := db.Exec(query, CompanyName, Applicant, InputId)
	if err != nil {
		return err
	}
	fmt.Println("Данные таблицы под номером", InputId, "обновлены")
	return nil
}

func readRecords(db *sql.DB) error {
	rows, err := db.Query(`SELECT * FROM "employee"`)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Результаты:")
	for rows.Next() {
		var CompanyName, Applicant string
		var InputId int
		if err := rows.Scan(&CompanyName, &Applicant, &InputId); err != nil {
			return err
		}
		fmt.Println("Имя компании:", CompanyName, "Имя аппликанта:", Applicant, "ID:", InputId)
	}

	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func deleteRecord(db *sql.DB) error {
	var InputId int
	fmt.Println("Введите Id строки, которую хотите удалить")
	fmt.Scan(&InputId)

	query := `DELETE FROM employee WHERE id = $1`
	_, err := db.Exec(query, InputId)
	if err != nil {
		return err
	}
	fmt.Println("Запись удалена успешно!")
	return nil
}
