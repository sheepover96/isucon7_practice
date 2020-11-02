import MySQLdb

conn = MySQLdb.connect(
    user='isucon',
    password='isucon',
    host='localhost',
    db='isubata'
)

cur = conn.cursor()

sql = 'select id, name, data from image'

res = cur.execute(sql)

conn.close()
