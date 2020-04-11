package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"otm/app/constants"
	"otm/app/error_package"
	"otm/app/providers/database"
	"otm/app/utils"
	"strings"
	"time"
)

const (
	Table     = "table"
	Entity    = "entity"
	Condition = "condition"
)

type Base struct{}

type LastUpdatedAtResult struct {
	LoginId       int
	LastUpdatedAt time.Time
	LastDeletedAt time.Time
}

//
//func (repo Base) GetLastUpdatedAt(ctx *gin.Context, namespaceIds []int) (map[int]time.Time, error) {
//	var results []LastUpdatedAtResult
//	resultSet := make(map[int]time.Time)
//	sql := "SELECT max(last_updated_at) as last_updated_at, max(last_deleted_at) as last_deleted_at,namespace_id FROM (SELECT max(updated_at) as last_updated_at, max(deleted_at) as last_deleted_at,namespace_id FROM domainmodel WHERE namespace_id in (?) GROUP BY namespace_id UNION SELECT max(updated_at) as last_updated_at, max(deleted_at) as last_deleted_at,namespace_id FROM rulechain WHERE namespace_id in (?) GROUP BY namespace_id UNION SELECT max(updated_at) as last_updated_at, max(deleted_at) as last_deleted_at,namespace_id FROM rules WHERE namespace_id in (?) GROUP BY namespace_id UNION SELECT max(updated_at) as last_updated_at, max(deleted_at) as last_deleted_at,namespace_id FROM rule_groups WHERE namespace_id in (?) GROUP BY namespace_id) AS X GROUP BY namespace_id;"
//	err := getDbInstance(ctx).Raw(sql, namespaceIds, namespaceIds, namespaceIds, namespaceIds).Scan(&results).Error
//
//	for _, result := range results {
//		if result.LastDeletedAt.After(result.LastUpdatedAt) {
//			resultSet[result.LoginId] = result.LastDeletedAt
//		} else {
//			resultSet[result.LoginId] = result.LastUpdatedAt
//		}
//	}
//
//	return resultSet, getErrorResponse(ctx, err, map[string]interface{}{
//		"table":         models.RuleTableName,
//		"namespace_ids": namespaceIds,
//	})
//}

func (repo Base) RunQuery(ctx *gin.Context, query string, arguments []interface{}) error {
	err := repo.GetDbInstance(ctx).Exec(query, arguments...).Error
	return getErrorResponse(ctx, err, map[string]interface{}{
		"query":     query,
		"arguments": arguments,
	})
}

// Find finds a model based on given conditions
func (repo Base) Find(ctx *gin.Context, model interface{}, condition map[string]interface{}) error {
	err := repo.GetDbInstance(ctx).Where(condition).First(model).Error
	gorm.IsRecordNotFoundError(err)
	return getErrorResponse(ctx, err, map[string]interface{}{
		"table":     utils.GetTypeName(model),
		"entity":    model,
		"condition": condition,
	})
}

// Find finds all models based on given conditions
func (repo Base) FindAll(ctx *gin.Context, model interface{}, condition map[string]interface{}) error {
	err := repo.GetDbInstance(ctx).Where(condition).Order("created_at desc").Find(model).Error

	return getErrorResponse(ctx, err, map[string]interface{}{
		"table":     utils.GetTypeName(model),
		"entity":    model,
		"condition": condition,
	})
}

func (repo Base) FindTrending(ctx *gin.Context, model interface{}, condition map[string]interface{}) error {
	err := repo.GetDbInstance(ctx).Where(condition).Order("applause desc").Find(model).Error

	return getErrorResponse(ctx, err, map[string]interface{}{
		"table":     utils.GetTypeName(model),
		"entity":    model,
		"condition": condition,
	})
}

func (repo Base) Update(ctx *gin.Context, model interface{}, attributes map[string]interface{}, condition map[string]interface{}) error {
	err := repo.GetDbInstance(ctx).Where(condition).First(model).Update(attributes).Error

	return getErrorResponse(ctx, err, map[string]interface{}{
		"table":     utils.GetTypeName(model),
		"entity":    model,
		"condition": condition,
	})

}

func (repo Base) Updates(ctx *gin.Context, model interface{}) error {
	err := repo.GetDbInstance(ctx).Model(model).Updates(model).Error

	return getErrorResponse(ctx, err, map[string]interface{}{
		"table":  utils.GetTypeName(model),
		"entity": model,
	})
}

// Create will insert the data specified by the entity
// Taking entity as interface as we can send any struct and perform an insert on it
func (repo Base) Create(ctx *gin.Context, model interface{}) error {
	err := repo.GetDbInstance(ctx).Create(model).Error

	return getErrorResponse(ctx, err, map[string]interface{}{
		"table":  utils.GetTypeName(model),
		"entity": model,
		"err":    err,
	})
}

// This will delete the entities if they exist
// and if no record exist then it does nothing. No error in that case either.
func (repo Base) Delete(ctx *gin.Context, model interface{}, condition map[string]interface{}) error {
	err := repo.GetDbInstance(ctx).Where(condition).Delete(model).Error

	return getErrorResponse(ctx, err, map[string]interface{}{
		"table":     utils.GetTypeName(model),
		"entity":    model,
		"condition": condition,
	})
}

// This will delete the entities if they exist
// and if no record exist then it does nothing. No error in that case either.
func (repo Base) UnscopedDelete(ctx *gin.Context, model interface{}, condition map[string]interface{}) error {
	err := repo.GetDbInstance(ctx).Where(condition).Unscoped().Delete(model).Error

	return getErrorResponse(ctx, err, map[string]interface{}{
		"table":     utils.GetTypeName(model),
		"entity":    model,
		"condition": condition,
	})
}

// Creates the record if the data don't exist for the given condition
// If the data exist then the record will be updated
func (repo Base) FirstOrCreate(ctx *gin.Context, model interface{}, condition map[string]interface{}) error {
	err := repo.GetDbInstance(ctx).
		Where(condition).
		FirstOrCreate(model).Error

	return getErrorResponse(ctx, err, map[string]interface{}{
		"table":     utils.GetTypeName(model),
		"model":     model,
		"condition": condition,
	})
}

// Find finds a model based on given conditions
func (repo Base) Last(ctx *gin.Context, model interface{}, condition map[string]interface{}) error {
	err := repo.GetDbInstance(ctx).Where(condition).Last(model).Error
	gorm.IsRecordNotFoundError(err)
	return getErrorResponse(ctx, err, map[string]interface{}{
		"table":     utils.GetTypeName(model),
		"entity":    model,
		"condition": condition,
	})
}

// GetDbInstance will return the db instance if available
// else returns set the new db instance in context and return
func (repo Base) GetDbInstance(ctx *gin.Context) *gorm.DB {
	db, ok := ctx.Get(constants.DB)

	if !ok || db == nil {
		db = database.GetClient().Instance()
	}

	return db.(*gorm.DB)
}

// Transaction will run the handlers in transaction
// Roll back the transaction if any error occurred in transaction
// commits if there are no errors

// SetDB will set the db instance in the context
// TODO: find why *gorm.DB throws nil pointer panic
func (repo Base) setTransactionDB(ctx *gin.Context, db interface{}) {
	ctx.Set(constants.DB, db)
}

func getErrorResponse(ctx *gin.Context, err error, data map[string]interface{}) error {
	if err != nil {
		message := err.Error()
		errorCode := error_package.CodeDatabaseError

		if strings.Contains(message, "Duplicate entry") {
			errorCode = error_package.UserAlreadyExisted
		}

		if gorm.IsRecordNotFoundError(err) {
			errorCode = error_package.RecordNotFound
		}

		fmt.Println(err, errorCode)
		return err
	}

	return nil
}

// Transaction will run the handlers in transaction
// Roll back the transaction if any error occurred in transaction
// commits if there are no errors
func (repo Base) Transaction(ctx *gin.Context, handlers ...func() (interface{}, error)) (data interface{}, ierr error) {
	db := repo.GetDbInstance(ctx).Begin()
	repo.setTransactionDB(ctx, db)

	defer func() {
		if rec := recover(); rec != nil {
			err := utils.GetError(rec)

			ierr = err
		}

		// var.(type) can only be used with switch
		switch ierr.(type) {
		case nil:
			err := db.Commit().Error

			if err != nil {
				ierr = err
			}

		default:
			err := db.Rollback().Error

			if err != nil {
				ierr = err
			}
		}

		// Reset the db instance in the context
		repo.setTransactionDB(ctx, nil)
	}()

	for _, handler := range handlers {
		data, ierr = handler()

		if ierr != nil {
			return data, ierr
		}
	}

	return data, ierr
}

// GetDbInstance will return the db instance if available
// else returns set the new db instance in context and return
func getDbInstance(ctx *gin.Context) *gorm.DB {
	db, ok := ctx.Get(constants.DB)

	if !ok || db == nil {
		db = database.GetClient().Instance()
	}

	return db.(*gorm.DB)
}
