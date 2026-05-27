package models

import (
	"fmt"

	"example.com/dad_midia/db"
)

type Midia struct {
	Matricula                   int64
	Nome                        string
	StreamingFavorito           string
	FrequenciaUsoRedesSociais   float64
	MeioPrincipalNoticias       string
	ComunicacaoDigitalPrincipal string
}

func (m *Midia) Save() error {

	query := `INSERT INTO midia(matricula, nome, streaming_favorito, freq_uso_redes_sociais, meio_principal_noticias, comunicacao_digital_principal) 
	VALUES (?, ?, ?, ?, ?, ?);`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(
		m.Matricula,
		m.Nome,
		m.StreamingFavorito,
		m.FrequenciaUsoRedesSociais,
		m.MeioPrincipalNoticias,
		m.ComunicacaoDigitalPrincipal,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	m.Matricula = id

	return err
}

func GetAllMidias(matricula int64) ([]Midia, error) {
	query := `SELECT * FROM midia WHERE matricula = ?`

	rows, err := db.DB.Query(query, matricula)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var dbEvents []Midia

	for rows.Next() {
		var midia Midia
		err := rows.Scan(
			&midia.Matricula,
			&midia.Nome,
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
	query := `SELECT * FROM midia WHERE matricula = ?`

	row := db.DB.QueryRow(query, id)

	var midia Midia
	err := row.Scan(&midia.Matricula, &midia.Nome, &midia.StreamingFavorito, &midia.FrequenciaUsoRedesSociais, &midia.MeioPrincipalNoticias, &midia.ComunicacaoDigitalPrincipal)
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
		WHERE matricula = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(m.Nome, m.StreamingFavorito, m.FrequenciaUsoRedesSociais, m.MeioPrincipalNoticias, m.FrequenciaUsoRedesSociais, m.ComunicacaoDigitalPrincipal)
	return err
}

func (m Midia) Delete() error {
	query := `
		DELETE FROM midia 
		WHERE matricula = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(m.Matricula)
	return err
}
