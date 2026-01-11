## Deploy via docker compose

```yml
services:
  backend:
    image: ghcr.io/y3eet/click-in/backend:latest
    environment:
      IS_PROD: "true"
      PORT: "8080"
      DATABASE_URL: ${DATABASE_URL}
      SECRET_KEY: ${SECRET_KEY}
      GOOGLE_CLIENT_ID: ${GOOGLE_CLIENT_ID}
      GOOGLE_CLIENT_SECRET: ${GOOGLE_CLIENT_SECRET}
      MINIO_ENDPOINT: ${MINIO_ENDPOINT}
      MINIO_ROOT_USER: ${MINIO_ROOT_USER}
      MINIO_ROOT_PASSWORD: ${MINIO_ROOT_PASSWORD}
      JWT_ACCESS_SECRET: ${JWT_ACCESS_SECRET}
      JWT_EXCHANGE_SECRET: ${JWT_EXCHANGE_SECRET}
      JWT_REFRESH_SECRET: ${JWT_REFRESH_SECRET}

      FRONTEND_URL: ${FRONTEND_URL:-http://localhost:3000}
      WEB_URL: ${WEB_URL:-http://localhost:3000}
      BASE_URL: ${BASE_URL:-http://localhost:8080}
      JWT_SECRET: ${JWT_SECRET:-}
    ports:
      - "8080:8080"
    restart: unless-stopped

  web:
    build:
      context: ./web
      dockerfile: Dockerfile
    environment:
      NEXT_PUBLIC_API_URL: ${NEXT_PUBLIC_API_URL}
      PORT: "3000"
    ports:
      - "3000:3000"
    depends_on:
      - backend
    restart: unless-stopped
```
