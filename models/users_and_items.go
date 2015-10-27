package models

type User struct{
  name string 'json:"user_name"'
  userID string 'json:"userID"'
}

type Product struct{
  name string 'json:"product_name"'
  productID string 'json:"productID"'
}
