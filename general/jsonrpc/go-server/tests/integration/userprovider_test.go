// +build integration

/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package integration

import (
	"context"
	"testing"
)
import (
	"github.com/stretchr/testify/assert"
)

func TestGetUserA000(t *testing.T) {
	user := &JsonRPCUser{}
	err := userProvider.GetUser(context.TODO(), []interface{}{"A000"}, user)
	assert.Nil(t, err)
	assert.Equal(t, "0", user.ID)
	assert.Equal(t, "Alex Stocks", user.Name)
	assert.Equal(t, int64(31), user.Age)
	assert.Equal(t, "MAN", user.Sex)
	assert.NotNil(t, user.Time)
}

func TestGetUserA001(t *testing.T) {
	user := &JsonRPCUser{}
	err := userProvider.GetUser(context.TODO(), []interface{}{"A001"}, user)
	assert.Nil(t, err)
	assert.Equal(t, "001", user.ID)
	assert.Equal(t, "ZhangSheng", user.Name)
	assert.Equal(t, int64(18), user.Age)
	assert.Equal(t, "MAN", user.Sex)
	assert.NotNil(t, user.Time)
}

func TestGetUserA002(t *testing.T) {
	user := &JsonRPCUser{}
	err := userProvider.GetUser(context.TODO(), []interface{}{"A002"}, user)
	assert.Nil(t, err)
	assert.Equal(t, "002", user.ID)
	assert.Equal(t, "Lily", user.Name)
	assert.Equal(t, int64(20), user.Age)
	assert.Equal(t, "WOMAN", user.Sex)
	assert.NotNil(t, user.Time)
}

func TestGetUserA003(t *testing.T) {
	user := &JsonRPCUser{}
	err := userProvider.GetUser(context.TODO(), []interface{}{"A003"}, user)
	assert.Nil(t, err)
	assert.Equal(t, "113", user.ID)
	assert.Equal(t, "Moorse", user.Name)
	assert.Equal(t, int64(30), user.Age)
	assert.Equal(t, "WOMAN", user.Sex)
	assert.NotNil(t, user.Time)
}

func TestGetUser0(t *testing.T) {
	user, err := userProvider.GetUser0("A003", "Moorse")
	assert.Nil(t, err)
	assert.NotNil(t, user)

	user, err = userProvider.GetUser0("A003", "MOORSE")
	assert.NotNil(t, err)
}

// FIXME: this doesn't work with jsonrpc
func TestGetUser2(t *testing.T) {
	user := &JsonRPCUser{}
	err := userProvider.GetUser2(context.TODO(), []interface{}{int32(64)}, user)
	assert.Nil(t, err)
	assert.Equal(t, "64", user.ID)
}

func TestGetUser3(t *testing.T) {
	err := userProvider.GetUser3()
	assert.Nil(t, err)
}

func TestGetUsers(t *testing.T) {
	users, err := userProvider.GetUsers([]interface{}{[]interface{}{"A002", "A003"}})
	assert.Nil(t, err)
	assert.Equal(t, "Lily", users[0].Name)
	assert.Equal(t, "Moorse", users[1].Name)
}
