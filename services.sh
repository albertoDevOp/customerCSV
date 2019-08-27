#/bin/bash
export PG_USER=root
export PG_PWD=root
export PG_DB=paack
case $1 in
  up)
    docker-compose -f ./docker-compose.yml up -d
    ;;
  restart)
    docker-compose restart
    ;;
  build)
    docker build -t csv_image -f images/Dockerfile images
    ;;
  *)
    echo "option not avaliable"
    ;;
esac