MYSQL="mysql -h$MYSQL_HOST -P$MYSQL_PORT -u$MYSQL_USER -p$MYSQL_PASSWORD --default_character_set utf8"

files=`ls build/support-files/sql/*.sql`
for f in $files
do
    echo "source the sql file: $f"
    $MYSQL bkiam < $f
done