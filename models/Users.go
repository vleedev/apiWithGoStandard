package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx"
)
/*
*	The model for users
*	Author: vlee.dev
 */
type User struct {
	ID			primitive.ObjectID 		`bson:"_id,omitempty"`
	Email  					string		`json:"email,omitempty"`
	FirstName				string		`json:"firstName,omitempty"`
	LastName				string		`json:"lastName,omitempty"`
	Password				string		`json:"password,omitempty"`
	Avatar					string		`json:"avatar,omitempty"`
	Telephone				uint8		`json:"telephone,omitempty"`
	Birthday				uint32		`json:"birthday,omitempty"`
	Facebook				string		`json:"facebook,omitempty"`
	Address					string		`json:"address,omitempty"`
	LocationLat				float64		`json:"locationLat,omitempty"`
	LocationLon				float64		`json:"locationLon,omitempty"`
	ResetToken				string		`json:"resetToken,omitempty"`
	VeriEmailToken			string		`json:"veriEmailToken,omitempty"`
	VeriPhoneToken			string		`json:"veriPhoneToken,omitempty"`
	Language				string		`json:"language,omitempty"`
	MainRole	primitive.ObjectID		`bson:"mainRole,omitempty"`
	Roles					bsonx.Arr	`bson:"roles,omitempty"`
	ChatBoxes				bsonx.Arr	`bson:"chatBoxes,omitempty"`
	UserActivities			bsonx.Arr	`bson:"userActivities,omitempty"`
	CreatedAt				int64		`json:"createdAt,omitempty"`
	UpdatedAt				int64		`json:"updatedAt,omitempty"`
	VerifiedByEmail			bool		`json:"verifiedByEmail" default:"false"`
	VerifiedByPhone			bool		`json:"verifiedByPhone" default:"false"`
}