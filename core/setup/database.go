package setup

import (
	"context"
	"errors"
	"fmt"
	// "git.internal.attains.cn/attains-cloud/service-acs/core/config/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

var (
	ErrDatabaseNameEmpty         = errors.New("database name cannot be empty")
	ErrDatabaseDriverUnsupported = errors.New("database driver unsupported")
)

func InitDatabase(_ context.Context, database *pkg.Database) (*gorm.DB, error) {
	database.Name = strings.TrimSpace(database.Name)
	if database.Name == "" {
		return nil, ErrDatabaseNameEmpty
	}

	switch database.Driver {
	case "mysql":
		dsn := database.Dsn
		if database.Dsn == "" {
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
				database.User,
				database.Pwd,
				database.Host,
				database.Port,
				database.DbName,
				database.Charset,
				true,
				"Local",
			)
		}
		orm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			SkipDefaultTransaction:                   false,
			NamingStrategy:                           nil,
			FullSaveAssociations:                     false,
			Logger:                                   nil,
			NowFunc:                                  nil,
			DryRun:                                   false,
			PrepareStmt:                              false,
			DisableAutomaticPing:                     false,
			DisableForeignKeyConstraintWhenMigrating: false,
			IgnoreRelationshipsWhenMigrating:         false,
			DisableNestedTransaction:                 false,
			AllowGlobalUpdate:                        false,
			QueryFields:                              false,
			CreateBatchSize:                          0,
			TranslateError:                           false,
			ClauseBuilders:                           nil,
			ConnPool:                                 nil,
			Dialector:                                nil,
			Plugins:                                  nil,
		})
		if err != nil {
			return nil, err
		}
		return orm, nil
	default:
		return nil, ErrDatabaseDriverUnsupported
	}
}
