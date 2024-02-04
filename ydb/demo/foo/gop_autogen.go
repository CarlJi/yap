package main

import (
	"errors"
	"github.com/goplus/yap/ydb"
	"time"
)

const _ = true

type foo struct {
	ydb.Sql
}

func main() {
//line ydb/demo/foo/foo_ydb.gox:82:1
	ydb.Gopt_AppGen_Main(new(ydb.AppGen), new(foo))
}

var ErrNoEmailAndTel = errors.New("no email and telephone")
//line ydb/demo/foo/foo.gop:9:1
func Rand() string {
//line ydb/demo/foo/foo.gop:10:1
	return ""
}
//line ydb/demo/foo/foo.gop:13:1
func Hmac(pwd string, salt string) string {
//line ydb/demo/foo/foo.gop:14:1
	return ""
}
//line ydb/demo/foo/foo_ydb.gox:5
func (this *foo) Main() {
//line ydb/demo/foo/foo_ydb.gox:5:1
	this.Engine__0("mysql")
//line ydb/demo/foo/foo_ydb.gox:7:1
	this.Table("user v0.1.0", func() {
//line ydb/demo/foo/foo_ydb.gox:8:1
		ydb.Gopt_Table_Gopx_Col__2[[32]byte](this, "id")
//line ydb/demo/foo/foo_ydb.gox:9:1
		ydb.Gopt_Table_Gopx_Col__0[string](this, "spwd")
//line ydb/demo/foo/foo_ydb.gox:10:1
		ydb.Gopt_Table_Gopx_Col__0[string](this, "salt")
//line ydb/demo/foo/foo_ydb.gox:11:1
		ydb.Gopt_Table_Gopx_Col__0[string](this, "nickname")
//line ydb/demo/foo/foo_ydb.gox:12:1
		ydb.Gopt_Table_Gopx_Col__0[string](this, "email")
//line ydb/demo/foo/foo_ydb.gox:13:1
		ydb.Gopt_Table_Gopx_Col__0[string](this, "tel")
//line ydb/demo/foo/foo_ydb.gox:14:1
		ydb.Gopt_Table_Gopx_Col__0[ydb.Date](this, "born")
//line ydb/demo/foo/foo_ydb.gox:15:1
		ydb.Gopt_Table_Gopx_Col__2[[6]ydb.DateTime](this, "ctime")
//line ydb/demo/foo/foo_ydb.gox:17:1
		this.Unique__0("id")
//line ydb/demo/foo/foo_ydb.gox:18:1
		this.Index__0("email")
//line ydb/demo/foo/foo_ydb.gox:19:1
		this.Index__0("tel")
//line ydb/demo/foo/foo_ydb.gox:20:1
		this.Index__0("born")
//line ydb/demo/foo/foo_ydb.gox:21:1
		this.Index__0("ctime")
	})
//line ydb/demo/foo/foo_ydb.gox:24:1
	this.Class("Users", func() {
//line ydb/demo/foo/foo_ydb.gox:25:1
		this.Use("user")
//line ydb/demo/foo/foo_ydb.gox:27:1
		this.Api("register", func(id string, pwd string, nickname string, email string, tel string, ctime time.Time) error {
//line ydb/demo/foo/foo_ydb.gox:28:1
			if email == "" && tel == "" {
//line ydb/demo/foo/foo_ydb.gox:29:1
				return ErrNoEmailAndTel
			}
//line ydb/demo/foo/foo_ydb.gox:31:1
			this.Limit__1(3, "email={email}")
//line ydb/demo/foo/foo_ydb.gox:32:1
			this.Limit__1(3, "tel={tel}")
//line ydb/demo/foo/foo_ydb.gox:34:1
			salt := Rand()
//line ydb/demo/foo/foo_ydb.gox:35:1
			spwd := Hmac(pwd, salt)
//line ydb/demo/foo/foo_ydb.gox:36:1
			this.Insert__1("id", id, "spwd", spwd, "salt", salt, "nickname", nickname, "email", email, "tel", tel, "ctime", ctime)
//line ydb/demo/foo/foo_ydb.gox:39:1
			return nil
		})
//line ydb/demo/foo/foo_ydb.gox:41:1
		this.Call__1("user", "pwd", "nickname", "", "", time.Now())
//line ydb/demo/foo/foo_ydb.gox:42:1
		this.Ret__1(ErrNoEmailAndTel)
//line ydb/demo/foo/foo_ydb.gox:43:1
		this.Call__1("user", "pwd", "nickname", "user@foo.com", "", time.Now())
//line ydb/demo/foo/foo_ydb.gox:44:1
		this.Ret__0(nil)
//line ydb/demo/foo/foo_ydb.gox:45:1
		this.Call__1("user", "pwd", "nickname", "user@foo.com", "13500000000", time.Now())
//line ydb/demo/foo/foo_ydb.gox:46:1
		this.Ret__1(ydb.ErrDuplicated)
//line ydb/demo/foo/foo_ydb.gox:48:1
		this.Api("login", func(id string, pwd string) bool {
//line ydb/demo/foo/foo_ydb.gox:49:1
			var spwd, salt string
//line ydb/demo/foo/foo_ydb.gox:50:1
			this.Query("id={id}")
//line ydb/demo/foo/foo_ydb.gox:51:1
			this.Ret__1("salt", &salt, "spwd", &spwd)
//line ydb/demo/foo/foo_ydb.gox:52:1
			return Hmac(pwd, salt) == spwd
		})
//line ydb/demo/foo/foo_ydb.gox:54:1
		this.Call__1("", "")
//line ydb/demo/foo/foo_ydb.gox:55:1
		this.Ret__1(false)
//line ydb/demo/foo/foo_ydb.gox:56:1
		this.Call__1("user", "pwd")
//line ydb/demo/foo/foo_ydb.gox:57:1
		this.Ret__1(true)
	})
//line ydb/demo/foo/foo_ydb.gox:60:1
	this.Table("article v0.1.0", func() {
//line ydb/demo/foo/foo_ydb.gox:61:1
		ydb.Gopt_Table_Gopx_Col__0[string](this, "id")
//line ydb/demo/foo/foo_ydb.gox:62:1
		ydb.Gopt_Table_Gopx_Col__1[string](this, "author", "name@user")
//line ydb/demo/foo/foo_ydb.gox:63:1
		ydb.Gopt_Table_Gopx_Col__0[string](this, "title")
//line ydb/demo/foo/foo_ydb.gox:64:1
		ydb.Gopt_Table_Gopx_Col__0[ydb.Blob](this, "body")
//line ydb/demo/foo/foo_ydb.gox:66:1
		this.Unique__0("id")
//line ydb/demo/foo/foo_ydb.gox:67:1
		this.Index__0("author")
//line ydb/demo/foo/foo_ydb.gox:69:1
		this.From("oldart v0.9.1", func() {
		})
	})
//line ydb/demo/foo/foo_ydb.gox:75:1
	this.Table("tag v0.1.0", func() {
//line ydb/demo/foo/foo_ydb.gox:76:1
		ydb.Gopt_Table_Gopx_Col__1[string](this, "id", "id@article")
//line ydb/demo/foo/foo_ydb.gox:77:1
		ydb.Gopt_Table_Gopx_Col__0[string](this, "tag")
//line ydb/demo/foo/foo_ydb.gox:79:1
		this.Unique__1([]string{"id", "tag"})
	})
//line ydb/demo/foo/foo_ydb.gox:82:1
	this.Class("Articles", func() {
	})
}
