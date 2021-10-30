package theater

import (
	"context"
	"encoding/json"
	"fmt"
	"ticketing/business/theater"
	"ticketing/helper/converter"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type mysqlTheaterRepository struct {
	Conn  *gorm.DB
	Redis *redis.Client
}

func NewMySQLTheaterRepository(conn *gorm.DB, redis *redis.Client) theater.Repository {
	return &mysqlTheaterRepository{
		Conn:  conn,
		Redis: redis,
	}
}

func (repository *mysqlTheaterRepository) Store(ctx context.Context, theaterDomain *theater.Domain) error {
	//! SAVE TO DB
	rec := fromDomain(*theaterDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	//! SAVE TO REDIS
	var recFind []Theater
	var allTheater []theater.Domain

	result2 := repository.Conn.Find(&recFind)
	if result2.Error != nil {
		return result2.Error
	}

	for _, value := range recFind {
		allTheater = append(allTheater, value.toDomain())
	}

	val, err := converter.ConvertStructToString(allTheater)
	if err != nil {
		fmt.Println("cannot marshal struct to string")
	}

	value, err := repository.Redis.Get("GetAll_Theater").Result()
	
	if val != value {
		err = repository.Redis.Set("GetAll_Theater", val, 0).Err()
		if err != nil {
			fmt.Println("Redis error set: ", err)
		}

		return nil
	} else if err != nil {
		fmt.Println("Redis error get: ", err)
	}

	return nil
}

func (repository *mysqlTheaterRepository) Delete(ctx context.Context, id int) error {
	theaterDelete := Theater{}
	result := repository.Conn.Where("id = ?", id).Delete(&theaterDelete)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlTheaterRepository) Update(ctx context.Context, theaterDomain *theater.Domain, id int) error {
	rec := fromDomain(*theaterDomain)

	result := repository.Conn.Where("id = ?", id).Updates(&rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlTheaterRepository) GetAll(ctx context.Context) ([]theater.Domain, error) {
	var allTheater []theater.Domain
	
	value, err := repository.Redis.Get("GetAll_Theater").Result()
	if err != nil {
		fmt.Println("Redis error get: ", err)
	}
	json.Unmarshal([]byte(value), &allTheater)

	return allTheater, nil
}
