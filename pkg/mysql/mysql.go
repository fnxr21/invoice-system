package mysql

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func DataBaseinit() {
	var err error

	// env database relation
	var dbconfig = os.Getenv("DB_CHARNTIME")
	var mst = os.Getenv("DB_MST")

	// env DATABASE HOST
	var dbuser = os.Getenv("DB_USER")
	var dbpass = os.Getenv("DB_PASS")
	var dbhost = os.Getenv("DB_HOST")
	var dbport = os.Getenv("DB_PORT")

	
	var mysqlconfig = dbuser + ":" + dbpass + "@tcp(" + dbhost + ":" + dbport + ")/"
	var gormConfig = &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}}
		
		fmt.Println(mysqlconfig,mst,dbconfig)
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s%s%s", mysqlconfig, mst, dbconfig)), gormConfig)
	if err != nil {
		fmt.Println("connect to database mst failed")
		panic(err)
	}

	// Retrieve the underlying sql.DB and configure the connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("get underlying sql.DB failed")
		panic(err)
	}
	// Set the connection pool parameters
	sqlDB.SetConnMaxIdleTime(30 * time.Minute) // max idle time  30 minuter
	sqlDB.SetConnMaxLifetime(2 * time.Hour)    //max life time conenection 2 hours
	sqlDB.SetMaxIdleConns(50)                  //max 50 idle connection
	sqlDB.SetMaxOpenConns(300)                 // max handle requst 300

	// scenario
	//High traffic (above 500 RPS)	,SetMaxOpenConns(500), SetMaxIdleConns(100)
	//Too many idle connections	 Decrease SetMaxIdleConns(30)
	//Too many accumulated connections	 Reduce SetMaxOpenConns(200)
	// Frequent stale connection errors	 Lower SetConnMaxLifetime(1 * time.Hour)

	// Ensure the database connection is alive
	// if err := sqlDB.Ping(); err != nil {
	// 	fmt.Println("Database ping failed:", err)
	// 	panic(err)
	// }

	fmt.Println("connected to database")
}
