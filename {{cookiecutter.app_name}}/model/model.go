package model

import (
	"fmt"
	"os"

	{% if cookiecutter.use_viper_config == "y" %}"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/config"{% endif %}
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB
var err error

func init() {
	{% if cookiecutter.use_viper_config == "y" %}
	dialect := config.GetString("dialect")
	dbPath := config.GetString("db_path")
	dbLogMode := config.GetBool("db_log_mode")
	{% else %}
	dialect := "mysql"
	dbPath := "127.0.0.1:3306"
	dbLogMode := false
	{% endif %}

	DB, err = gorm.Open(dialect, dbPath)
	if err != nil {
		fmt.Printf("err:%s\n", err)
		os.Exit(1)
	}

	DB.LogMode(dbLogMode)
	// DB.AutoMigrate(&User{})
	// DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
}
