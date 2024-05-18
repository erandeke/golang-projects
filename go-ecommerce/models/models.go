package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name      *string            `json:"first_name" validate:"required,min=2 max=30"`
	Last_Name       *string            `json:"last_name"  validate:"required,min=2,max=30"`
	Password        *string            `json:"password"   validate:"required, min=6"`
	Phone           *int               `json:"phone"`
	Email           *string            `json:"email"`
	Token           *string            `json:"token"`
	Refresh_Token   *string            `json:"refresh_token"`
	User_Id         string             `json:"user_id"`
	UserCart        []ProductUser      `json:"usercart" bson:"ProductUser"`
	Created_At      time.Time          `json:"createdAt"`
	Updated_At      time.Time          `json:"UpdatedAt"`
	Address_Details []Address          `json:"AddressDetails" bson:"Address"`
	Order_Status    []Order            `json:"order" bson:"orders"` // should this be just order , why do we need slices
}

type Product struct {
	Product_Id   primitive.ObjectID `json:"_id" bson:"_id"`
	Product_Name *string            `json:"product_name"`
	Price        *uint64            `json:"price"` //allows only positive
	Rating       *uint8             `json:"rating"`
	Image        *string            `json:"Image"` //store the url of image
}

type ProductUser struct {
	Product_Id   primitive.ObjectID `json:"_id" bson:"_id"`
	Product_Name *string            `json:"product_name"`
	Price        *uint64            `json:"price"`
	Rating       *uint8             `json:"rating"`
	Image        *string            `json:"Image"`
}

type Address struct {
	Address_ID primitive.ObjectID `json:"_id" bson:"_id"`
	Street     *string            `json:"street"`
	PinCode    *string            `json:"pincode"`
}

type Order struct {
	Order_Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Order_Cart     []ProductUser      `json:"order_cart" bson:"productuser"`
	Ordered_At     time.Time          `json:"orderd_at"`
	Price          *uint64            `json:"price"`
	Discount       *uint64            `json:"discount"`
	Payment_Method Payment            `json:"payment_method" bson:"payment"`
}

type Payment struct {
	Digital bool
	COD     bool
}
