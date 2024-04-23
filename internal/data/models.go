package data

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type DBModel struct {
	DB *sql.DB
}
type Models struct {
	Users  UserModel
	Tokens TokenModel
}

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

func NewModels(db *sql.DB) Models {
	return Models{
		Users:  UserModel{DB: db}, // Initialize a new UserModel instance.
		Tokens: TokenModel{DB: db},
	}
}

func (m *DBModel) Insert(moduleInfo *ModuleInfo) error {
	_, err := m.DB.Exec("INSERT INTO module_info (created_at, updated_at, module_name, module_duration, exam_type, version) VALUES ($1, $2, $3, $4, $5, $6)",
		moduleInfo.CreatedAt, moduleInfo.UpdatedAt, moduleInfo.ModuleName, moduleInfo.ModuleDuration, moduleInfo.ExamType, moduleInfo.Version)
	if err != nil {
		return fmt.Errorf("failed to insert module info: %w", err)
	}
	return nil
}

func (m *DBModel) Retrieve(id int) (*ModuleInfo, error) {
	var moduleInfo ModuleInfo
	row := m.DB.QueryRow("SELECT * FROM module_info WHERE id = $1", id)
	err := row.Scan(&moduleInfo.ID, &moduleInfo.CreatedAt, &moduleInfo.UpdatedAt, &moduleInfo.ModuleName, &moduleInfo.ModuleDuration, &moduleInfo.ExamType, &moduleInfo.Version)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve module info: %w", err)
	}
	return &moduleInfo, nil
}

func (m *DBModel) Update(moduleInfo *ModuleInfo) error {
	_, err := m.DB.Exec("UPDATE module_info SET created_at = $1, updated_at = $2, module_name = $3, module_duration = $4, exam_type = $5, version = $6 WHERE id = $7",
		moduleInfo.CreatedAt, moduleInfo.UpdatedAt, moduleInfo.ModuleName, moduleInfo.ModuleDuration, moduleInfo.ExamType, moduleInfo.Version, moduleInfo.ID)
	if err != nil {
		return fmt.Errorf("failed to update module info: %w", err)
	}
	return nil
}

func (m *DBModel) Delete(id int) error {
	_, err := m.DB.Exec("DELETE FROM module_info WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete module info: %w", err)
	}
	return nil
}

// Defence

func (m *DBModel) InsertDepartmentInfo(depInfo *DepartmentInfo) error {
	err := m.DB.QueryRow("INSERT INTO department_info (department_name, staff_quantity, department_director, module_id) VALUES ($1, $2, $3, $4) RETURNING id",
		depInfo.DepartmentName, depInfo.StaffQuantity, depInfo.DepartmentDirector, depInfo.ModuleID).Scan(&depInfo.ID)
	if err != nil {
		return fmt.Errorf("failed to insert department info: %w", err)
	}
	return nil
}

func (m *DBModel) RetrieveDepartmentInfo(id int) (*DepartmentInfo, error) {
	var departmentInfo DepartmentInfo
	row := m.DB.QueryRow("SELECT id, department_name, staff_quantity, department_director, module_id FROM department_info WHERE id = $1", id)
	err := row.Scan(&departmentInfo.ID, &departmentInfo.DepartmentName, &departmentInfo.StaffQuantity, &departmentInfo.DepartmentDirector, &departmentInfo.ModuleID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve department info: %w", err)
	}
	return &departmentInfo, nil
}

// Assingment 2

// func (m *DBModel) InsertUserInfo(userInfo *UserInfo) error {
// 	query := `
//         INSERT INTO user_info (created_at, updated_at, name, surname, email, password_hash, role, activated, version)
//         VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
//         RETURNING id`

// 	row := m.DB.QueryRow(query, userInfo.CreatedAt, userInfo.UpdatedAt, userInfo.Name, userInfo.Surname, userInfo.Email, userInfo.PasswordHash, userInfo.Role, userInfo.Activated, userInfo.Version)

// 	err := row.Scan(&userInfo.ID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (m *DBModel) GetUserInfo(id int) (*UserInfo, error) {
// 	query := `
//         SELECT id, created_at, updated_at, name, surname, email, password_hash, role, activated, version
//         FROM user_info
//         WHERE id = $1`

// 	userInfo := &UserInfo{}

// 	err := m.DB.QueryRow(query, id).Scan(&userInfo.ID, &userInfo.CreatedAt, &userInfo.UpdatedAt, &userInfo.Name, &userInfo.Surname, &userInfo.Email, &userInfo.PasswordHash, &userInfo.Role, &userInfo.Activated, &userInfo.Version)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return userInfo, nil
// }

// func (m *DBModel) GetAllUserInfo() ([]*UserInfo, error) {
// 	query := `
//         SELECT id, created_at, updated_at, name, surname, email, password_hash, role, activated, version
//         FROM user_info`

// 	rows, err := m.DB.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var userInfos []*UserInfo

// 	for rows.Next() {
// 		userInfo := &UserInfo{}
// 		err := rows.Scan(&userInfo.ID, &userInfo.CreatedAt, &userInfo.UpdatedAt, &userInfo.Name, &userInfo.Surname, &userInfo.Email, &userInfo.PasswordHash, &userInfo.Role, &userInfo.Activated, &userInfo.Version)
// 		if err != nil {
// 			return nil, err
// 		}
// 		userInfos = append(userInfos, userInfo)
// 	}

// 	return userInfos, nil
// }

// func (m *DBModel) UpdateUserInfo(userInfo *UserInfo) error {
// 	query := `
//         UPDATE user_info
//         SET updated_at = $1, name = $2, surname = $3, email = $4, password_hash = $5, role = $6, activated = $7, version = $8
//         WHERE id = $9`

// 	_, err := m.DB.Exec(query, userInfo.UpdatedAt, userInfo.Name, userInfo.Surname, userInfo.Email, userInfo.PasswordHash, userInfo.Role, userInfo.Activated, userInfo.Version, userInfo.ID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (m *DBModel) DeleteUserInfo(id int) error {
// 	query := `
//         DELETE FROM user_info
//         WHERE id = $1`

// 	_, err := m.DB.Exec(query, id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
