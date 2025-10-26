1. jalankan cd api_golang_tj di terminal/cmd, lalu ketikan docker build -t image-repo/builder/service_tj .
2. di path service_transjakarta, ketikan di dalam terminal/cmd docker compose -f docker-compose-api-golang-tj.yaml -p service_tranjakarta up -d
3. import DB postgres dengan file create_table.sql
