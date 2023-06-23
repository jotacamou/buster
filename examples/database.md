## Examples on how to use the `database` library

Creating a new database pool:
```
import "github.com/jotacamou/buster/database"

func main() {
    dbh, err := databae.NewDatabaseHandler("dev")
    if err != nil {
        // Handle error
    }
```

Get all transactions:
```
// GetAllTransactions() returns a slice of transactions
transactions, err := dbh.GetAllTransactions()
if err != nil {
    // Handle error
}
```

Update an existing transaction:
```
trx, err = dbh.GetTransactionByReferenceId(referenceId)
if err != nil {
    // Handle error
}

// update the transaction values
trx.Status = database.Canceled
trx.Amount = 200

// commit transaction update
dbh.UpdateTransaction(trx)
```

Updating an existing transaction status by Reference ID:
```
// Prepare the transaction update
st := database.PendingStatus
trxUpdate := databaseTransactionUpdate(Status: &st)

// Commit the transaction
err = dbh.UpdateTransactionByReferenceId(referenceId, trxUpdate)
if err != nil {
    // Handle error
}
```
