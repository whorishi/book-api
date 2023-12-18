This a book rest api for a book store which stores the data of the books present in the bookstore and perform CRUD operations.

1. GET: we can get all the books available in the book store.
2. POST: Add new book to the database.
3. PUT: Update book details with given id.
4. DELETE: Delete a book with a given id, which is removed or sold from book store.


Sequence Diagram for this Book API:

![sequence Diagram](https://github.com/whorishi/book-api/assets/76156125/ec3963f6-efff-480b-89ef-5f3abcb77d4e)


To initiate docker use these commands:
> docker pull mysql

> docker run --name sample-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_database -p 3307:3306 -d mysql:8.0.30

> docker exec -it sample-mysql mysql -uroot -proot123 test_database -e "CREATE TABLE books (id INT AUTO_INCREMENT PRIMARY KEY, title VARCHAR(255) NOT NULL, author VARCHAR(255), publisher VARCHAR(255), price INT, category VARCHAR(255));"

To Check database use

> docker exec -it sample-mysql bash

> mysql -uroot -p

Now Enter Password i.e root123 and change database to test_database.

Now you can perform sql quesries on database 



We have used the default port i.e 8000

use 'http://localhost:8000/books' to access the GET and POST Query.

use 'http://localhost:8000/books/{id}' to access the PUT and DELETE Query. // {} surrounding id should not be used.



