package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"hotel-rooms/constants"
	"hotel-rooms/helpers"
	"hotel-rooms/internal/models"
	"time"
)

type RoomsRepository struct {
	DB    *gorm.DB
	Redis *redis.ClusterClient
}

func NewRoomsRepository(db *gorm.DB, rdb *redis.ClusterClient) *RoomsRepository {
	return &RoomsRepository{
		DB:    db,
		Redis: rdb,
	}
}

func (r *RoomsRepository) AddRoom(ctx context.Context, room *models.Room) error {
	return r.DB.WithContext(ctx).Create(room).Error
}

func (r *RoomsRepository) GetAllRooms(ctx context.Context) ([]models.Room, error) {
	var rooms []models.Room

	result, err := r.Redis.Get(ctx, constants.RedisKeyAllRooms).Result()
	if err == nil && result != "" {
		if err = json.Unmarshal([]byte(result), &rooms); err != nil {
			helpers.Logger.Warn("failed to unmarshal rooms data")
			return nil, err
		}

		helpers.Logger.Info("success get data all rooms from redis")
		return rooms, nil
	}

	subquery := r.DB.Table("rooms").
		Select("MIN(id) as id").
		Group("room_type_id")

	err = r.DB.WithContext(ctx).
		Joins("JOIN (?) as r2 ON rooms.id = r2.id", subquery).
		Preload("RoomType.RoomPhotos").Preload("RoomType.RoomFeatures").
		Find(&rooms).Error

	if err != nil {
		return nil, err
	}

	go func() {
		ctx = context.Background()
		roomsJson, err := json.Marshal(rooms)
		if err != nil {
			helpers.Logger.Warn("error marshal rooms data to json")
			return
		}

		err = r.Redis.Set(ctx, constants.RedisKeyAllRooms, string(roomsJson), time.Hour*24).Err()
		if err != nil {
			helpers.Logger.Warn("failed to set all rooms to redis: ", err)
			return
		}

		helpers.Logger.Info("success set data all rooms to redis")
	}()

	return rooms, nil
}

func (r *RoomsRepository) GetRoomDetails(ctx context.Context, id int) (*models.Room, error) {
	var room models.Room

	result, err := r.Redis.Get(ctx, fmt.Sprintf(constants.RedisKeyDetailsRoom, string(rune(id)))).Result()
	if err == nil && result != "" {
		if err = json.Unmarshal([]byte(result), &room); err != nil {
			return nil, err
		}

		helpers.Logger.Info("success get room details from redis")
		return &room, nil
	}

	err = r.DB.WithContext(ctx).Where("id = ?", id).Preload("RoomType.RoomPhotos").Preload("RoomType.RoomFeatures").First(&room).Error
	if err != nil {
		return nil, err
	}

	go func() {
		ctx = context.Background()
		roomJson, err := json.Marshal(room)
		if err != nil {
			helpers.Logger.Warn("error marshal room data to json")
			return
		}

		err = r.Redis.Set(ctx, fmt.Sprintf(constants.RedisKeyDetailsRoom, string(rune(id))), string(roomJson), time.Hour*24).Err()
		if err != nil {
			helpers.Logger.Warn("failed to set room details to redis: ", err)
			return
		}
	}()

	return &room, nil
}

func (r *RoomsRepository) EditRoom(ctx context.Context, roomID int, newData map[string]interface{}) error {
	err := r.DB.WithContext(ctx).Model(&models.Room{}).Where("id = ?", roomID).Updates(newData).Error
	if err != nil {
		return err
	}

	go func() {
		ctx = context.Background()

		err = r.Redis.Del(ctx, constants.RedisKeyAllRooms).Err()
		if err != nil {
			helpers.Logger.Warn("failed to delete all rooms from redis: ", err)
			return
		}

		err = r.Redis.Del(ctx, fmt.Sprintf(constants.RedisKeyDetailsRoom, string(rune(roomID)))).Err()
		if err != nil {
			helpers.Logger.Warn("failed to delete room details from redis: ", err)
			return
		}

		helpers.Logger.Info("success delete room from redis")
	}()

	return nil
}

func (r *RoomsRepository) DeleteRoom(ctx context.Context, roomID int) error {
	err := r.DB.WithContext(ctx).Where("id = ?", roomID).Delete(&models.Room{}).Error
	if err != nil {
		return err
	}

	go func() {
		ctx = context.Background()

		err = r.Redis.Del(ctx, constants.RedisKeyAllRooms).Err()
		if err != nil {
			helpers.Logger.Warn("failed to delete all rooms from redis: ", err)
			return
		}

		err = r.Redis.Del(ctx, fmt.Sprintf(constants.RedisKeyDetailsRoom, string(rune(roomID)))).Err()
		if err != nil {
			helpers.Logger.Warn("failed to delete room details from redis: ", err)
			return
		}

		helpers.Logger.Info("success delete room from redis")
	}()

	return nil
}

func (r *RoomsRepository) GetRoomAvailability(ctx context.Context, totalBooked []models.RoomBookedResponse) ([]*models.Room, error) {
	var rooms []*models.Room

	result, err := r.Redis.Get(ctx, constants.RedisKeyRoomsAvailable).Result()
	if err == nil && result != "" {
		if err = json.Unmarshal([]byte(result), &rooms); err != nil {
			helpers.Logger.Warn("failed to unmarshal rooms data")
			return nil, err
		}

		helpers.Logger.Info("success get data rooms available from redis")
		return rooms, nil
	}

	subquery := r.DB.Table("rooms").
		Select("MIN(id) as id").
		Group("room_type_id")

	err = r.DB.WithContext(ctx).
		Joins("JOIN (?) as r2 ON rooms.id = r2.id", subquery).
		Preload("RoomType.RoomPhotos").Preload("RoomType.RoomFeatures").
		Order("rooms.room_type_id ASC").
		Find(&rooms).Error

	if err != nil {
		return nil, err
	}

	for i, room := range rooms {
		if int(room.RoomType.ID) == totalBooked[i].RoomTypeID && room.RoomType.TotalRooms-totalBooked[i].TotalBooked >= 0 {
			room.Available = true
		} else {
			room.Available = false
		}
	}

	go func() {
		ctx = context.Background()
		roomsJson, err := json.Marshal(rooms)
		if err != nil {
			helpers.Logger.Warn("error marshal rooms available data to json")
			return
		}

		err = r.Redis.Set(ctx, constants.RedisKeyRoomsAvailable, string(roomsJson), time.Hour*24).Err()
		if err != nil {
			helpers.Logger.Warn("failed to set all rooms available to redis: ", err)
			return
		}

		helpers.Logger.Info("success set data rooms available to redis")
	}()

	return rooms, nil
}
