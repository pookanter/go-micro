FROM cosmtrek/air

WORKDIR /app

COPY modules modules

RUN cd modules/db && go mod tidy