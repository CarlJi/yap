/*
 * Copyright (c) 2024 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ydb

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"github.com/goplus/gop/ast"
)

const (
	GopPackage = true
)

// -----------------------------------------------------------------------------

type dbTable = Table
type dbClass = Class

type Sql struct {
	driver  string
	tables  map[string]*Table
	classes map[string]*Class
	db      *sql.DB
	*dbTable
	*dbClass
}

func (p *Sql) initSql() {
	p.tables = make(map[string]*Table)
	p.classes = make(map[string]*Class)
}

// Engine initializes database by specified engine name.
func (p *Sql) Engine__0(name string, src ...ast.Node) {
	if defaultDataSource, ok := engineDataSource[name]; ok {
		db, err := sql.Open(name, defaultDataSource)
		if err != nil {
			log.Panicln("sql.Open:", err)
		}
		p.db = db
		p.driver = name
	}
}

// Engine returns engine name of the database.
func (p *Sql) Engine__1() string {
	return p.driver
}

// Table creates a new table by a spec.
func (p *Sql) Table(nameVer string, spec func(), src ...ast.Node) {
	pos := strings.IndexByte(nameVer, ' ') // user v0.1.0
	if pos < 0 {
		log.Panicln("table name should have a version: eg. `user v0.1.0`")
	}
	name, ver := nameVer[:pos], strings.TrimLeft(nameVer[pos+1:], " \t")
	tbl := newTable(name, ver)
	p.dbTable = tbl
	p.tables[name] = tbl
	spec()
	tbl.create(context.TODO(), p)
	p.dbTable = nil
}

// Class creates a new class by a spec.
func (p *Sql) Class(name string, spec func(), src ...ast.Node) {
	cls := newClass(name)
	p.dbClass = cls
	p.classes[name] = cls
	spec()
	cls.create(context.TODO(), p)
	p.dbClass = nil
}

// -----------------------------------------------------------------------------

var (
	engineDataSource map[string]string // engineName => defaultDataSource
)

// Register registers a engine and its default data source.
func Register(name, defaultDataSource string) {
	engineDataSource[name] = defaultDataSource
}

// -----------------------------------------------------------------------------

type AppGen struct {
}

func (p *AppGen) initApp() {
}

func Gopt_AppGen_Main(app interface{ initApp() }, workers ...interface{ initSql() }) {
	app.initApp()
	if me, ok := app.(interface{ MainEntry() }); ok {
		me.MainEntry()
	}
	for _, worker := range workers {
		worker.initSql()
		worker.(interface{ Main() }).Main()
	}
}

// -----------------------------------------------------------------------------
