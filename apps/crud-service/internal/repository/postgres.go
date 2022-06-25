package repository

import (
  "apps/apps/models"
  grpcservice "apps/apps/pkg/grpc"
  "context"
  "database/sql"
  "errors"
  "fmt"
)

type Repo struct {
  db *sql.DB
}

func NewRepository(db *sql.DB) *Repo {
  return &Repo{
    db: db,
  }
}

func (r *Repo) GetPosts(postsId []int32) ([]*grpcservice.GetPostByIdResponse, error) {
  str := ""
  for index, id := range postsId {
    str += fmt.Sprintf("post_id=%d", id)
    if index != len(postsId)-1 {
      str += " OR "
    }
  }
  dataQuery := fmt.Sprintf(`SELECT post_id, user_id, title, body FROM data WHERE %s`, str)
  rows, err := r.db.Query(dataQuery)
  if err != nil {
    return nil, err
  }
  array := []*grpcservice.GetPostByIdResponse{}
  for rows.Next() {
    arr := &grpcservice.GetPostByIdResponse{}
    if err := rows.Scan(&arr.Id, &arr.UserId, &arr.Title, &arr.Body); err != nil {
      return nil, err
    }
    array = append(array, arr)
  }
  return array, nil
}

func (r *Repo) GetPostById(id int) (*models.Data, error) {
  selectQuery := `SELECT post_id, user_id, title, body FROM data WHERE post_id = $1`
  post := &models.Data{}
  row := r.db.QueryRow(selectQuery, id)
  if err := row.Scan(&post.ID, &post.UserID, &post.Title, &post.Body); err != nil {
    return nil, errors.New("post with this id doesn't exist")
  }
  return post, nil
}

func (r *Repo) DeletePostById(id int) error {
  deleteQuery := `DELETE FROM data WHERE post_id = $1`
  ctx := context.Background()
  tx, err := r.db.BeginTx(ctx, nil)
  if err != nil {
    return err
  }
  res, err := tx.ExecContext(ctx, deleteQuery, id)
  status, _ := res.RowsAffected()
  if status == 0 {
    return errors.New("post with id doesn't exist")
  }
  if err != nil {
    tx.Rollback()
    return err
  }
  err = tx.Commit()
  if err != nil {
    return err
  }
  return nil
}

func (r *Repo) UpdatePostById(id int, userId int, title string, body string) error {
  updateQuery := `UPDATE data SET title = $1, body = $2 WHERE post_id = $3 AND user_id = $4`
  ctx := context.Background()
  tx, err := r.db.BeginTx(ctx, nil)
  if err != nil {
    return err
  }
  res, err := tx.ExecContext(ctx, updateQuery, title, body, id, userId)
  status, _ := res.RowsAffected()
  if status == 0 {
    return errors.New("post with id doesn't exist")
  }
  if err != nil {
    tx.Rollback()
    return err
  }
  err = tx.Commit()
  if err != nil {
    return err
  }
  return nil
}
