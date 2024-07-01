
FROM mysql:oraclelinux9

# MySQLをrootで実行する設定に書き換える
RUN sed -e 's/^ *user *= *mysql$/user = root/' -i /db/my.cnf