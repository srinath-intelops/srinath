package daos

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/srinath-intelops/srinath/test-component/pkg/rest/server/daos/clients/sqls"
	"github.com/srinath-intelops/srinath/test-component/pkg/rest/server/models"
)

type TestresourceDao struct {
	sqlClient *sqls.MySQLClient
}

func migrateTestresources(r *sqls.MySQLClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS testresources(
		ID int NOT NULL AUTO_INCREMENT,
        
		Name INT NOT NULL,
	    PRIMARY KEY (ID)
	);
	`
	_, err := r.DB.Exec(query)
	return err
}

func NewTestresourceDao() (*TestresourceDao, error) {
	sqlClient, err := sqls.InitMySQLDB()
	if err != nil {
		return nil, err
	}
	err = migrateTestresources(sqlClient)
	if err != nil {
		return nil, err
	}
	return &TestresourceDao{
		sqlClient,
	}, nil
}

func (testresourceDao *TestresourceDao) CreateTestresource(m *models.Testresource) (*models.Testresource, error) {
	insertQuery := "INSERT INTO testresources(Name) values(?)"
	res, err := testresourceDao.sqlClient.DB.Exec(insertQuery, m.Name)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			if mysqlErr.Number == 1062 {
				return nil, sqls.ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id
	log.Debugf("testresource created")
	return m, nil
}

func (testresourceDao *TestresourceDao) UpdateTestresource(id int64, m *models.Testresource) (*models.Testresource, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	testresource, err := testresourceDao.GetTestresource(id)
	if err != nil {
		return nil, err
	}
	if testresource == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE testresources SET Name = ? WHERE Id = ?"
	res, err := testresourceDao.sqlClient.DB.Exec(updateQuery, m.Name, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("testresource updated")
	return m, nil
}

func (testresourceDao *TestresourceDao) DeleteTestresource(id int64) error {
	deleteQuery := "DELETE FROM testresources WHERE Id = ?"
	res, err := testresourceDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("testresource deleted")
	return nil
}

func (testresourceDao *TestresourceDao) ListTestresources() ([]*models.Testresource, error) {
	selectQuery := "SELECT * FROM testresources"
	rows, err := testresourceDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var testresources []*models.Testresource
	for rows.Next() {
		m := models.Testresource{}
		if err = rows.Scan(&m.Id, &m.Name); err != nil {
			return nil, err
		}
		testresources = append(testresources, &m)
	}
	if testresources == nil {
		testresources = []*models.Testresource{}
	}
	log.Debugf("testresource listed")
	return testresources, nil
}

func (testresourceDao *TestresourceDao) GetTestresource(id int64) (*models.Testresource, error) {
	selectQuery := "SELECT * FROM testresources WHERE Id = ?"
	row := testresourceDao.sqlClient.DB.QueryRow(selectQuery, id)

	m := models.Testresource{}
	if err := row.Scan(&m.Id, &m.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}
	log.Debugf("testresource retrieved")
	return &m, nil
}
