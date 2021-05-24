package modal

import (
	"database/sql"
)

type PhoneNumber struct {
	ID      int    `json:"id"`
	Country string `json:"country"`
	Number  string `json:"number"`
	Code    string `json:"code"`
	State   string `json:"state"`
}

func (p *PhoneNumber) GetPhoneNumber(db *sql.DB) error {
	return db.QueryRow("SELECT id, country, number, code, state FROM phone_numbers WHERE id=$1",
		p.ID).Scan(&p.ID, &p.Country, &p.Number, &p.Code, &p.Code)
}

func (p *PhoneNumber) UpdatePhoneNumber(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE phone_numbers SET country=$1, number=$2, code=$3, state=$4 WHERE id=$5",
			p.Country, p.Number, p.Code, p.State)

	return err
}

func (p *PhoneNumber) DeletePhoneNumber(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM phone_numbers WHERE id=$1", p.ID)

	return err
}

func (p *PhoneNumber) CreatePhoneNumber(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO phone_numbers(country, number, code, state) VALUES($1, $2, $3, $4) RETURNING id",
		p.Country, p.Number, p.Code, p.State).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetPhoneNumbers(db *sql.DB, start, count int) (numbers []PhoneNumber, err error) {
	rows, err := db.Query(
		"SELECT id, country, number, code, state FROM phone_numbers LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var p PhoneNumber
		if err := rows.Scan(&p.ID, &p.Country, &p.Number, &p.Code, &p.State); err != nil {
			return nil, err
		}
		numbers = append(numbers, p)
	}
	return
}

func GetPhoneNumbersCounts(db *sql.DB) (int, error) {
	rows, err := db.Query(
		"SELECT count(*) as count_Phone FROM phone_numbers ")

	if err != nil {
		return 0, err
	} else {

		var count_Phone int
		for rows.Next() {
			rows.Scan(&count_Phone)
		}
		return count_Phone, nil
	}
}
