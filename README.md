<h1>E - commerce app</h1>

<ul>
    <li>Home Page</li>
    <li>Register, Login</li>
    <li>Product</li>
</ul>

<h3>Technology Stack</h3>
<ul>
    <li>Golang
        <ul>
            <li>Gorm</li>
            <li>Gin</li>
        </ul>
    </li>
    <li>PostgreSQL</li>
    <li>Docker</li>
</ul>

<ul> User :
<li>Login / Auth</li>
<li>Comment Product</li>
<li>Rate Product</li>
<li>Search Item</li>
<li>Purchasing Item</li>
<li>Filter Item by price & rating</li>
</ul>

* <h4> SignUp API</h4> 
 POST http://localhost:8181/signup

`{
    "username": "user2",
    "email address": "email2",
    "password": "abscdeffege"
}`

* <h4> LogIn API </h4> 
POST http://localhost:8181/login

`{
"email": "email1",
"password": "abscdeffege"
}`

* <h4> Publish Item </h4> 
POST http://localhost:8181/user/items

`{
"name": "item3",
"description": "desc item3",
"price": 28.3
}`

* <h4> Filter Item by price & rating</h4>
GET http://localhost:8181/user/filtered_items

`{
"min_price": 10.0,
"max_price": 200.0,
"min_rating": 2,
"max_rating": 4
}`
* <h4> Search Item by name</h4>

GET http://localhost:8181/user/search?name=item_name
* <h4> Purchase Item</h4>
POST http://localhost:8181/user/order?user_id=xx&item_id=xx
* <h4> Rate Item</h4>
PUT http://localhost:8181/user/items?item_id=1&user_id=2

`{
"Rating": 3
}`
* <h4> Comment Item</h4>
PUT http://localhost:8181/user/item/comment?item_id=1&user_id=1

`{
"comment": "comment 1"
}`
