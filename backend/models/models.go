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

	fmt.Print("Введите имя компании: ")
	fmt.Scan(&CompanyName)
	fmt.Print("Введите имя Апликанта: ")
	fmt.Scan(&Applicant)
	fmt.Print("Введите ID: ")
	fmt.Scan(&InputId)

	query := `INSERT INTO "employee" ("company", "applicant", "id") VALUES ($1, $2, $3)`
	_, err := db.Exec(query, CompanyName, Applicant, InputId)
	if err != nil {
		return err
	}
	fmt.Println("Данные добавлены успешно!")
	fmt.Println("-------------------------")
	return nil
}

func updateRecord(db *sql.DB) error {
	var InputId int
	fmt.Println("Введите ID записи, которую хотите обновить:")
	fmt.Scan(&InputId)

	var currentCompanyName, currentApplicant string
	query := `SELECT company, applicant FROM employee WHERE id = $1`
	err := db.QueryRow(query, InputId).Scan(&currentCompanyName, &currentApplicant)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Запись с таким ID не найдена.")
			return nil
		}
		return err
	}
	fmt.Println("-------------------------")
	fmt.Println("Текущие данные:")
	fmt.Println("Имя компании:", currentCompanyName)
	fmt.Println("Имя аппликанта:", currentApplicant)
	fmt.Println("-------------------------")
	// Prompt for new data
	var newCompanyName, newApplicant string
	fmt.Println("Введите новое имя компании (оставьте пустым, чтобы не изменять):")
	fmt.Scan(&newCompanyName)
	fmt.Println("Введите новое имя аппликанта (оставьте пустым, чтобы не изменять):")
	fmt.Scan(&newApplicant)

	if newCompanyName == "" {
		newCompanyName = currentCompanyName
	}
	if newApplicant == "" {
		newApplicant = currentApplicant
	}

	updateQuery := `UPDATE employee SET company = $1, applicant = $2 WHERE id = $3`
	_, err = db.Exec(updateQuery, newCompanyName, newApplicant, InputId)
	if err != nil {
		return err
	}

	fmt.Println("Данные таблицы под номером", InputId, "обновлены")
	fmt.Println("-------------------------")
	return nil
}

func readRecords(db *sql.DB) error {
	rows, err := db.Query(`SELECT * FROM "employee"`)
	if err != nil {
		return err
	}
	defer rows.Close()
	fmt.Println("-------------------------")
	fmt.Println("Результаты:")
	for rows.Next() {
		var CompanyName, Applicant string
		var InputId int
		if err := rows.Scan(&CompanyName, &Applicant, &InputId); err != nil {
			return err
		}
		fmt.Println("Имя компании:", CompanyName, "\nИмя аппликанта:", Applicant, "\nID:", InputId)
	}

	if err := rows.Err(); err != nil {
		return err
	}
	fmt.Println("-------------------------")
	return nil
}

func deleteRecord(db *sql.DB) error {
	var InputId int
	fmt.Println("Введите Id строки, которую хотите удалить:")
	fmt.Scan(&InputId)

	var currentCompanyName, currentApplicant string
	query := `SELECT company, applicant FROM employee WHERE id = $1`
	err := db.QueryRow(query, InputId).Scan(&currentCompanyName, &currentApplicant)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Запись с таким ID не найдена.")
			fmt.Println("-------------------------")
			return nil
		}
		return err
	}
	fmt.Println("-------------------------")
	fmt.Println("Вы собираетесь удалить следующую запись:")
	fmt.Println("-------------------------")
	fmt.Println("Имя компании:", currentCompanyName)
	fmt.Println("Имя аппликанта:", currentApplicant)
	fmt.Println("-------------------------")

	var confirmation string
	fmt.Println("Вы уверены, что хотите удалить эту запись? (да/нет):")
	fmt.Scan(&confirmation)

	if confirmation != "да" {
		fmt.Println("Удаление отменено.")
		fmt.Println("-------------------------")
		return nil
	}

	deleteQuery := `DELETE FROM employee WHERE id = $1`
	_, err = db.Exec(deleteQuery, InputId)
	if err != nil {
		return err
	}

	fmt.Println("Запись удалена успешно!")
	fmt.Println("-------------------------")
	return nil
}
