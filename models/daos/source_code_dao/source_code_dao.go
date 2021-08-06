package source_code_dao

import (
	"context"
	"sync"

	"github.com/tranHieuDev23/IdeTwo/models/daos/collections"
	"github.com/tranHieuDev23/IdeTwo/models/source_code"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SourceCodeDao struct {
	collections collections.DbCollections
}

var instance *SourceCodeDao = nil
var once sync.Once

func GetInstance() SourceCodeDao {
	once.Do(func() {
		collections := collections.GetInstance()
		instance = &SourceCodeDao{
			collections: collections,
		}
	})
	return *instance
}

func (dao *SourceCodeDao) CreateSourceCode(source source_code.SourceCode) {
	_, err := dao.collections.SourceCodeCollection.InsertOne(context.TODO(), source)
	if err != nil {
		panic(err)
	}
}

func (dao *SourceCodeDao) GetSourceCode(id string) *source_code.SourceCode {
	filters := bson.D{{Key: "id", Value: id}}
	result := dao.collections.SourceCodeCollection.FindOne(context.TODO(), filters)
	if result.Err() != nil {
		return nil
	}
	source := source_code.SourceCode{}
	result.Decode(&source)
	return &source
}

func (dao *SourceCodeDao) UpdateSourceCode(source source_code.SourceCode) *source_code.SourceCode {
	filters := bson.D{{Key: "id", Value: source.Id}}
	result := dao.collections.SourceCodeCollection.FindOneAndReplace(
		context.TODO(),
		filters,
		source,
		options.FindOneAndReplace().SetReturnDocument(options.After))
	if result.Err() != nil {
		return nil
	}
	newSource := source_code.SourceCode{}
	result.Decode(&newSource)
	return &newSource
}

func (dao *SourceCodeDao) DeleteSourceCode(id string) *source_code.SourceCode {
	filters := bson.D{{Key: "id", Value: id}}
	result := dao.collections.SourceCodeCollection.FindOneAndDelete(context.TODO(), filters)
	if result.Err() != nil {
		return nil
	}
	newSource := source_code.SourceCode{}
	result.Decode(&newSource)
	return &newSource
}
