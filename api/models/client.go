package models

import (
	"fmt"
	"database/sql"
)


type Client struct {
	ID   int64    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Phone int    `json:"phone"`
}

func GetClients(db *sql.DB) []Client {
	query := "SELECT id, name, age, phone FROM clients"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	var clients []Client
	for rows.Next() {
		var client Client
		if err := rows.Scan(&client.ID, &client.Name, &client.Age, &client.Phone); err != nil {
			fmt.Println(err)
			return nil
		}
		clients = append(clients, client)
	}
	return clients
}

func GetClientByID(db *sql.DB, id int64) (Client, error) {
	query := "SELECT id, name, age, phone FROM clients WHERE id = ?"
	row := db.QueryRow(query, id)
	var client Client
	if err := row.Scan(&client.ID, &client.Name, &client.Age, &client.Phone); err != nil {
		return Client{}, err
	}

	return client, nil
}

func CreateClient(db *sql.DB, client Client) (int64, error) {
	query := "INSERT INTO clients (name, age, phone) VALUES (?, ?, ?) RETURNING id"

	res, err := db.Exec(query, client.Name, client.Age, client.Phone)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func UpdateClient(db *sql.DB, client Client) error {
	query := "UPDATE clients SET name = ?, age = ?, phone = ? WHERE id = ?"
	_, err := db.Exec(query, client.Name, client.Age, client.Phone, client.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteClient(db *sql.DB, id int64) error {
	query := "DELETE FROM clients WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}