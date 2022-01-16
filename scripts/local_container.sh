docker container create --name inedu_kampusmerdeka_container -p 8080:8080 inedu_kampusmerdeka:1.0.0
docker container start inedu_kampusmerdeka

docker run --name mysql_inedu_kampusmerdeka -e MYSQL_ROOT_PASSWORD=haji12345 -d mysql:8.0