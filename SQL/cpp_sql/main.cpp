/*************************************************************************
    > File Name: main.cpp
    > Author: SongLee
    > E-mail: lisong.shine@qq.com
    > Created Time: 2014年05月05日 星期一 00时30分45秒
    > Personal Blog: http://songlee24.github.io
 ************************************************************************/
#include<iostream>
#include "MyDB.h"
using namespace std;

// macro
#define FUN             RETRIEVE

#define INQUIRE         0
#define GET_VERSION     1
#define CREATE_DB       2
#define CRE_AND_POP_TAB 3
#define RETRIEVE        4

void finish_with_error(MYSQL *con)
{
    fprintf(stderr, "%s\n", mysql_error(con));
    mysql_close(con);
    exit(1);
}

int main()
{
#if FUN == INQUIRE
	MyDB db;
	db.initDB("127.0.0.1", "lyc64", "lyc7758321321", "crashcourse");
	db.exeSQL("select * from products");
#elif FUN == GET_VERSION
    printf("MySQL client version: %s\n", mysql_get_client_info());
#elif FUN == CREATE_DB
    MYSQL *con = mysql_init(NULL);

    if (con == NULL)
    {
        fprintf(stderr, "%s\n", mysql_error(con));
        exit(1);
    }

    if (mysql_real_connect(con, "127.0.01", "lyc64", "lyc7758321321",
            NULL, 0, NULL, 0) == NULL)
    {
        fprintf(stderr, "%s\n", mysql_error(con));
        mysql_close(con);
        exit(1);
    }

    if (mysql_query(con, "DROP DATABASE IF EXISTS testdb"))
    {
        fprintf(stderr, "%s\n", mysql_error(con));
        mysql_close(con);
        exit(1);
    }

    if (mysql_query(con, "CREATE DATABASE testdb"))
    {
        fprintf(stderr, "%s\n", mysql_error(con));
        mysql_close(con);
        exit(1);
    }

    mysql_close(con);
#elif FUN == CRE_AND_POP_TAB
    MYSQL *con = mysql_init(NULL);

    if (con == NULL)
    {
        fprintf(stderr, "%s\n", mysql_error(con));
        exit(1);
    }

    if (mysql_real_connect(con, "127.0.0.1", "lyc64", "lyc7758321321",
            "testdb", 0, NULL, 0) == NULL)
    {
        finish_with_error(con);
    }

    if (mysql_query(con, "DROP TABLE IF EXISTS cars")) {
        finish_with_error(con);
    }

    if (mysql_query(con, "CREATE TABLE cars(id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(255), price INT)")) {
        finish_with_error(con);
    }

    if (mysql_query(con, "INSERT INTO cars VALUES(1,'Audi',52642)")) {
        finish_with_error(con);
    }

    if (mysql_query(con, "INSERT INTO cars VALUES(2,'Mercedes',57127)")) {
        finish_with_error(con);
    }

    if (mysql_query(con, "INSERT INTO cars VALUES(3,'Skoda',9000)")) {
        finish_with_error(con);
    }

    if (mysql_query(con, "INSERT INTO cars VALUES(4,'Volvo',29000)")) {
        finish_with_error(con);
    }

    if (mysql_query(con, "INSERT INTO cars VALUES(5,'Bentley',350000)")) {
        finish_with_error(con);
    }

    if (mysql_query(con, "INSERT INTO cars VALUES(6,'Citroen',21000)")) {
        finish_with_error(con);
    }

    if (mysql_query(con, "INSERT INTO cars VALUES(7,'Hummer',41400)")) {
        finish_with_error(con);
    }

    if (mysql_query(con, "INSERT INTO cars VALUES(8,'Volkswagen',21600)")) {
        finish_with_error(con);
    }

    mysql_close(con);
#elif FUN == RETRIEVE
    MYSQL *con = mysql_init(NULL);

    if (con == NULL)
    {
        fprintf(stderr, "mysql_init() failed\n");
        exit(1);
    }

    if (mysql_real_connect(con, "127.0.0.1", "lyc64", "lyc7758321321",
            "testdb", 0, NULL, 0) == NULL)
    {
        finish_with_error(con);
    }

    if (mysql_query(con, "SELECT * FROM cars"))
    {
        finish_with_error(con);
    }

    MYSQL_RES *result = mysql_store_result(con);

    if (result == NULL)
    {
        finish_with_error(con);
    }

    int num_fields = mysql_num_fields(result);

    MYSQL_ROW row;

    while ((row = mysql_fetch_row(result)))
    {
        for(int i = 0; i < num_fields; i++)
        {
            printf("%s ", row[i] ? row[i] : "NULL");
        }

        printf("\n");
    }

    mysql_free_result(result);
    mysql_close(con);

#endif
	return 0;
}