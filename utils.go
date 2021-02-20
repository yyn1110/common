package common

import (

	"database/sql"
	"errors"
	"fmt"
	"net/smtp"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"strings"
	"syscall"
	"testing"
	"sync"
)

//////////////////////////////////////////////////////////////////////
// check param invalid, if not assert
//////////////////////////////////////////////////////////////////////
func Assert(expr bool, message string) {
	if !expr {
		panic(message)
	}
}

func TestAssert(t *testing.T, expr bool, message string) {
	if !expr {
		t.Fatal(message)
	}
}

func AssertEqual(t *testing.T, s string, x, y interface{}) {
	if !reflect.DeepEqual(x, y) {
		t.Fatalf("%s: %#v, %#v", s, x, y)
	}
}

func AssertNotEqual(t *testing.T, s string, x, y interface{}) {
	if reflect.DeepEqual(x, y) {
		t.Fatalf("%s: %#v", s, x)
	}
}

func CheckParam(expr bool) {
	Assert(expr, "invalid param")
}

func WaitKill() {
	// Wait for terminating signal
	sc := make(chan os.Signal, 2)
	signal.Notify(sc, syscall.SIGTERM, syscall.SIGINT)
	<-sc
}

// get function name
func GetFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

//create sql template table, the table name must start with <template_domain_id>
//and the string <template_domain_id> only in the template table name
func CreateTemplateTable(driver *sql.DB, templateTable, domain string) error {
	if driver == nil || templateTable == "" || domain == "" {
		return errors.New("check param failed")
	}
	SQL, err := generateRealSql(templateTable, domain)
	if err != nil {
		return err
	}
	return CreateTable(driver, SQL)
}

func CreateDatabase(driver *sql.DB, database string) error {
	if len(database) <= 0 {
		fmt.Printf("database [%s] is invalid", database)
		return ErrInvalidParam
	}

	SQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET = utf8 DEFAULT COLLATE = utf8_unicode_ci", database)

	if _, err := driver.Exec(SQL); err != nil {
		fmt.Printf("create database [%s] failed: err[%v]", database, err)
		return err
	}

	return nil
}

func DropDatabase(driver *sql.DB, database string) error {
	if len(database) <= 0 {
		fmt.Printf("database [%s] is invalid", database)
		return ErrInvalidParam
	}

	SQL := fmt.Sprintf("DROP DATABASE IF EXISTS %s", database)

	if _, err := driver.Exec(SQL); err != nil {
		fmt.Printf("drop database [%s] failed: err[%v]", database, err)
		return err
	}

	return nil
}

func CreateTable(driver *sql.DB, table string) error {
	if driver == nil || table == "" {
		return errors.New("check param failed")
	}
	if _, err := driver.Exec(table); err != nil {
		return err
	}
	return nil
}

func DropTemplateTable(driver *sql.DB, templateTable, domain string) error {
	if driver == nil || len(templateTable) <= 0 || len(domain) <= 0 {
		return ErrInvalidParam
	}
	tableName, err := generateRealTableName(templateTable, domain)
	if err != nil {
		return err
	}
	return DropTable(driver, tableName)
}

func DropTable(driver *sql.DB, tableName string) error {
	if driver == nil || len(tableName) <= 0 {
		return ErrInvalidParam
	}
	SQL := fmt.Sprintf("DROP TABLE %s", tableName)
	stmt, err := driver.Prepare(SQL)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func TruncateTemplateTable(driver *sql.DB, templateTable, domain string) error {
	if driver == nil || len(templateTable) <= 0 || len(domain) <= 0 {
		return ErrInvalidParam
	}
	tableName, err := generateRealTableName(templateTable, domain)
	if err != nil {
		return err
	}
	return TruncateTable(driver, tableName)
}

func TruncateTable(driver *sql.DB, tableName string) error {
	if driver == nil || len(tableName) <= 0 {
		return ErrInvalidParam
	}
	SQL := fmt.Sprintf("TRUNCATE TABLE %s", tableName)
	stmt, err := driver.Prepare(SQL)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func generateRealTableName(sqlStr, domain string) (string, error) {
	var tableName string
	split1 := strings.Split(sqlStr, "<template_domain_id>")
	if len(split1) != 2 {
		return tableName, errors.New("uknown sql table template")
	}
	split2 := strings.Split(split1[1], " ")
	if len(split2) <= 0 {
		return tableName, errors.New("uknown sql table template")
	}
	tableName = domain + split2[0]
	return tableName, nil
}

func generateRealSql(sqlStr, domain string) (string, error) {
	var sql string
	split := strings.Split(sqlStr, "<template_domain_id>")
	if len(split) != 2 {
		return sql, errors.New("uknown sql table template")
	}
	sql = split[0] + domain + split[1]
	return sql, nil
}

// send email by user := "xxx@test.com" with password
// through host := "smtp.test.com:25" to := "yyy@test.com;zzz@test.com"
func SendMail(user, password, smtpHost, to, subject, mailType, content string) error {
	addr := strings.Split(smtpHost, ":")
	auth := smtp.PlainAuth("", user, password, addr[0])
	var contentType string
	if mailType == "html" {
		contentType = "Content-Type: text/html" + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	body := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " +
		subject + "\r\n" + contentType + "\r\n\r\n" + content)
	// dest mail split by ;
	list := strings.Split(to, ";")
	return smtp.SendMail(smtpHost, auth, user, list, body)
}

// DBOptions 数据库选项
type DBOptions struct {
	Addr        string
	User        string
	Password    string
	Database    string
	MaxOpenConn int
	MaxIdleConn int
}

// OpenDatabase 打开数据库
func OpenDatabase(driver string, opts DBOptions) (*sql.DB, error) {
	source := fmt.Sprintf("%s:%s@tcp(%s)/%s", opts.User, opts.Password, opts.Addr, opts.Database)
	db, err := sql.Open(driver, source)
	if err != nil {
		fmt.Printf("[OpenDB] open(%s, %s): %v", driver, source, err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		fmt.Printf("[OpenDB] ping(%s, %s): %v", driver, source, err)
		return nil, err
	}
	db.SetMaxOpenConns(opts.MaxOpenConn)
	db.SetMaxIdleConns(opts.MaxIdleConn)
	return db, nil
}

func InLock(lock *sync.Mutex, fun func()) {
	lock.Lock()
	defer lock.Unlock()

	fun()
}

func InReadLock(lock *sync.RWMutex, fun func()) {
	lock.RLock()
	defer lock.RUnlock()

	fun()
}

func InWriteLock(lock *sync.RWMutex, fun func()) {
	lock.Lock()
	defer lock.Unlock()

	fun()
}
