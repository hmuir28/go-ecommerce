“E-commerce API Documentation ----------------------------”
“This API provides functionality for user management, product management, and shopping cart operations. Below are the available endpoints and instructions on how to run the app locally.”

### **API Endpoints**”

#### **User Authentication & Management**

* **POST** `/users/signup`
Create a new user account.
**Handler:** `controllers.SignUp()`
* **POST** `/users/login`
Log in an existing user.
**Handler:** `controllers.Login()` 

#### **Admin Product Management**

* **POST** `/admin/addproduct`
Add a new product (admin access only).
**Handler:** `controllers.ProductViewerAdmin()`

#### **Product Management**
* **POST** `/products`
Create a new product.
**Handler:** `controllers.CreateProduct()`
* **GET** `/products`
Retrieve all available products.
**Handler:** `controllers.FindProducts()`
* **GET** `/products/:id`
Retrieve a specific product by its ID.
**Handler:** `controllers.FindProductById()`
* **GET** `/products/search`
Search for products based on query parameters.
**Handler:** `controllers.FindProductByQuery()`

#### **Shopping Cart Operations**
* **GET** `/addtocart`
Add a product to the user's cart.
**Handler:** `app.AddProductToCart()`
* **GET** `/removeitem`
Remove an item from the user's cart.
**Handler:** `app.RemoveItem()`
* **GET** `/cartcheckout`
Checkout all items in the cart and process the order.
**Handler:** `app.BuyItemFromCart()`
* **GET** `/instantbuy`
Buy a product instantly without adding it to the cart.
**Handler:** `app.InstantBuyer()`

### **Running the App Locally**
To run the application locally, ensure you have **Make** installed. Use the following command: bash `make run` This will start the server, and the app will be available at `http://localhost:8080`.

### **Dependencies**
Make sure you have the following installed:
* **Go** (latest version recommended)
* **MongoDB** (for data storage)
* **Gin Web Framework** (`go get github.com/gin-gonic/gin`)

### **Contact**
For any questions or issues, feel free to open an issue or contact the development team.