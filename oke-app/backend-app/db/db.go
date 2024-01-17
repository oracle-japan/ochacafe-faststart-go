package db

import (
	"fmt"
	"os"
	"time"

	"github.com/oracle-japan/ochacafe-faststart-go/oke-app/backend-app/repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDbInfo() string {
	dbInfo := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	return dbInfo
}

func GetDBInfo() *gorm.DB {
	dbInfo := getDbInfo()

	db, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		panic("Failed to open database" + err.Error())
	}

	return db
}

func SetupDB() {

	db := GetDBInfo()
	db.AutoMigrate(&repo.Items{})

	Items := []*repo.Items{
		{
			Name:       "速習Golang",
			Date:       time.Date(2024, 2, 7, 19, 00, 00, 000000, time.UTC).Format("20060102150405"),
			Topics:     "Golang",
			Presenters: "Takuya Niita",
		},
		{
			Name:       "基礎から学ぶNext.js",
			Date:       "TBD",
			Topics:     "SSR, SSG, React",
			Presenters: "Kyotaro Nonaka",
		},
		{
			Name:       "シン・Kafka",
			Date:       "TBD",
			Topics:     "kRaft, Strimzi Operator",
			Presenters: "Shuhei Kawamura",
		},
		{
			Name:       "NewSQLのランドスケープ",
			Date:       "TBD",
			Topics:     "TiDB, CockroachDB",
			Presenters: "Yutaka Ichikawa",
		},
		{
			Name:       "Kubernetesで作るIaaS基盤",
			Date:       "TBD",
			Topics:     "KubeVirt, VM, libvirt, QEMU",
			Presenters: "Takuya Niita",
		},
		{
			Name:       "LLMのエコシステム",
			Date:       "TBD",
			Topics:     "Langchain, LlamaIndex, Vector Database",
			Presenters: "Shuhei Kawamura",
		},
	}

	db.Exec("TRUNCATE items")

	db.Create(Items)
}
