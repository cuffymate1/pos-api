# POS API

ระบบ POS API พัฒนาโดยใช้ภาษา Go และ Framework Fiber พร้อมด้วย GORM สำหรับการจัดการฐานข้อมูล PostgreSQL และ JWT สำหรับการยืนยันตัวตน

## 🧰 Stack ที่ใช้

- [Go](https://golang.org/)
- [Fiber](https://github.com/gofiber/fiber) – Web Framework
- [GORM](https://gorm.io/) – ORM สำหรับ Go
- [PostgreSQL](https://www.postgresql.org/) – ระบบฐานข้อมูล
- [JWT](https://github.com/golang-jwt/jwt) – JSON Web Token สำหรับ Auth
- [godotenv](https://github.com/joho/godotenv) – โหลดค่าจากไฟล์ `.env`

---

## 🚀 เริ่มต้นใช้งาน

### 1. Clone โปรเจกต์

```bash
git clone https://github.com/cuffymate1/pos-api.git
cd pos-api

go mod init github.com/cuffymate1/pos-api
go get github.com/gofiber/fiber/v2
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/golang-jwt/jwt/v5
go get github.com/joho/godotenv

go run .

```

### สร้างไฟล์ .env --> config/.env
- DB_HOST=localhost
- DB_PORT=5432
- DB_USER=postgres
- DB_PASSWORD=yourpassword
- DB_NAME=yourdbname
- JWT_SECRET=your_jwt_secret