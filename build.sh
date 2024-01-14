docker build --tag aevanra/mictakepic:latest \
    --build-arg MONGO_URI=$MONGO_URI \
    --build-arg DB_NAME=$DB_NAME \
    --build-arg USERS_COLLECTION=$USERS_COLLECTION \
    --build-arg SESSION_SECRET=$SESSION_SECRET \
    --build-arg HOME_IMAGE_DIR_NAME=$HOME_IMAGE_DIR_NAME \
    . && \
docker push aevanra/mictakepic:latest

