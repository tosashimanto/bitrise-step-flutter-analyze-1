package main

import (
	"testing"
)

func TestHasAnalyzeError(t *testing.T) {
	cases := []struct {
		title        string
		cmdOutput    string
		failSeverity string
		want         bool
	}{
		{
			"no error level violation fails on info",
			`info • This class inherits from a class marked as @immutable, and therefore should be immutable (all instance fields must be final) • lib/pages/home/home_page.dart:17:7 • must_be_immutable
			info • This class inherits from a class marked as @immutable, and therefore should be immutable (all instance fields must be final) • lib/pages/login/login_page.dart:11:7 • must_be_immutable
			info • Unused import: 'package:silkthread/service/ads_manager.dart' • lib/pages/post/image_edit.dart:8:8 • unused_import
			info • Unused import: 'dart:ui' • lib/pages/post/image_full_view.dart:1:8 • unused_import
			info • Unused import: 'package:cloud_firestore/cloud_firestore.dart' • lib/pages/post/image_full_view.dart:4:8 • unused_import
			info • Unused import: 'package:font_awesome_flutter/font_awesome_flutter.dart' • lib/pages/post/image_full_view.dart:6:8 • unused_import
			`,
			"warning",
			false,
		},
		{
			"contains error and info level violation fails on error",
			`info • Unused import: 'dart:math' • lib/package.dart:3:8 • unused_import
			error • Expected to find ';' • lib/package.dart:3:8 • expected_token
			`,
			"error",
			true,
		},
		{
			"contains error and info level violation fails on warning",
			`info • Unused import: 'dart:math' • lib/package.dart:3:8 • unused_import
			error • Expected to find ';' • lib/package.dart:3:8 • expected_token
			`,
			"warning",
			true,
		},
		{
			"contains error and info level violation fails on info",
			`info • Unused import: 'dart:math' • lib/package.dart:3:8 • unused_import
			error • Expected to find ';' • lib/package.dart:3:8 • expected_token
			`,
			"info",
			true,
		},
		{
			"contains warning level violation fails on error",
			`warning • Unused import: 'dart:math' • lib/package.dart:3:8 • unused_import
			warning • Expected to find ';' • lib/package.dart:3:8 • expected_token
			`,
			"error",
			false,
		},
		{
			"contains warning level violation fails on warning",
			`warning • Unused import: 'dart:math' • lib/package.dart:3:8 • unused_import
			warning • Expected to find ';' • lib/package.dart:3:8 • expected_token
			`,
			"warning",
			true,
		},
		{
			"contains warning level violation fails on info",
			`warning • Unused import: 'dart:math' • lib/package.dart:3:8 • unused_import
			warning • Expected to find ';' • lib/package.dart:3:8 • expected_token
			`,
			"info",
			true,
		},
	}

	for _, tt := range cases {
		if got := hasAnalyzeError(tt.cmdOutput, tt.failSeverity); got != tt.want {
			t.Errorf("%s: got %t want %t", tt.title, got, tt.want)
		}
	}
}
