package config

import (
	"log"

	"github.com/torrez/pkg"
	"github.com/torrez/src/models"
	"github.com/torrez/src/repository"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	createAdminUser(db)
}

func createAdminUser(db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)

	existingAdmin, err := userRepo.FindByEmail("admin@gmail.com")
	if err == nil && existingAdmin != nil {
		log.Println("El usuario administrador ya existe. No se creará nuevamente.")
		return
	}

	hashedPassword, err := pkg.HashPassword("admin")
	if err != nil {
		log.Fatal("Error al hashear la contraseña", err)
	}

	admin := &models.User{
		Name:     "admin",
		Email:    "admin@gmail.com",
		Phone:    "12345",
		Rol:      models.Administrador,
		Password: hashedPassword,
	}

	if err := userRepo.Create(admin); err != nil {
		log.Fatal("Error al crear el usuario administrador", err)
	} else {
		log.Println("Usuario administrador creado exitosamente.")
	}
}
