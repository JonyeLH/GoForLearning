// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redis

import (
	caches "MyGo_middleware/common/Cache"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
	"time"
)

func TestRedisCache(t *testing.T) {
	bm, err := caches.NewCache("engine_config", `{"key":"auth","conn": "47.110.138.80:6379","dbNum":"0","password":"123456"}`)
	if err != nil {
		t.Error("init err")
	}
	timeoutDuration := 300 * time.Second
	if err = bm.Put("astaxie", 1, timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("astaxie") {
		t.Error("check err")
	}

	if err = bm.Put("astaxie", 1, timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if v, _ := redis.Int(bm.Get("astaxie")); v != 1 {
		t.Error("get err")
	}

	r1 := bm.Incr("astaxie")
	fmt.Println(r1)

	if v, _ := redis.Int(bm.Get("astaxie")); v != 2 {
		t.Error("get err")
	}

	r2 := bm.Decr("astaxie")
	fmt.Println(r2)

	if v, _ := redis.Int(bm.Get("astaxie")); v != 1 {
		t.Error("get err")
	}

	//src-test string
	if err = bm.Put("astaxie", "author", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("astaxie") {
		t.Error("check err")
	}

	if v, _ := redis.String(bm.Get("astaxie")); v != "author" {
		t.Error("get err")
	}

	//src-test GetMulti
	if err = bm.Put("astaxie1", "author1", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("astaxie1") {
		t.Error("check err")
	}

	vv, _ := bm.GetMulti([]string{"astaxie", "astaxie1"})
	if len(vv) != 2 {
		t.Error("GetMulti ERROR")
	}
	if v, _ := redis.String(vv[0], nil); v != "author" {
		t.Error("GetMulti ERROR")
	}
	if v, _ := redis.String(vv[1], nil); v != "author1" {
		t.Error("GetMulti ERROR")
	}

}
