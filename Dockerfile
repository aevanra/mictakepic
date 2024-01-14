FROM golang:1.21.5

ARG MONGO_URI
ARG DB_NAME
ARG USERS_COLLECTION
ARG SESSION_SECRET
ARG HOME_IMAGE_DIR_NAME

# Create and set working directory
RUN mkdir /app
WORKDIR /app

# Set environment variables
RUN echo "MONGO_URI=\"${MONGO_URI}\"" > /app/.env
RUN echo "DB_NAME=\"${DB_NAME}\"" >> /app/.env
RUN echo "USERS_COLLECTION=\"${USERS_COLLECTION}\"" >> /app/.env
RUN echo "SESSION_SECRET=\"${SESSION_SECRET}\"" >> /app/.env
RUN echo "HOME_IMAGE_DIR_NAME=\"${HOME_IMAGE_DIR_NAME}\"" >> /app/.env


# Copy app files
COPY . /app

RUN go mod download
RUN go build main.go    

EXPOSE 8082
CMD ["./main"]
