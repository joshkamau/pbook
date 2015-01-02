package dao

import (
	_ "github.com/lib/pq"
	"database/sql"
	"bitbucket.org/joshnet/pbook/model"
)

const (
	DB_DRIVER string = "postgres"
	DB_CONN string = "dbname=phonebook user=postgres password=pass sslmode=disable"
)

func getDB()(*sql.DB, error){
	db, err := sql.Open(DB_DRIVER, DB_CONN)
	return db, err
}

func SaveContact(c *model.Contact)(int, error){
	db, err := getDB()
	defer db.Close()

	if err != nil {
		return 0, err
	}

	var id int
	queryErr := db.QueryRow("INSERT INTO contact (name, phone_number) VALUES($1, $2) RETURNING id",
		c.Name, c.PhoneNumber).Scan(&id)
	if queryErr != nil {
		return 0, queryErr
	}

	return id, nil
}

func DeleteContact(id int) error {
	db, err := getDB()
	defer db.Close()
	if err != nil {
		return err
	}
	_, queryErr := db.Exec("DELETE FROM contact WHERE id = $1", id)

	if queryErr != nil {
		return queryErr
	}
	return nil
}

func GetByName(name string)([]*model.Contact, error){
	db, err := getDB()
	defer db.Close()

	if err != nil {
		return nil, err
	}

	rows, queryErr := db.Query("SELECT id, name, phone_number FROM contact WHERE LOWER(name) LIKE LOWER($1)", "%"+name +"%")

	if queryErr != nil {
		return nil, queryErr
	}
	var contacts []*model.Contact
	for rows.Next() {
		var id int
		var name string
		var phoneNumber string
		rows.Scan(&id, &name, &phoneNumber)
		contacts = append(contacts, &model.Contact{Id:id, Name:name, PhoneNumber:phoneNumber})
	}
	return contacts, nil
}

func GetAll()([]*model.Contact, error){
	return GetByName("%")
}
