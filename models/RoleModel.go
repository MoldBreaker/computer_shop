package models

/*CREATE TABLE Role (
    role_id INT NOT NULL AUTO_INCREMENT,
    role_name VARCHAR(50) NULL,
    PRIMARY KEY (role_id)
);*/

type RoleModel struct {
	RoleId int
	RoleName string
}
