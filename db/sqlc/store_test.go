package db

import "context"

func TestTransferTx(tx *testing.T){
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// run a concurrent transfer transaction
	n := 5
	amount := int64(10)


	// check errors after go routine returns back to main routine after its execution:
	// check errors using channels:
	errs := make(chan error)

	for i := 0; i< n; i++ {
		// this go routine is running separate 
		// from the main function TestTransferTx it is being run in:
		go func ()  {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID: account2.ID,
				Amount: amount,
			})

			// need brackets at the end for it to run the goroutine!!!:
		}()
	}
}