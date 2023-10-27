package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

func main() {
	//var modelName string
	var migrationName string
	var outputDir string

	//flag.StringVar(&modelName, "model", "", "GORM model name")
	flag.StringVar(&migrationName, "name", "", "Migration name")
	flag.StringVar(&outputDir, "output", "./database/migration", "Output directory for migrations")
	flag.Parse()

	if migrationName == "" {
		fmt.Println("Usage: make migrate -name <migration_name>")
		os.Exit(1)
	}

	var methodName string
	if strings.Contains(migrationName, "_") {
		str := strings.ReplaceAll(migrationName, "_", " ")
		methodName = strings.ReplaceAll(toCapitalize(str), " ", "")
	} else {
		methodName = strings.ReplaceAll(toCapitalize(migrationName), " ", "")
	}

	// Create the migration file content
	timestamp := time.Now().Format("20060102150405")
	migrationContent := fmt.Sprintf(`package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)
	
func %sTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "%s_%s",
		Migrate: func(tx *gorm.DB) error {
			// Add migration logic here

			// if err := tx.Table("users").AutoMigrate(&Users{}); err != nil {
			//	return err
			// }
			// if err := tx.Exec(CREATE UNIQUE INDEX users_email_uidx ON users (email) WHERE deleted_at IS NULL;).Error; err != nil {
			//	return err
			// }

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			// Add rollback logic here
			return tx.Migrator().DropTable("tableName")
		},
	}
}
`, methodName, timestamp, migrationName)

	// Create the migration file
	migrationFileName := fmt.Sprintf("%s/%s_%s.go", outputDir, timestamp, migrationName)
	err := os.WriteFile(migrationFileName, []byte(migrationContent), 0644)
	if err != nil {
		fmt.Printf("Error creating migration file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Migration file created: %s\n", migrationFileName)
}

func toCapitalize(str string) string {
	var output []rune
	isWord := true
	for _, val := range str {
		if isWord && unicode.IsLetter(val) {
			output = append(output, unicode.ToUpper(val))
			isWord = false
		} else if !unicode.IsLetter(val) {
			isWord = true
			output = append(output, val)
		} else {
			output = append(output, val)
		}
	}
	return string(output)
}
