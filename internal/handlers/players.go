package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"sportteam/internal/db"
	"sportteam/internal/models"
)

// HandlePlayers — обработчик /players
func HandlePlayers(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")

	var rows *sql.Rows
	var err error

	if q == "" {
		rows, err = db.DB.Query(`
			SELECT
				p.id,
				pn.last_name || ' ' || pn.first_name ||
				COALESCE(' ' || pn.middle_name, '') AS full_name,
				c.city_name,
				p.height_cm,
				p.weight_kg
			FROM player p
			JOIN player_name pn ON p.player_name_id = pn.id
			JOIN city c        ON p.city_id = c.id
			ORDER BY p.id;
		`)
	} else {
		pattern := "%" + q + "%"
		rows, err = db.DB.Query(`
			SELECT
				p.id,
				pn.last_name || ' ' || pn.first_name ||
				COALESCE(' ' || pn.middle_name, '') AS full_name,
				c.city_name,
				p.height_cm,
				p.weight_kg
			FROM player p
			JOIN player_name pn ON p.player_name_id = pn.id
			JOIN city c        ON p.city_id = c.id
			WHERE (pn.last_name || ' ' || pn.first_name || ' ' || COALESCE(pn.middle_name, '')) ILIKE $1
			   OR c.city_name ILIKE $1
			ORDER BY p.id;
		`, pattern)
	}

	if err != nil {
		http.Error(w, "Ошибка запроса к базе данных", http.StatusInternalServerError)
		log.Printf("Ошибка db.Query: %v\n", err)
		return
	}
	defer rows.Close()

	var players []models.Player

	for rows.Next() {
		var pl models.Player
		if err := rows.Scan(&pl.ID, &pl.FullName, &pl.City, &pl.Height, &pl.Weight); err != nil {
			http.Error(w, "Ошибка чтения данных из БД", http.StatusInternalServerError)
			log.Printf("Ошибка rows.Scan: %v\n", err)
			return
		}
		players = append(players, pl)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Ошибка при работе с данными БД", http.StatusInternalServerError)
		log.Printf("rows.Err: %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(players); err != nil {
		http.Error(w, "Ошибка кодирования JSON", http.StatusInternalServerError)
		log.Printf("Ошибка json.Encode: %v\n", err)
		return
	}
}
