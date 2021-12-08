package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
	"strconv"
	"strings"
)

type NotificationRepositoryImpl struct {
}

func NewNotificationRepository() NotificationRepository {
	return &NotificationRepositoryImpl{}
}

func (repository *NotificationRepositoryImpl) Health(ctx context.Context, tx *sql.Tx) int {
	return 1
}

func (repository *NotificationRepositoryImpl) HealthCheckDB(ctx context.Context, tx *sql.Tx) domain.Notification {
	SQL := "SELECT notification_id FROM rns_notification where notification_id = $1"
	rows, err := tx.QueryContext(ctx, SQL, 1)
	helper.PanicIfError(err, "error notification repository at func healthcheckdb when execute query")
	defer rows.Close()

	notification := domain.Notification{}

	if rows.Next() {
		err := rows.Scan(&notification.NotificationId)
		helper.PanicIfError(err, "error notification repository at func healthcheckdb when scan row")
	}

	return notification
}

func (repository *NotificationRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, params helper.FilterNotificationFindByUserId) []domain.NotificationFindById {
	no := 1
	SQL := "SELECT notification_id, user_id, type, title, message, is_read, icon_image_url, image_url, date_in" +
		" FROM rns_notification" +
		" WHERE user_id = $1"
	sqlParams := []interface{}{params.UserId}

	if strings.Title(params.Type) == "Info" || strings.Title(params.Type) == "Promo" {
		no++
		SQL = fmt.Sprintf(SQL+" and type = $%d", no)
		sqlParams = append(sqlParams, params.Type)
	}

	if len(params.IsRead) > 0 {
		no++
		SQL = fmt.Sprintf(SQL+" and is_read = $%d", no)
		isRead, err := strconv.ParseBool(params.IsRead)
		helper.PanicIfError(err, "error notification repository at func findbyuserid when convert is read")
		sqlParams = append(sqlParams, isRead)
	}

	if params.StatusRecord != "" {
		no++
		SQL = fmt.Sprintf(SQL+" and status_record = $%d", no)
		sqlParams = append(sqlParams, strings.ToUpper(params.StatusRecord))
	} else {
		no++
		SQL = fmt.Sprintf(SQL+" and status_record != $%d", no)
		sqlParams = append(sqlParams, "DELETED")
	}

	SQL = SQL + " ORDER BY date_in DESC"

	if params.Page > 0 && params.ItemsPerPage > 0 {
		no++
		SQL = fmt.Sprintf(SQL+" OFFSET $%d", no)
		no++
		SQL = fmt.Sprintf(SQL+" LIMIT $%d", no)
		sqlParams = append(sqlParams, (params.Page-1)*params.ItemsPerPage)
		sqlParams = append(sqlParams, params.ItemsPerPage)
	} else {
		if params.ItemsPerPage > 0 {
			no++
			SQL = fmt.Sprintf(SQL+" LIMIT $%d", no)
			sqlParams = append(sqlParams, params.ItemsPerPage)
		}
	}

	stmt, err := tx.PrepareContext(ctx, SQL)
	helper.PanicIfError(err, "error notification repository at func findbyuserid when prepare context")
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, sqlParams...)
	helper.PanicIfError(err, "error notification repository at func findbyuserid when execute query context")

	var notifications []domain.NotificationFindById
	for rows.Next() {
		notification := domain.NotificationFindById{}
		err := rows.Scan(&notification.NotificationId, &notification.UserId, &notification.Type, &notification.Title, &notification.Message, &notification.IsRead, &notification.IconImageUrl, &notification.ImageUrl, &notification.DateIn)
		helper.PanicIfError(err, "error notification repository at func findbyuserid when scan row")
		notifications = append(notifications, notification)
	}
	return notifications
}

func (repository *NotificationRepositoryImpl) Save(ctx context.Context, tx *sql.Tx,
	notification domain.Notification) domain.Notification {
	var notification_id int
	SQL := "INSERT INTO rns_notification" +
		" (user_id, title, type, message, is_read, icon_image_url, image_url, sent_status, " +
		" date_in, user_in, status_record) " +
		" VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING notification_id"

	result := tx.QueryRowContext(ctx, SQL, notification.UserId, notification.Title, notification.Type,
		notification.Message, notification.IsRead, notification.IconImageUrl, notification.ImageUrl,
		notification.SentStatus, notification.DateIn, notification.UserIn, notification.StatusRecord)
	if err := result.Scan(&notification_id); err != nil { // scan will release the connection
		helper.PanicIfError(err, "error notification repository at func save when scan row")
	}

	notification.NotificationId = notification_id
	return notification
}

func (repository *NotificationRepositoryImpl) FindOneByNotificationId(ctx context.Context, tx *sql.Tx, notificationId int) (domain.Notification, error) {
	SQL := "select notification_id, user_id, title, type, message, is_read, " +
		" icon_image_url, image_url, " +
		" date_in, user_in, date_up, user_up, status_record " +
		" FROM rns_notification " +
		" WHERE notification_id = ($1)"
	rows, err := tx.QueryContext(ctx, SQL, notificationId)
	helper.PanicIfError(err, "error notification repository at func findonebynotificationid when execute query")
	defer rows.Close()

	notification := domain.Notification{}
	if rows.Next() {
		err := rows.Scan(&notification.NotificationId, &notification.UserId,
			&notification.Type, &notification.Title, &notification.Message,
			&notification.IsRead, &notification.IconImageUrl, &notification.ImageUrl,
			&notification.DateIn, &notification.UserIn, &notification.DateUp,
			&notification.UserUp, &notification.StatusRecord)
		helper.PanicIfError(err, "error notification repository at func findonebynotificationid when scan row")
		return notification, nil
	} else {
		return notification, errors.New("notification is not found")
	}
}

func (repository *NotificationRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, notification domain.Notification) domain.Notification {
	var notificationId int
	SQL := "UPDATE rns_notification SET user_id = $1, type = $2, title = $3, message = $4, " +
		" is_read = $5, icon_image_url = $6, image_url = $7, status_record = $8, date_up = $9, user_up = $10, " +
		" request_raw = $11, response_raw = $12" +
		" WHERE notification_id = $13 RETURNING notification_id"
	result := tx.QueryRowContext(ctx, SQL, notification.UserId, notification.Type, notification.Title, notification.Message,
		notification.IsRead, notification.IconImageUrl, notification.ImageUrl, notification.StatusRecord,
		notification.DateUp, notification.UserUp, notification.RequestRaw, notification.ResponseRaw, notification.NotificationId)
	if err := result.Scan(&notificationId); err != nil { // scan will release the connection
		helper.PanicIfError(err, "error notification repository at func update when scan row")
	}

	return notification
}

func (repository *NotificationRepositoryImpl) UpdateRead(ctx context.Context, tx *sql.Tx, notification domain.Notification) domain.Notification {
	var notificationId int
	SQL := "UPDATE rns_notification SET is_read = $1, status_record = $2, date_up = $3, user_up = $4 " +
		" WHERE notification_id = $5 RETURNING notification_id"
	result := tx.QueryRowContext(ctx, SQL, notification.IsRead, notification.StatusRecord,
		notification.DateUp, notification.UserUp, notification.NotificationId)
	if err := result.Scan(&notificationId); err != nil { // scan will release the connection
		helper.PanicIfError(err, "error notification repository at func updateread when scan row")
	}

	return notification
}

func (repository *NotificationRepositoryImpl) DeleteSoft(ctx context.Context, tx *sql.Tx, notification domain.Notification) domain.Notification {
	var notificationId int
	SQL := "UPDATE rns_notification SET status_record = $1, date_up = $2, user_up = $3 " +
		" WHERE notification_id = $4 RETURNING notification_id"
	result := tx.QueryRowContext(ctx, SQL, notification.StatusRecord,
		notification.DateUp, notification.UserUp, notification.NotificationId)
	if err := result.Scan(&notificationId); err != nil { // scan will release the connection
		helper.PanicIfError(err, "error notification repository at func deletesoft when execute query")
	}

	return notification
}

func (repository *NotificationRepositoryImpl) DeleteHard(ctx context.Context, tx *sql.Tx, notification domain.Notification) {
	SQL := "DELETE FROM rns_notification WHERE notification_id = $1"
	_, err := tx.QueryContext(ctx, SQL, notification.NotificationId)
	helper.PanicIfError(err, "error notification repository at func delete hard when execute query")
}
