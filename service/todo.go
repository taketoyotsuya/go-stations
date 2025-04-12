package service

import (
	"context"
	"database/sql"
	"log"

	"github.com/TechBowl-japan/go-stations/model"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
func (s *TODOService) CreateTODO(ctx context.Context, subject, description string) (*model.TODO, error) {
	const (
		insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
		confirm = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	// INSERT クエリを準備
	stmt, err := s.db.PrepareContext(ctx, insert) //stmtはsql.Stmt型（ステートメント）
	//nilはポインタ型などの特定の変数が値を持たないことを示すために使用される
	//PrepareContextは、SQL文を実行するための準備を行うメソッドで、SQL文を実行するためのステートメントを返す
	if err != nil {
		log.Println("PrepareContextでエラー発生")
		return nil, err
	}
	//ステートメントをクローズ(deferを使用することにより、周囲の関数が終了した後に自動的にクローズされる)
	defer stmt.Close()

	// INSERT クエリを実行
	//ExecContextは、行を返さずにクエリを実行します。引数は、クエリ内のプレースホルダ（後から値や文字を入力できるように、一時的に確保しておく値や文字、場所のこと）パラメータのものです。
	result, err := stmt.ExecContext(ctx, subject, description)
	if err != nil {
		log.Println("ExecContextでエラー発生")
		return nil, err
	}

	// 挿入されたレコードの ID を取得（ExecContextメソッドの戻り値から保存したTODOのIDを取得）
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("LastInsertIdでエラー発生")
		return nil, err
	}

	// 挿入されたレコードを取得（IDを利用して保存したTODOをQueryRowContextメソッドを利用して読み取り）
	row := s.db.QueryRowContext(ctx, confirm, id)
	todo := &model.TODO{}

	//todo.Studentが何なのかを確認するため
	log.Println(todo)
	log.Println(&todo.Subject)

	err = row.Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		log.Println("Row.Scanでエラー発生:", err)
		return nil, err
	}

	return todo, nil
}

// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)

	return nil, nil
}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	return nil, nil
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	return nil
}
