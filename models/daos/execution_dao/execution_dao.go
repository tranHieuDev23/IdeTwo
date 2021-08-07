package execution_dao

import (
	"context"
	"sync"

	"github.com/tranHieuDev23/IdeTwo/models/daos/collections"
	"github.com/tranHieuDev23/IdeTwo/models/execution"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ExecutionDao struct {
	collections collections.DbCollections
}

var instance *ExecutionDao = nil
var once sync.Once

func GetInstance() ExecutionDao {
	once.Do(func() {
		collections := collections.GetInstance()
		instance = &ExecutionDao{
			collections: collections,
		}
	})
	return *instance
}

func (dao *ExecutionDao) CreateExecution(execution execution.Execution) {
	_, err := dao.collections.ExecutionCollection.InsertOne(context.TODO(), execution)
	if err != nil {
		panic(err)
	}
}

func (dao *ExecutionDao) GetExecution(id string) *execution.Execution {
	filters := bson.D{{Key: "id", Value: id}}
	doc := dao.collections.ExecutionCollection.FindOne(context.TODO(), filters)
	if doc.Err() != nil {
		return nil
	}
	result := execution.Execution{}
	doc.Decode(&result)
	return &result
}

func (dao *ExecutionDao) UpdateExecution(exec execution.Execution) *execution.Execution {
	filters := bson.D{{Key: "id", Value: exec.Id}}
	doc := dao.collections.ExecutionCollection.FindOneAndReplace(
		context.TODO(),
		filters,
		exec,
		options.FindOneAndReplace().SetReturnDocument(options.After))
	if doc.Err() != nil {
		return nil
	}
	result := execution.Execution{}
	doc.Decode(&result)
	return &result
}

func (dao *ExecutionDao) DeleteExecution(id string) *execution.Execution {
	filters := bson.D{{Key: "id", Value: id}}
	doc := dao.collections.ExecutionCollection.FindOneAndDelete(context.TODO(), filters)
	if doc.Err() != nil {
		return nil
	}
	result := execution.Execution{}
	doc.Decode(&result)
	return &result
}
