# bank

## Setup

### make
run → 
```powershell
sudo apt install make
```

### sqlc
run → 
```powershell
sudo apt install sqlc
```

### Docker
- Install docker desktop
- navigate to download folder
- run → 
```powershell
Start-Process "Docker Desktop Installer.exe" -Verb RunAs -Wait -ArgumentList "install --installation-dir=C:\Docker\"
```

### PostgreSQL Image
To manage the PostgreSQL image using Docker, follow these steps:

1. **Pull the Docker image**
   - Command:
     ```powershell
     docker pull postgres
     ```
   You can copy this command and paste it into your terminal.

2. **Create a Docker container**
   - Command:
     ```powershell
     docker run --name go-bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres
     ```
   You can copy this command and paste it into your terminal.

3. **Connect to the database**
   - Command:
     ```powershell
     docker exec -it go-bank psql -U root
     ```
   You can copy this command and paste it into your terminal.

### DATABASE MIGRATION
```powershell
migrate create -ext sql -dir db/migration -seq init_schema
```

```powershell
migrate -path db/migration -database "postgresql://root:root@localhost:5432/bank?sslmode=disable" -verbose up
```