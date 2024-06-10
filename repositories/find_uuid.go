package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/minecraft_bakend_web/config"
)

func FindUUIDByUsername(username string) (string, error) {
	configUser := config.LoadConfig()
	db, err := sql.Open("sqlite3", configUser.Ruta_Luck_Perms)
	if err != nil {
		return "", fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	var uuid string
	query := "SELECT uuid FROM luckperms_players WHERE username = ?"
	err = db.QueryRow(query, strings.ToLower(username)).Scan(&uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no se encontró ningún UUID para el nombre '%s'", username)
		}
		return "", fmt.Errorf("failed to query database: %w", err)
	}

	return uuid, nil
}
