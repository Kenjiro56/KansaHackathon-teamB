package db

import (
	"KansaiHack-Friday/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DBConnect Successful",
	})
}

var DB *gorm.DB

func Connect() error {
	dsn := os.Getenv("DATABASE_URL") //←デプロイする時はDB_URLをここに入れる
	if dsn == "" {
		// 環境変数が設定されていない場合、デフォルトのDSNを使用
		dsn = "host=db user=user password=password dbname=mydatabase port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	fmt.Println("Database connection established")
	return nil
}

func Migrate() error {
	models := []interface{}{
		&models.User{},
		&models.Obj{},
		&models.Todo{},
		// migrationしたいModelを記述
	}

	err := DB.AutoMigrate(models...)
	if err != nil {
		return err
	}
	fmt.Println("Database migrated")
	return nil
}

func DropAllTables(db *gorm.DB) error {
	// PostgreSQLで全テーブルを削除
	err := db.Exec(`
			DO $$ DECLARE
					r RECORD;
			BEGIN
					-- publicスキーマ内のすべてのテーブルを削除
					FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
							EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
					END LOOP;
			END $$;
	`).Error

	if err != nil {
		log.Fatalf("Failed to drop all tables: %v", err)
		return err
	}

	log.Println("All tables dropped successfully")
	return nil
}

func SeedData(db *gorm.DB) error {
	// ダミーユーザーを作成
	users := []models.User{
		{UserName: "Alice", Email: "alice@example.com", Password: "password123", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserName: "Bob", Email: "bob@example.com", Password: "password123", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return fmt.Errorf("failed to insert user %v: %w", user.UserName, err)
		}
	}

	// ダミーオブジェクトを作成
	objs := []models.Obj{
		{UserID: 1, ObjTitle: "First Object", DeleteFlag: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserID: 2, ObjTitle: "Second Object", DeleteFlag: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, obj := range objs {
		if err := db.Create(&obj).Error; err != nil {
			return fmt.Errorf("failed to insert obj %v: %w", obj.ObjTitle, err)
		}
	}

	return nil
}
