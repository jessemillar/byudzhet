#!/usr/bin/env bash

echo "Stopping running application"
ssh $DEPLOY_USER@$DEPLOY_HOST 'docker stop byudzhet'
ssh $DEPLOY_USER@$DEPLOY_HOST 'docker rm byudzhet'

echo "Pulling latest version"
ssh $DEPLOY_USER@$DEPLOY_HOST 'docker pull jessemillar/byudzhet:latest'

echo "Starting the new version"
ssh $DEPLOY_USER@$DEPLOY_HOST 'docker run -d --restart=always --name byudzhet -p 8000:8000 jessemillar/byudzhet:latest'

echo "Success!"

exit 0