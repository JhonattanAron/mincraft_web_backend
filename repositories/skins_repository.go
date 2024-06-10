package repositories

import (
	"fmt"

	"github.com/minecraft_bakend_web/models"
)

func ListSkinHandler() []models.Skin_Data {
	var uuid string
	var skin_identifier string
	db, err := GetConnectionDb()
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	defer db.Close()

	result, err := db.Query("SELECT uuid, skin_identifier FROM sr_players")
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	skins := []models.Skin_Data{}
	for result.Next() {
		var skin models.Skin_Data
		err = result.Scan(&uuid, &skin_identifier)
		if err != nil {
			fmt.Println("Error al escanear:", err.Error())
			continue
		}
		skin.Uuid = uuid
		skin.Skin_identifier = skin_identifier
		skins = append(skins, skin)
	}

	return skins
}

func GetSkinByUuid(uuid string) (skin string) {
	var skin_identifier string
	db, err := GetConnectionDb()
	if err != nil {
		fmt.Println("Err", err.Error())
		return "eror al Conectarse con la Db" + err.Error()
	}
	defer db.Close()
	result, err := db.Query("SELECT skin_identifier FROM sr_players WHERE uuid =?", uuid)
	if err != nil {
		fmt.Println("Err", err.Error())
		return "Eror al Ejcutar la consulta " + err.Error()
	}
	for result.Next() {
		err = result.Scan(&skin_identifier)
		if err != nil {
			fmt.Println("Error al escanear:", err.Error())
			continue
		}
	}
	return skin_identifier
}
