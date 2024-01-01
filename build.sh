docker build --tag aevanra/mictakepic:latest \
    --build-arg MONGO_URI=$MONGO_URI \
    --build-arg SMB_SHARE_HOST=$SMB_SHARE_HOST \
    --build-arg SMB_USER=$SMB_USER \
    --build-arg SMB_PASS=$SMB_PASS \
    --build-arg DB_NAME=$DB_NAME \
    --build-arg USERS_COLLECTION=$USERS_COLLECTION \
    --build-arg SESSION_SECRET=$SESSION_SECRET \
    . && \
docker push aevanra/mictakepic:latest

