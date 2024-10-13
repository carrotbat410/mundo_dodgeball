package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	RoomId               primitive.ObjectID `json:"room_id" bson:"_id,omitempty"`
	Title                string             `json:"title" bson:"title"`
	Id                   string             `json:"id" bson:"id"`                                         //방장 아이디
	Username             string             `json:"username" bson:"username,omitempty"`                   //방장 이름
	Content              string             `json:"content" bson:"content"`                               //방 설명
	Password             string             `json:"password,omitempty" bson:"password"`                   //방 비밀번호
	IsStart              bool               `json:"is_start" bson:"is_start"`                             //게임 시작했는지
	MaxParticipantNumber int                `json:"max_participant_number" bson:"max_participant_number"` //최대 참가자 수
	NowParticipantNumber int                `json:"now_participant_number" bson:"now_participant_number"` //현재 참가자 수
	ParticipantIdList    []string           `json:"participant_id_list" bson:"participant_id_list"`
	CreatedAt            time.Time          `bson:"created_at"`
	UpdatedAt            time.Time          `bson:"updated_at"`
	IsDelete             bool               `json:"is_delete" bson:"is_delete"` //softdelete되었는지
}
