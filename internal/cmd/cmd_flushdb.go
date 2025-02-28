// Copyright (c) 2022-present, DiceDB contributors
// All rights reserved. Licensed under the BSD 3-Clause License. See LICENSE file in the project root for full license information.

package cmd

import (
	"github.com/dicedb/dice/internal/errors"
	"github.com/dicedb/dice/internal/shardmanager"
	"github.com/dicedb/dice/internal/store"
)

var cFLUSHDB = &CommandMeta{
	Name:      "FLUSHDB",
	HelpShort: "FLUSHDB deletes all keys.",
	Eval:      evalFLUSHDB,
	Execute:   executeFLUSHDB,
}

func init() {
	CommandRegistry.AddCommand(cFLUSHDB)
}

// FLUSHDB deletes all keys.
// The function expects no arguments
//
// Parameters:
//   - c *Cmd: The command context containing the arguments
//   - s *dstore.Store: The data store instance
//
// Returns:
//   - *CmdRes: OK or nil
//   - error: Error if wrong number of arguments
func evalFLUSHDB(c *Cmd, s *store.Store) (*CmdRes, error) {
	if len(c.C.Args) != 0 {
		return cmdResNil, errors.ErrWrongArgumentCount("FLUSHDB")
	}

	store.Reset(s)
	return cmdResOK, nil
}

func executeFLUSHDB(c *Cmd, sm *shardmanager.ShardManager) (*CmdRes, error) {
	for _, shard := range sm.Shards() {
		_, err := evalFLUSHDB(c, shard.Thread.Store())
		if err != nil {
			return nil, err
		}
	}
	return cmdResOK, nil
}
