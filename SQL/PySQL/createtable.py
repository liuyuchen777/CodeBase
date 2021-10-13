#!/usr/bin/python3

import pymysql

# 打开数据库连接
db = pymysql.connect("localhost","lyc64","lyc7758321321")

# 使用 cursor() 方法创建一个游标对象 cursor
mycursor = db.cursor()

mycursor.execute("SELECT VERSION()")

data = mycursor.fetchone()

print ("Database version : %s " % data)

mycursor.execute("DROP DATABASE IF EXISTS mydatabase")
# create database
mycursor.execute("CREATE DATABASE mydatabase")
# change database
mycursor.execute("USE mydatabase")
# 使用 execute()  方法执行 SQL 查询 
mycursor.execute("DROP TABLE IF EXISTS employee")
# create table 
sql = """CREATE TABLE EMPLOYEE (
        FIRST_NAME  CHAR(20) NOT NULL,
        LAST_NAME  CHAR(20),
        AGE INT,  
        SEX CHAR(1),
        INCOME FLOAT )"""

mycursor.execute(sql)

# sql insert
# SQL 插入语句
sql = """INSERT INTO EMPLOYEE(FIRST_NAME,
        LAST_NAME, AGE, SEX, INCOME)
        VALUES ('Mac', 'Mohan', 20, 'M', 2000)"""
try:
    # 执行sql语句
    mycursor.execute(sql)
    # 提交到数据库执行
    db.commit()
except:
    # Rollback in case there is any error
    db.rollback()

db.close()