FROM cosmtrek/air AS development

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["air", "-c", ".air.toml"]