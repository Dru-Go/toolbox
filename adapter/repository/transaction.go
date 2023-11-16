package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/bokwoon95/sq"
	"github.com/dru-go/noah-toolbox/domain"
	"github.com/dru-go/noah-toolbox/utils"
	"github.com/google/uuid"
)

const (
	DateFormat = "2006-01-02 15:04:05"
)

type TRANSACTIONS struct {
	sq.TableStruct
	ID, MATERIALID, PROJECT, CATEGORY, COMPANY, TRANSACTIONTYPE, REFERENCE_NO sq.StringField
	RECEIVED, ISSUED, RETURN, BALANCE                                         sq.NumberField
	UNITPRICE, TOTALPRICE, CUMULATIVE                                         sq.StringField
	SOURCE, DESTINATION, REMARK, CREATEDAT, UPDATEDAT, DELETEDAT              sq.StringField
}

type ITransactionRepository interface {
	LastTransaction(id string) (domain.Transaction, error)
	Fetch(domain.ComputeFilter) ([]domain.Transaction, error)
	BulkCreate([]domain.Transaction) error
	BulkUpdate([]domain.Transaction) error
}

func (m TRANSACTIONS) TransactionMapper() func(row *sq.Row) domain.Transaction {
	return func(row *sq.Row) domain.Transaction {
		return domain.Transaction{
			Id:              row.StringField(m.ID),
			MaterialId:      row.StringField(m.MATERIALID),
			Project:         row.StringField(m.PROJECT),
			Company:         row.StringField(m.COMPANY),
			TransactionType: row.StringField(m.TRANSACTIONTYPE),
			ReferenceNo:     row.StringField(m.REFERENCE_NO),
			Units:           0,
			Quantity:        0,
			Received:        row.IntField(m.RECEIVED),
			Issued:          row.IntField(m.ISSUED),
			Returns:         row.IntField(m.RETURN),
			Balance:         row.IntField(m.BALANCE),
			UnitPrice:       row.StringField(m.UNITPRICE),
			TotalPrice:      row.StringField(m.TOTALPRICE),
			Cumulative:      row.StringField(m.CUMULATIVE),
			Remark:          row.StringField(m.REMARK),
			Source:          row.StringField(m.SOURCE),
			Destination:     row.StringField(m.DESTINATION),
			CreatedAt:       row.StringField(m.CREATEDAT),
			UpdatedAt:       row.StringField(m.UPDATEDAT),
			DeletedAt:       row.StringField(m.DELETEDAT),
		}
	}
}

func (ts TRANSACTIONS) Values(transactions []domain.Transaction) func(col *sq.Column) {
	created_at := time.Now().UTC()
	return func(col *sq.Column) {
		for _, t := range transactions {
			col.SetString(ts.ID, uuid.New().String())
			col.SetString(ts.MATERIALID, t.MaterialId)
			col.SetString(ts.PROJECT, t.Project)
			col.SetString(ts.COMPANY, t.Company)
			col.SetString(ts.TRANSACTIONTYPE, t.TransactionType)
			col.SetString(ts.REFERENCE_NO, t.ReferenceNo)
			col.SetString(ts.REMARK, t.Remark)
			col.SetString(ts.SOURCE, t.Source)
			col.SetString(ts.DESTINATION, t.Destination)
			col.SetString(ts.CREATEDAT, created_at.Format(DateFormat))
			col.SetString(ts.UPDATEDAT, time.Now().Format("0000-00-00 00:00:00"))
			col.SetString(ts.DELETEDAT, time.Now().Format("0000-00-00 00:00:00"))
			col.SetInt(ts.RECEIVED, t.Received)
			col.SetInt(ts.ISSUED, t.Issued)
			col.SetInt(ts.RETURN, t.Returns)
			col.SetInt(ts.BALANCE, t.Balance)
			col.SetString(ts.UNITPRICE, t.UnitPrice)
			col.SetString(ts.TOTALPRICE, t.TotalPrice)
			col.SetString(ts.CUMULATIVE, t.Cumulative)
		}
	}
}

func (m TRANSACTIONS) TransactionRowMapper(transactions []domain.Transaction) [][]any {
	rowValues := make([][]any, len(transactions))

	updated_at := time.Now().UTC()
	for i, t := range transactions {
		rowValues[i] = []any{t.Id, t.Balance, t.UnitPrice, t.Cumulative, updated_at.Format(DateFormat)}
	}
	fmt.Print(utils.PrettyPrint(rowValues))
	return rowValues
}

// NOTE id should be validated
func (repo Repository) LastTransaction(id string) (domain.Transaction, error) {
	transaction := sq.New[TRANSACTIONS]("transaction")
	var non_deleted = transaction.DELETEDAT.IsNull()

	selected, err := sq.FetchOne(sq.Log(repo.Db), sq.
		From(transaction).
		Where(transaction.ID.EqString(id), non_deleted).
		SetDialect(sq.DialectMySQL),
		transaction.TransactionMapper(),
	)
	if err != nil && err != sql.ErrNoRows {
		return selected, err
	}
	// Get the selected transaction, the compute the previous based on the date
	query, err := sq.FetchOne(repo.Db, sq.
		From(transaction).
		Where(transaction.COMPANY.EqString(selected.Company),
			transaction.PROJECT.EqString(selected.Project),
			transaction.MATERIALID.EqString(selected.MaterialId),
			sq.Expr("createdAt < {}", selected.CreatedAt),
			non_deleted).
		GroupBy(transaction.CREATEDAT.Desc()).
		SetDialect(sq.DialectMySQL),
		transaction.TransactionMapper(),
	)

	if err != nil && err != sql.ErrNoRows {
		return query, err
	}

	return query, nil
}

func (repo Repository) Fetch(filter domain.ComputeFilter) ([]domain.Transaction, error) {
	transaction := sq.New[TRANSACTIONS]("transactions")
	var non_deleted = transaction.DELETEDAT.IsNull()
	var filters []sq.Predicate = []sq.Predicate{non_deleted}

	if len(filter.Ids) > 0 {
		filters = append(filters, transaction.ID.In(filter.Ids))
	}

	if filter.MaterialId != "" {
		filters = append(filters, transaction.MATERIALID.EqString(filter.MaterialId))
	}

	if filter.Date.StartDate != "" || filter.Date.EndDate != "" {
		if filter.Date.StartDate != "" {
			filters = append(filters, sq.Expr("createdAt > {}", filter.Date.StartDate))
		}

		if filter.Date.EndDate != "" {
			filters = append(filters, sq.Expr("createdAt < {}", filter.Date.EndDate))
		}
	}

	query, err := sq.FetchAll(sq.Log(repo.Db), sq.
		From(transaction).
		Where(filters...).
		GroupBy(transaction.CREATEDAT.Desc()).
		SetDialect(sq.DialectMySQL),
		transaction.TransactionMapper(),
	)
	if err != nil && err != sql.ErrNoRows {
		return query, err
	}
	return query, nil
}

func (repo Repository) BulkCreate(transactions []domain.Transaction) error {
	transaction := sq.New[TRANSACTIONS]("")
	tx, err := repo.Db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer tx.Rollback()
	result, err := sq.Exec(tx, sq.
		InsertInto(transaction).
		ColumnValues(transaction.Values(transactions)).SetDialect(sq.DialectMySQL),
	)
	fmt.Println(result.RowsAffected)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repo Repository) BulkUpdate(transactions []domain.Transaction) error {
	transaction := sq.New[TRANSACTIONS]("")
	tx, err := repo.Db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer tx.Rollback()
	tmp := sq.SelectValues{
		Alias:     "tmp",
		Columns:   []string{"id", "balance", "unitPrice", "cumulative", "updatedAt"},
		RowValues: transaction.TransactionRowMapper(transactions),
	}

	result, err := sq.Exec(tx, sq.MySQL.
		Update(transaction).
		Join(tmp, tmp.Field("id").Eq(transaction.ID)).
		Set(
			transaction.BALANCE.Set(tmp.Field("balance")),
			transaction.UNITPRICE.Set(tmp.Field("unitPrice")),
			transaction.CUMULATIVE.Set(tmp.Field("cumulative")),
			transaction.UPDATEDAT.Set(tmp.Field("updatedAt")),
		),
	)
	fmt.Println(result.RowsAffected)

	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
