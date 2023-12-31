/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package migrationscripts

import (
	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
)

var _ plugin.MigrationScript = (*addExcludeBotsAsFirstReviewer20230415)(nil)

type transformationRule202304 struct {
	ExcludedBotsAsFirstReviewer string `gorm:"index;type:varchar(255)"`
}

func (transformationRule202304) TableName() string {
	return "_tool_gitlab_transformation_rules"
}

type addExcludeBotsAsFirstReviewer20230415 struct{}

func (script *addExcludeBotsAsFirstReviewer20230415) Up(basicRes context.BasicRes) errors.Error {
	return basicRes.GetDal().AutoMigrate(&transformationRule202304{})
}

func (*addExcludeBotsAsFirstReviewer20230415) Version() uint64 {
	return 20230415254711
}

func (*addExcludeBotsAsFirstReviewer20230415) Name() string {
	return "add excluded_bots_as_first_reviewer for _tool_gitlab_transformation_rules"
}
