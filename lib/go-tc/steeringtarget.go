package tc

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/apache/trafficcontrol/lib/go-util"
)

type SteeringTarget struct {
	DeliveryService   DeliveryServiceName `json:"deliveryService" db:"deliveryservice_name"`
	DeliveryServiceID int                 `json:"deliveryServiceId" db:"deliveryservice"`
	Target            DeliveryServiceName `json:"target" db:"target_name"`
	TargetID          int                 `json:"targetId" db:"target"`
	Type              string              `json:"type" db:"type"`      // TODO enum?
	TypeID            int                 `json:"typeId" db:"type_id"` // TODO enum?
	Value             util.JSONIntStr     `json:"value" db:"value"`
}

type SteeringTargetNullable struct {
	DeliveryService   *DeliveryServiceName `json:"deliveryService" db:"deliveryservice_name"`
	DeliveryServiceID *uint64              `json:"deliveryServiceId" db:"deliveryservice"`
	Target            *DeliveryServiceName `json:"target" db:"target_name"`
	TargetID          *uint64              `json:"targetId" db:"target"`
	Type              *string              `json:"type" db:"type_name"` // TODO enum?
	TypeID            *int                 `json:"typeId" db:"type_id"` // TODO enum?
	Value             *util.JSONIntStr     `json:"value" db:"value"`
}

func (st SteeringTargetNullable) Validate(tx *sql.Tx) error {
	errs := []string{}
	if st.TypeID == nil {
		errs = append(errs, "missing typeId")
	}
	if st.Value == nil {
		errs = append(errs, "missing value")
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "; "))
	}
	return nil
}
