package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

const (
	StatusPending   = "PENDING"
	StatusCanceled  = "CANCELED"
	StatusCompleted = "COMPLETED"
	StatusCreated   = "CREATED"
)

// DatabaseHandler contains the database connection pool reference.
type DatabaseHandler struct {
	DB *sql.DB
}

// TransactionUpdate contains references to the values to be updated by
// UpdateTransactionByReferenceId().  Only Status and Amount can change.
type TransactionUpdate struct {
	Status *string
	Amount *float32
}

// NewDatabaseHandler returns a new DatabaseHandler based on the
// environment database key on the configuration file.
func NewDatabaseHandler(env string) (*DatabaseHandler, error) {
	dbconf, err := loadConfig(env)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(dbconf.Driver, generateDSN(dbconf))
	if err != nil {
		return nil, err
	}

	dbh := &DatabaseHandler{
		DB: db,
	}

	return dbh, nil
}

// loadConfig parses the database configuration for the given environment key.
func loadConfig(env string) (*DatabaseConfig, error) {
	var config interface{}

	configFile, err := ioutil.ReadFile("database.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}

	jsonConf, err := json.Marshal(config.(map[string]interface{})[env])
	if err != nil {
		return nil, err
	}

	dbconf := &DatabaseConfig{}
	json.Unmarshal(jsonConf, &dbconf)

	return dbconf, nil
}

// CreateTransaction inserts a new row to the transaction table.
func (dbh *DatabaseHandler) CreateTransaction(referenceId, externalId, status string, amount float32) error {
	stmt, err := dbh.DB.Prepare("INSERT INTO transaction(reference_id, external_id, status, amount) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(referenceId, externalId, status, amount)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTransaction updates an existing transaction.
func (dbh *DatabaseHandler) UpdateTransaction(trx *Transaction) error {
	stmt, err := dbh.DB.Prepare("UPDATE transaction set status = ?, amount = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(trx.Status, trx.Amount, trx.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTransactionByReferenceId updates a transaction by passing new values from a TransactionUpdate.
func (dbh *DatabaseHandler) UpdateTransactionByReferenceId(referenceId string, update TransactionUpdate) error {
	trx, err := dbh.GetTransactionByReferenceId(referenceId)
	if err != nil {
		return err
	}

	if update.Status != nil {
		trx.Status = *update.Status
	}

	if update.Amount != nil {
		trx.Amount = *update.Amount
	}

	if err = dbh.UpdateTransaction(trx); err != nil {
		return err
	}

	return nil
}

// GetTransaction retrieves a transaction by ID filter.
func (dbh *DatabaseHandler) GetTransaction(id int) (*Transaction, error) {
	trx := &Transaction{}
	rows := dbh.DB.QueryRow("SELECT * from transaction WHERE id = ?", id)
	err := rows.Scan(&trx.ID, &trx.ReferenceID, &trx.ExternalID, &trx.Status, &trx.Amount, &trx.Created)
	if err != nil {
		return nil, err
	}

	return trx, nil
}

// GetTransactionByReferenceId retrieves a transaction by Reference ID.
func (dbh *DatabaseHandler) GetTransactionByReferenceId(referenceId string) (*Transaction, error) {
	trx := &Transaction{}
	rows := dbh.DB.QueryRow("SELECT * from transaction WHERE reference_id = ?", referenceId)
	err := rows.Scan(&trx.ID, &trx.ReferenceID, &trx.ExternalID, &trx.Status, &trx.Amount, &trx.Created)
	if err != nil {
		return nil, err
	}

	return trx, nil
}

// GetAllTransactions returns a slice with all of the transactions in the table.
func (dbh *DatabaseHandler) GetAllTransactions() ([]Transaction, error) {
	rows, err := dbh.DB.Query("SELECT * from transaction")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []Transaction

	for rows.Next() {
		trx := Transaction{}
		_ = rows.Scan(&trx.ID, &trx.ReferenceID, &trx.ExternalID, &trx.Status, &trx.Amount, &trx.Created)
		transactions = append(transactions, trx)
	}

	err = rows.Err()

	return transactions, err
}

// generateDSN returns a DSN string to pass to the mysql db driver based on configuration.
func generateDSN(conf *DatabaseConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
}
