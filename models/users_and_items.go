package models

type User struct{
  name string 'json:"user_name" form:"user_name"'
  userID string 'json:"userID form:"userID'
}

type Product struct{
  name string 'json:"product_name" form:"product_name"'
  productID string 'json:"productID" form:"productID"'
}
