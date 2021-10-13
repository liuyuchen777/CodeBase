#!/usr/bin/python3

import pymysql

# 打开数据库连接
db = pymysql.connect("localhost","lyc64","lyc7758321321","crashcourse" )

# 使用 cursor() 方法创建一个游标对象 cursor
cursor = db.cursor()

# 使用 execute()  方法执行 SQL 查询 
cursor.execute("SELECT VERSION()")

# 使用 fetchone() 方法获取单条数据.
data = cursor.fetchone()

print ("Database version : %s " % data)

# 查询数据库
sql = "SELECT * from products"

try:
    cursor.execute(sql)
    results = cursor.fetchall()
    print(type(results))
    print(results)
    """
    for row in results:
        prod_id = row[0]
        vend_id = row[1]
        prod_name = row[2]
        prod_price = row[3]
        prod_desc = row[4]
        print("prod_id = %s, vend_id = %d, prod_name = %s, prod_price = %d, prod_desc = %s" %\
                (prod_id, vend_id, prod_name, prod_price, prod_desc))
    """
except:
    print("[Error] unable to fetch data")
    
# 关闭数据库连接
db.close()