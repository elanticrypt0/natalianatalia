package webcore_features

import (
	"log"

	"github.com/k23dev/go4it"
	"github.com/k23dev/natalianatalia/api/models"
	"github.com/k23dev/natalianatalia/pkg/webcore"
	"gorm.io/gorm"
)

func SeedCMD(gas *webcore.GasonlineApp) {
	db := gas.App.DB.Primary
	seedCategories(db)
	log.Println("###")
}

func SeedTable(gas *webcore.GasonlineApp, table string) {
	// todo
	// data := []models.Zone{}
	// file := fmt.Sprintf("./seeds/" + table + ".json")
	// go4it.ReadOrParseJson(file, &zone_list)
	// gas.App.DB.Primary.Save(&zone_list)
	// log.Println("Zones seeded")
}

func seedCategories(db *gorm.DB) {
	cat_list := []models.Category{}
	go4it.ReadAndParseJson("./seeds/categories.json", &cat_list)
	db.Save(&cat_list)
	log.Println("Categories seeded")
}
