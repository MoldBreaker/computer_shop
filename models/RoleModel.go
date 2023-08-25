package models

import "database/sql"

/*CREATE TABLE Role (
    role_id INT NOT NULL AUTO_INCREMENT,
    role_name VARCHAR(50) NULL,
    PRIMARY KEY (role_id)
);*/

type RoleModel struct {
	RoleId int
	RoleName string
}

func(RM * RoleModel) GetTableName() string {
	return `Role`
}

func(a *RoleModel) ScanToRoleModel(rows *sql.Rows) (*RoleModel, error) {
	err := rows.Scan(&a.RoleId, &a.RoleName)
	if err != nil {
		return nil, err
	}
	return a, nil
}