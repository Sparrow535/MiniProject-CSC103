package model

import (
	"database/sql"
	"errors"
	"myapp/dataBase/postgres"
)

type Room struct {
	Block  string `json:"block"`
	Floor  string `json:"floor"`
	RoomNo string `json:"room_no"`
}
type RoomG struct {
	Block  string `json:"block"`
	Floor  string `json:"floor"`
	RoomNo string `json:"room_no"`
}

const queryInsertRoom = "INSERT INTO room(block, floor, room_no) VALUES ($1, $2, $3) RETURNING block, floor, room_no;"

func (room *Room) CreateRoom() error {
	row := postgres.Db.QueryRow(queryInsertRoom, room.Block, room.Floor, room.RoomNo)
	err := row.Scan(&room.Block, &room.Floor, &room.RoomNo)
	return err
}

func GetRoomByNumber(roomNo string) (*Room, error) {
	var room Room

	// Query the database to fetch room details
	err := postgres.Db.QueryRow("SELECT room_no, block, floor FROM room WHERE room_no = $1", roomNo).Scan(&room.RoomNo, &room.Block, &room.Floor)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("room not found")
		}
		return nil, err
	}

	return &room, nil
}

const queryInsertRoomG = "INSERT INTO roomg(block, floor, room_no) VALUES ($1, $2, $3) RETURNING block, floor, room_no;"

func (room *RoomG) CreateRoomG() error {
	row := postgres.Db.QueryRow(queryInsertRoomG, room.Block, room.Floor, room.RoomNo)
	err := row.Scan(&room.Block, &room.Floor, &room.RoomNo)
	return err
}

func GetRoomByNumberG(roomNo string) (*RoomG, error) {
	var room RoomG

	// Query the database to fetch room details
	err := postgres.Db.QueryRow("SELECT room_no, block, floor FROM roomg WHERE room_no = $1", roomNo).Scan(&room.RoomNo, &room.Block, &room.Floor)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("room not found")
		}
		return nil, err
	}

	return &room, nil
}
