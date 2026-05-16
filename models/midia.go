package models

import (
	"fmt"

	"example.com/dad_midia/db"
)

type Midia struct {
	ID                          int64
	Estudante                   int64
	StreamingFavorito           string
	FrequenciaUsoRedesSociais   float64
	MeioPrincipalNoticias       string
	ComunicacaoDigitalPrincipal string
}

func (m *Midia) Save() error {

	query := `INSERT INTO midia(estudante, streaming_favorito, freq_uso_redes_sociais, meio_principal_noticias, comunicacao_digital_principal) 
	VALUES (?, ?, ?, ?, ?);`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(
		m.Estudante,
		m.StreamingFavorito,
		m.FrequenciaUsoRedesSociais,
		m.MeioPrincipalNoticias,
		m.ComunicacaoDigitalPrincipal,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	m.ID = id

	return err
}

func GetAllMidias(estudante int64) ([]Midia, error) {
	query := `SELECT * FROM Midia WHERE estudante = ?`

	rows, err := db.DB.Query(query, estudante)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var dbEvents []Midia

	for rows.Next() {
		var midia Midia
		err := rows.Scan(
			&midia.ID,
			&midia.Estudante,
			&midia.StreamingFavorito,
			&midia.FrequenciaUsoRedesSociais,
			&midia.MeioPrincipalNoticias,
			&midia.ComunicacaoDigitalPrincipal,
		)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		dbEvents = append(dbEvents, midia)
	}

	return dbEvents, nil
}

func GetMidiaByID(id int64) (*Midia, error) {
	query := `SELECT * FROM midia WHERE id_midia = ?`

	row := db.DB.QueryRow(query, id)

	var midia Midia
	err := row.Scan(&midia.ID, &midia.Estudante, &midia.StreamingFavorito, &midia.FrequenciaUsoRedesSociais, &midia.MeioPrincipalNoticias, &midia.ComunicacaoDigitalPrincipal)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &midia, nil
}

func (m Midia) Update() error {
	query := `
		UPDATE midia 
		SET streaming_favorito = ?, freq_uso_redes_sociais = ?, meio_principal_noticias = ?, comunicacao_digital_principal = ? 
		WHERE id_midia = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(m.StreamingFavorito, m.FrequenciaUsoRedesSociais, m.MeioPrincipalNoticias, m.FrequenciaUsoRedesSociais, m.ComunicacaoDigitalPrincipal)
	return err
}

func (m Midia) Delete() error {
	query := `
		DELETE FROM midia 
		WHERE id_midia = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(m.ID)
	return err
}
