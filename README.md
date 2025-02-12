
# **Golang MySQL Boilerplate**  
A modular and scalable boilerplate for building web applications using **Golang**, **Gin**, and **MySQL**.

---

## **Features**  
- **Gin Framework** – High-performance HTTP web framework  
- **MySQL Integration** – Seamless database operations  
- **Modular Structure** – Scalable and easy-to-maintain codebase  
- **Auto CRUD Generation** – Quickly generate CRUD modules  
- **Hot Reloading** – Develop efficiently with `air`  

---

## **Installation**  
Clone the repository and navigate to the project folder:  
```bash
git clone git@github.com:MuhammadRifat/golang-mysql-boilerplate.git
cd golang-mysql-boilerplate
```

Install dependencies:  
```bash
go mod tidy
```

Copy environment variables:  
```bash
cp .env.example .env
```
Update `.env` with your database credentials.

---

## **Running the App**  
Start the application with **Air (hot reloading)**:  
```bash
air
```
Or run manually:  
```bash
go run ./src/main.go
```

---

## **Generate CRUD Module**  
Easily generate a new module with:  
```bash
chmod +x generate_module.sh
./generate_module.sh <module-name>
```
Replace `<module-name>` with your desired module (e.g., `user`, `product`).

---

## 📂 **Project Structure**  
```
golang-mysql-boilerplate/
│── src/
│   ├── config/            # Configuration files
│   ├── modules/           # Feature-based modules
│   │   ├── auth/          # Authentication module
│   │   ├── user/          # User module
│   ├── routes/            # API route definitions
│   ├── seed/              # Database seeding logic
│   ├── util/              # Utility/helper functions
│── main.go                # Application entry point
│── .air.toml              # Air (hot reloading) configuration
│── .env                   # Environment variables (ignored in Git)
│── .env.example           # Example environment configuration
│── .gitignore             # Git ignore file
│── generate_module.go     # Go script to generate modules
│── generate_module.sh     # Shell script to automate module creation
│── go.mod                 # Go modules dependencies
│── go.sum                 # Dependency checksums
│── README.md              # Project documentation
```

---

## **Contributing**  
Contributions are welcome! Feel free to submit issues and pull requests.

---

## **License**  
This project is licensed under the **MIT License**.

---
