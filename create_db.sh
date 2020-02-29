#docker-machine create --driver=virtualbox --virtualbox-memory 4096 php-piscine
docker-machine ssh php-piscine sudo mkdir -p /etc/mysql/conf.d
docker-machine ssh php-piscine sudo mkdir -p /var/lib/mysql
docker-machine ssh php-piscine sudo mount -t vboxsf -o uid=999,gid=50 conf.d /etc/mysql/conf.d
docker-machine ssh php-piscine sudo mount -t vboxsf -o uid=999,gid=50 mysql_base /var/lib/mysql
eval $(docker-machine env php-piscine)
docker run --detach --name=mysql-container --env="MYSQL_ROOT_PASSWORD=password" --publish 3306:3306 --volume=/var/lib/mysql:/var/lib/mysql --volume=/etc/mysql/conf.d:/etc/mysql/conf.d  mysql
#mysql -uroot -ppassword -h $(docker-machine ip php-piscine)
#mysql -uroot -ppassword -h $(docker-machine ip php-piscine) db_tmatthew < base_student.sql > output.log
