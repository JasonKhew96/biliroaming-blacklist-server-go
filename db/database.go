package db

import (
	"biliroaming-blacklist-server-go/config"
	"biliroaming-blacklist-server-go/models"
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Database struct {
	db      *sql.DB
	context context.Context
}

// bilibili_users

func (db *Database) UpsertBiliUser(biliUser *models.BilibiliUser) error {
	return biliUser.Upsert(db.context, db.db, true, []string{"uid"}, boil.Whitelist("counter", "is_whitelist", "ban_until"), boil.Infer())
}

func (db *Database) BanBiliUser(uid int64, banUntil time.Duration) error {
	biliUser, err := db.GetBiliUser(uid)
	if err != nil && err == sql.ErrNoRows {
		return db.UpsertBiliUser(&models.BilibiliUser{
			UID:      uid,
			BanUntil: null.TimeFrom(time.Now().Add(banUntil)),
		})
	} else if err != nil {
		return err
	}
	biliUser.BanUntil = null.TimeFrom(time.Now().Add(banUntil))
	return db.UpsertBiliUser(biliUser)
}

func (db *Database) UnbanBiliUser(uid int64) error {
	biliUser, err := db.GetBiliUser(uid)
	if err != nil {
		return err
	}
	biliUser.BanUntil = null.TimeFrom(time.Time{})
	return db.UpsertBiliUser(biliUser)
}

func (db *Database) WhiteBiliUser(uid int64, white bool) error {
	biliUser, err := db.GetBiliUser(uid)
	if err != nil && err == sql.ErrNoRows {
		return db.UpsertBiliUser(&models.BilibiliUser{
			UID:         uid,
			IsWhitelist: white,
		})
	} else if err != nil {
		return err
	}
	biliUser.IsWhitelist = white
	return db.UpsertBiliUser(biliUser)
}

func (db *Database) GetBiliUser(uid int64) (*models.BilibiliUser, error) {
	return models.BilibiliUsers(models.BilibiliUserWhere.UID.EQ(uid)).One(db.context, db.db)
}

func (db *Database) IncBiliUserCounter(uid int64) (int64, error) {
	biliUser, err := db.GetBiliUser(uid)
	if err == sql.ErrNoRows {
		if err := db.UpsertBiliUser(&models.BilibiliUser{
			UID:     uid,
			Counter: 1,
		}); err != nil {
			return -1, err
		}
	} else if err != nil {
		return -1, err
	}
	biliUser.Counter++
	return biliUser.Update(db.context, db.db, boil.Infer())
}

// admins

func (db *Database) UpsertAdmin(admin *models.Admin) error {
	return admin.Upsert(db.context, db.db, true, []string{"id"}, boil.Whitelist("permissions"), boil.Infer())
}

func (db *Database) GetAdmin(id int64) (*models.Admin, error) {
	return models.Admins(models.AdminWhere.ID.EQ(id)).One(db.context, db.db)
}

func (db *Database) RemoveAdmin(id int64) (int64, error) {
	admin, err := db.GetAdmin(id)
	if err != nil {
		return -1, err
	}
	return admin.Delete(db.context, db.db)
}

// users

func (db *Database) GetUserRecords(uid int64) (models.RecordSlice, error) {
	return models.Records(models.RecordWhere.UID.EQ(uid), qm.OrderBy("modified_at DESC", qm.Limit(8))).All(db.context, db.db)
}

func (db *Database) GetTotalUser() (int64, error) {
	return models.BilibiliUsers().Count(db.context, db.db)
}

func (db *Database) InsertRecord(uid int64, description string, chatId, msgId int64) (int, error) {
	record := models.Record{
		UID:         uid,
		Description: description,
		ChatID:      null.Int64From(chatId),
		MessageID:   null.Int64From(msgId),
	}
	// https://t.me/c/1755746887/1236
	err := record.Insert(db.context, db.db, boil.Infer())
	return record.RecordID, err
}

// records

func (db *Database) GetTotalRecord() (int64, error) {
	return models.Records().Count(db.context, db.db)
}

func (db *Database) GetRecordCount(uid int64) (int64, error) {
	return models.Records(models.RecordWhere.UID.EQ(uid)).Count(db.context, db.db)
}

func (db *Database) GetRecord(uid int64, offset int) (*models.Record, error) {
	return models.Records(models.RecordWhere.UID.EQ(uid), qm.OrderBy("modified_at DESC"), qm.Offset(offset)).One(db.context, db.db)
}

func (db *Database) GetRecords(uid int64, limit int) (models.RecordSlice, error) {
	return models.Records(models.RecordWhere.UID.EQ(uid), qm.OrderBy("modified_at DESC"), qm.Limit(limit)).All(db.context, db.db)
}

// reports

const (
	FILE_TYPE_NONE int16 = iota
	FILE_TYPE_ANIMATION
	FILE_TYPE_DOCUMENT
	FILE_TYPE_PHOTO
	FILE_TYPE_VIDEO
)

func (db *Database) GetTotalReport() (int64, error) {
	return models.Reports().Count(db.context, db.db)
}

func (db *Database) InsertReport(uid int64, description string, fileType int16, fileId string) (int, error) {
	report := models.Report{
		UID:         uid,
		Description: description,
		FileType:    fileType,
		FileID:      fileId,
	}
	err := report.Insert(db.context, db.db, boil.Infer())
	return report.ReportID, err
}

func (db *Database) UpdateReport(reportId int, fileType int16, fileId string) (int64, error) {
	report, err := db.GetReport(reportId)
	if err != nil {
		return -1, err
	}
	report.FileType = fileType
	report.FileID = fileId
	return report.Update(db.context, db.db, boil.Infer())
}

func (db *Database) GetReport(reportId int) (*models.Report, error) {
	return models.Reports(models.ReportWhere.ReportID.EQ(reportId)).One(db.context, db.db)
}

func New(config config.DatabaseConfig) (*Database, error) {
	// boil.DebugMode = true
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s", config.Host, config.Port, config.Name, config.User, config.Pass))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{
		db:      db,
		context: context.Background(),
	}, nil
}
