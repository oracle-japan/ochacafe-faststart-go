package db

import (
	"fmt"
	"os"
	"time"

	"github.com/oracle-japan/ochacafe-faststart-go/oke-app/backend-app/repo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDbInfo() string {
	dbInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	return dbInfo
}

func SetupDB() {
	dbInfo := GetDbInfo()
	db, err := gorm.Open(mysql.Open(dbInfo), &gorm.Config{})
	if err != nil {
		panic("Failed to open database" + err.Error())
	}

	db.Exec("TRUNCATE items")

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

	db.Create(Items)
}
