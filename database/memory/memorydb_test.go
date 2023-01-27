// Copyright 2022 bnb-chain. All Rights Reserved.
//
// Distributed under MIT license.
// See file LICENSE for detail or copy at https://opensource.org/licenses/MIT

package memory

import (
	"testing"

	"github.com/zecrey-labs/zecrey-smt/database"
	"github.com/zecrey-labs/zecrey-smt/database/dbtest"
)

func TestMemoryDB(t *testing.T) {
	t.Run("DatabaseSuite", func(t *testing.T) {
		dbtest.TestDatabaseSuite(t, func() database.TreeDB {
			return NewMemoryDB()
		})
	})
}
