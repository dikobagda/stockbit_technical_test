/* Create Table User */
CREATE TABLE stockbit_user
( ID INT(11) NOT NULL AUTO_INCREMENT,
  UserName VARCHAR(30) NOT NULL,
  Parent INT(11),
  CONSTRAINT user_pk PRIMARY KEY (ID)
);

/* Insert Table User */
INSERT INTO `stockbit_user` (`ID`, `UserName`, `Parent`) VALUES (1, 'ALI', '2');
INSERT INTO `stockbit_user` (`ID`, `UserName`, `Parent`) VALUES (2, 'Budi', '0');
INSERT INTO `stockbit_user` (`ID`, `UserName`, `Parent`) VALUES (3, 'Cecep', '1')

/* To do */
/*
Terdapat sebuah table "USER" yg memiliki 3 kolom: ID, UserName, Parent. Di mana:
Kolom ID adalah Primary Key
Kolom UserName adalah Nama User
Kolom Parent adalah ID dari User yang menjadi Creator untuk User tertentu.
eg.
——————————————————————————
| ID | UserName | Parent |
——————————————————————————
| 1  | Ali      |   2    |
| 2  | Budi     |   0    |
| 3  | Cecep    |   1    |
—————————————————————————-
Tuliskan SQL Query untuk mendapatkan data berisi:
——————————————————————————————————
| ID | UserName | ParentUserName |
——————————————————————————————————
| 1  | Ali      |     Budi       |
| 2  | Budi     |     NULL       |
| 3  | Cecep    |     Ali        |
——————————————————————————————————
*Kolom ParentUserName adalah UserName berdasarkan value Parent
*/

/* Answer */
SELECT stu.ID, stu.UserName, sta.UserName as ParentUserName
FROM stockbit_user stu
LEFT JOIN stockbit_user sta ON (sta.ID = stu.Parent)
ORDER by stu.ID