package database

import (
	"context"
	"log"
	"rahmet/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

var users = []models.User{
	models.User{
		Name:     "Beknar",
		Surname:  "Danabek",
		Password: "password",
		Email:    "beknar.danabek@bcc.kz",
	},
	models.User{
		Name:     "Nurlan",
		Surname:  "Kambar",
		Password: "password",
		Email:    "nurlan.kambar@bcc.kz",
	},
}

var events = []models.Event{
	models.Event{
		Title:       "Pohod v goory",
		Description: "Pohod v goory",
	},
	models.Event{
		Title:       "Pohod v goory",
		Description: "Pohod v goory",
	},
}

func Migrate() {
	//Instance.AutoMigrate(&models.User{})
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	Instance.Migrator().DropTable("users")
	Instance.Migrator().DropTable("events")
	//err := Instance.Migrator().DropTable(&models.Event{}, &models.User{}).Error
	//if err != nil {
	//	log.Fatalf("cannot drop table: %v", err)
	//}

	err := Instance.WithContext(ctx).Debug().AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = Instance.WithContext(ctx).Debug().AutoMigrate(&models.Event{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err := Instance.WithContext(ctx).Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		events[i].CreatorID = users[i].ID

		err = Instance.WithContext(ctx).Debug().Model(&models.Event{}).Create(&events[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}

	log.Println("Database Migration Completed...")
}
