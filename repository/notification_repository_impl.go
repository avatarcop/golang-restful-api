package repository

import (
	"context"
	"database/sql"
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
	helper.PanicIfError(err)
	defer rows.Close()

	notification := domain.Notification{}

	if rows.Next() {
		err := rows.Scan(&notification.NotificationId)
		helper.PanicIfError(err)
	}

	return notification
}

func (repository *NotificationRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, params domain.NotificationFindByUserId) []domain.Notification {
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
		helper.PanicIfError(err)
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
	helper.PanicIfError(err)
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, sqlParams...)
	helper.PanicIfError(err)

	var notifications []domain.Notification
	for rows.Next() {
		notification := domain.Notification{}
		err := rows.Scan(&notification.NotificationId, &notification.UserId, &notification.Type, &notification.Title, &notification.Message, &notification.IsRead, &notification.IconImageUrl, &notification.ImageUrl, &notification.DateIn)
		helper.PanicIfError(err)
		notifications = append(notifications, notification)
	}
	return notifications
}

func (repository *NotificationRepositoryImpl) Save(ctx context.Context, tx *sql.Tx,
	notification domain.Notification) domain.Notification {
	var notification_id int
	SQL := "INSERT INTO rns_notification" +
		" (notification_id, user_id, title, type, message, is_read, icon_image_url, image_url, sent_status, " +
		" date_in, user_in, status_record) " +
		" VALUES ((select max(notification_id) from rns_notification)+1,$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING notification_id"

	result := tx.QueryRowContext(ctx, SQL, notification.UserId, notification.Title, notification.Type,
		notification.Message, notification.IsRead, notification.IconImageUrl, notification.ImageUrl,
		notification.SentStatus, notification.DateIn, notification.UserIn, notification.StatusRecord)
	if err := result.Scan(&notification_id); err != nil { // scan will release the connection
		panic(err)
	}

	notification.NotificationId = notification_id
	return notification
}
