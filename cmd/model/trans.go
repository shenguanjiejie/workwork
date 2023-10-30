package model

/*
Copyright © 2022 shenguanjiejie <835166018@qq.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

type TransInfo struct {
	From        string `json:"from" bson:"from"`
	To          string `json:"to" bson:"to"`
	TransResult []struct {
		Src string `json:"src" bson:"src"`
		Dst string `json:"dst" bson:"dst"`
	} `json:"trans_result" bson:"trans_result"`
}
