//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, January 2020
//

package homo

import (
	"fmt"
	"github.com/countstarlight/homo/logger"
	"go.uber.org/zap"
	"os"
	"runtime/debug"
)

// Run service
func Run(handle func(Context) error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "service is stopped with panic: %s\n%s", r, string(debug.Stack()))
		}
	}()
	c, err := newContext()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[%s][%s] failed to create context: %s\n", c.sn, c.in, err.Error())
		logger.S.Errorw("failed to create context", zap.Error(err))
		return
	}
	logger.S.Infof("service starting: ", os.Args)
	err = handle(c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[%s][%s] service is stopped with error: %s\n", c.sn, c.in, err.Error())
		logger.S.Errorw("service is stopped with error", zap.Error(err))
	} else {
		logger.S.Info("service stopped")
	}
}