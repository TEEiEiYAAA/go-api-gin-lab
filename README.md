# Teerat Srisukho 6609650442
Subject: CS367 Web Service Development Concepts

---

# Student API with Gin
REST API สำหรับการจัดการกับข้อมูลของนักศึกษา ซึ่งถูกพัฒนาด้วยภาษา Go, Gin และ SQLite

## ความสามารถระบบ
- ดึงข้อมูลนักศึกษาทั้งหมด (GET)
- ดึงข้อมูลนักศึกษาตาม Id (GET)
- เพิ่มนักศึกษาใหม่ (POST)
- แก้ไขข้อมูลนักศึกษา (PUT)
- ลบนักศึกษา (DELETE)
และมีการตรวจสอบความถูกต้องของข้อมูล (Validation)

## Project Structure
```
go-api-gin-lab/
├── config/
│   └── database.go
├── handlers/
│   └── student_handler.go
├── models/
│   └── student.go
├── repositories/
│   └── student_repository.go
├── services/
│   └── student_service.go
├── main.go
├── go.mod
├── go.sum
└── students.db
```

## Prerequisites
- Go version 1.21 ขึ้นไป
- Git

## How to Run

1. Clone repository
```bash
git clone https://github.com/TEEiEiYAAA/go-api-gin-lab.git
cd go-api-gin-lab
```
2. ติดตั้ง dependencies
```bash
go mod download
```
3. Run server
```bash
go run main.go
```
ซึ่งหลังจากรันจะมีข้อความขึ้นว่า:
```
[GIN-debug] Listening and serving HTTP on :8080
```
และ server ของเราจะทำงานที่
```
http://localhost:8080
```

## Testing API Endpoint
- เราจะใช้ Postman ในการทดสอบ
- Base URL: http://localhost:8080
- ตั้งค่า Headers โดยมี Key: Content-Type และ Value: application/json

### 1. ดึงข้อมูลนักศึกษาทั้งหมด
**GET** `/students`

#### Response: 200 OK

```json
[
  {
    "id": "650001",
    "name": "John Doe",
    "major": "Computer Science",
    "gpa": 3.5
  }
]
```

#### URL

```bash
http://localhost:8080/students
```

### 2. ดึงนักศึกษาตาม ID

**GET** `/students/:id`

#### สำเร็จ: 200 OK

```json
{
  "id": "650001",
  "name": "John Doe",
  "major": "Computer Science",
  "gpa": 3.5
}
```

#### ไม่พบข้อมูล: 404

```json
{
  "error": "Student not found"
}
```

#### URL

```bash
http://localhost:8080/students/650001
```
### 3. เพิ่มนักศึกษา

**POST** `/students

#### สำเร็จ: 201 Created

```json
{
  "id": "650001",
  "name": "John Doe",
  "major": "Computer Science",
  "gpa": 3.5
}
```
แต่ก็จะมีกรณีอื่น ๆ ที่จะเกิด Validation Error
- ไม่ใส่รหัสนักศึกษา(เว้นว่าง)
```
{
    "error": "id must not be empty"
}
```
- ไม่ใส่ชื่อ
```
{
    "error": "name must not be empty"
}
```
- GPA ไม่อยู่ในช่วง 0-4
```
{
    "error": "gpa must be between 0.00 and 4.00"
}
```

### 4. แก้ไขข้อมูลนักศึกษา

**PUT** `/students/:id

#### สำเร็จ: 200 OK
และระบบจะคืนข้อมูลที่ update แล้ว แต่ก็มี Error ได้ เช่น
- 400 ข้อมูลไม่ถูกต้อง
```
{
    "error": "Invalid JSON format"
}
```
- 404 ไม่พบนักศึกษา
```
{
    "error": "Student not found"
}
```

### 5. ลบข้อมูลนักศึกษา

**DELETE** `/students/:id

#### สำเร็จ: 204 No Content
แต่ก็มี Error ได้ เช่น
- 404 ไม่พบนักศึกษา
```
{
    "error": "Student not found"
}
```
