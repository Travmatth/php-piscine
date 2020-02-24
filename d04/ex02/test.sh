PW_FILE=${PWD}/../private/passwd
rm $PW_FILE
curl -d login=toto1 -d passwd=titi1 -d submit=OK "http://${HOSTNAME:-$(hostname)}:8080/j04/ex01/create.php"
#OK
more ~/http/Piscines/j04/htdocs/private/passwd
# a:1:
# {
# 	i:0;
# 	a:2:{
# 		s:5:"login";
# 		s:5:"toto1";
# 		s:6:"passwd";
# 		s:128:"2bdd45b3c828273786937ac1b4ca7908a431019e8b93c9fd337317f92fac80dace29802bedc33d9259c8b55d1572cb8a6c1df8579cdaa02256099ed52a905d38";
# 	}
# }
curl -d login=toto1 -d passwd=titi1 -d submit=OK "http://${HOSTNAME:-$(hostname)}:8080/j04/ex01/create.php"
# ERROR
curl -d login=toto2 -d passwd= -d submit=OK "http://${HOSTNAME:-$(hostname)}:8080/j04/ex01/create.php"
# ERROR$