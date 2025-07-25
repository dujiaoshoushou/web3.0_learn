package utils

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"strings"
	"time"
)

var db *gorm.DB

func InitDB() {
	connectionString := viper.GetString("db.dsn")
	logWriter = LogWriter()
	loggLevel := gormLogger.Warn
	switch strings.ToUpper(viper.GetString("app.mode")) {
	case APP_RELEASE:
		loggLevel = gormLogger.Warn
	case APP_DEBUG, APP_TEST:
		fallthrough
	default:
		loggLevel = gormLogger.Info
	}
	customerLogger := gormLogger.New(log.New(logWriter, "\r\n", log.LstdFlags), gormLogger.Config{
		SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
		LogLevel:                  loggLevel,              // 日志级别
		Colorful:                  false,                  // 禁用彩色打印
		IgnoreRecordNotFoundError: false,
		ParameterizedQueries:      false,
	})
	conf := &gorm.Config{
		SkipDefaultTransaction: false,
		// NamingStrategy tables, columns naming strategy
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		// FullSaveAssociations full save associations
		//FullSaveAssociations: false,
		// Logger
		Logger: customerLogger,

		//// DryRun generate sql without execute
		//DryRun: false,
		//// PrepareStmt executes the given query in cached statement
		//PrepareStmt: false,
		//
		//// DisableAutomaticPing
		//DisableAutomaticPing: false,
		//// DisableForeignKeyConstraintWhenMigrating
		//DisableForeignKeyConstraintWhenMigrating: true,
		//// IgnoreRelationshipsWhenMigrating
		//IgnoreRelationshipsWhenMigrating: false,
		//// DisableNestedTransaction disable nested transaction
		//DisableNestedTransaction: false,
		//// AllowGlobalUpdate allow global update
		//AllowGlobalUpdate: false,
		//// QueryFields executes the SQL query with all fields of the table
		//QueryFields: false,
		//// CreateBatchSize default create batch size
		//CreateBatchSize: 0,
		//// TranslateError enabling error translation
		//TranslateError: false,
		//// PropagateUnscoped propagate Unscoped to every other nested statement
		//PropagateUnscoped: false,
		//
		//// ClauseBuilders clause builder
		//ClauseBuilders: nil,
		//// ConnPool db conn pool
		//ConnPool: nil,
		//// Dialector database dialector
		//Dialector: nil,
		//// Plugins registered plugins
		//Plugins: nil,
	}

	newDb, err := gorm.Open(mysql.Open(connectionString), conf)
	if err != nil {
		log.Fatal(err)
	} else {
		db = newDb
	}
}

func DB() *gorm.DB {
	return db
}
