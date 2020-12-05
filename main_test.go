package main

import (
	"fmt"
	"testing"
)

func TestHasAnalyzeError(t *testing.T) {
	cases := []struct {
		title        string
		cmdOutput    string
		failSeverity string
		want         bool
	}{
		// Info severity
		{
			"レベルにerrorなし、メッセージの中にerrorワードあり",
			`info • This class inherits from a class marked as @immutable, and therefore should be immutable (all instance fields must be final) • lib/pages/home/home_page.dart:17:7 • must_be_immutable
			info • This class inherits from a class marked as @immutable, and therefore should be immutable (all instance fields must be final) • lib/pages/login/login_page.dart:11:7 • must_be_immutable
			info • Unused import: 'package:silkthread/service/ads_manager.dart' • lib/pages/post/image_edit.dart:8:8 • unused_import
			info • Unused import: 'dart:ui' • lib/pages/post/image_full_view.dart:1:8 • unused_import
			warning • Unused import: 'error:cloud_firestore/cloud_firestore.dart' • lib/pages/post/image_full_view.dart:4:8 • unused_import
			info • Unused import: 'package:font_awesome_flutter/font_awesome_flutter.dart' • lib/pages/post/image_full_view.dart:6:8 • unused_import
			`,
			"error",
			false,
		},
		{
			"レベルにerrorあり、メッセージの中にerrorワードあり",
			`info • This class inherits from a class marked as @immutable, and therefore should be immutable (all instance fields must be final) • lib/pages/home/home_page.dart:17:7 • must_be_immutable
			info • This class inherits from a class marked as @immutable, and therefore should be immutable (all instance fields must be final) • lib/pages/login/login_page.dart:11:7 • must_be_immutable
			info • Unused import: 'package:silkthread/service/ads_manager.dart' • lib/pages/post/image_edit.dart:8:8 • unused_import
			info • Unused import: 'dart:ui' • lib/pages/post/image_full_view.dart:1:8 • unused_import
			info • Unused import: 'error:cloud_firestore/cloud_firestore.dart' • lib/pages/post/image_full_view.dart:4:8 • unused_import
			error • Unused import: 'package:font_awesome_flutter/font_awesome_flutter.dart' • lib/pages/post/image_full_view.dart:6:8 • unused_import
			`,
			"error",
			true,
		},
		{
			"Given info severity and info message that not contains the protected word info when the output is validated then expect to be fail",
			createFailMessage(infoLevel, ""),
			infoLevel,
			true,
		},
		{
			"Given info severity and info message that contains the protected word info when the output is validated then expect to be fail",
			createFailMessage(infoLevel, infoLevel),
			infoLevel,
			true,
		},
		{
			"Given info severity and warning message that not contains the protected word info when the output is validated then expect to be fail",
			createFailMessage(warningLevel, ""),
			infoLevel,
			true,
		},
		{
			"Given info severity and warning message that contains the protected word info when the output is validated then expect to be fail",
			createFailMessage(warningLevel, infoLevel),
			infoLevel,
			true,
		},
		{
			"Given info severity and error message that not contains the protected word info when the output is validated then expect to be fail",
			createFailMessage(errorLevel, ""),
			infoLevel,
			true,
		},
		{
			"Given info severity and error message that contains the protected word info when the output is validated then expect to be fail",
			createFailMessage(errorLevel, infoLevel),
			infoLevel,
			true,
		},
		// Warning severity
		{
			"Given warning severity and info message that not contains the protected word warning when the output is validated then expect not to be fail",
			createFailMessage(infoLevel, ""),
			warningLevel,
			false,
		},
		{
			"Given warning severity and info message that contains the protected word warning when the output is validated then expect not to be fail",
			createFailMessage(infoLevel, warningLevel),
			warningLevel,
			false,
		},
		{
			"Given warning severity and warning message that not contains the protected word warning when the output is validated then expect to be fail",
			createFailMessage(warningLevel, ""),
			warningLevel,
			true,
		},
		{
			"Given warning severity and warning message that contains the protected word warning when the output is validated then expect to be fail",
			createFailMessage(warningLevel, warningLevel),
			warningLevel,
			true,
		},
		{
			"Given warning severity and error message that not contains the protected word warning when the output is validated then expect to be fail",
			createFailMessage(errorLevel, ""),
			warningLevel,
			true,
		},
		{
			"Given warning severity and error message that contains the protected word warning when the output is validated then expect to be fail",
			createFailMessage(errorLevel, warningLevel),
			warningLevel,
			true,
		},
		// Error severity
		{
			"Given error severity and info message that not contains the protected word error when the output is validated then expect not to be fail",
			createFailMessage(infoLevel, ""),
			errorLevel,
			false,
		},
		{
			"Given error severity and info message that contains the protected word error when the output is validated then expect not to be fail",
			createFailMessage(infoLevel, errorLevel),
			errorLevel,
			false,
		},
		{
			"Given error severity and warning message that not contains the protected word error when the output is validated then expect not to be fail",
			createFailMessage(warningLevel, ""),
			errorLevel,
			false,
		},
		{
			"Given error severity and warning message that contains the protected word error when the output is validated then expect not to be fail",
			createFailMessage(warningLevel, errorLevel),
			errorLevel,
			false,
		},
		{
			"Given error severity and error message that not contains the protected word error when the output is validated then expect to be fail",
			createFailMessage(errorLevel, ""),
			errorLevel,
			true,
		},
		{
			"Given error severity and error message that contains the protected word error when the output is validated then expect to be fail",
			createFailMessage(errorLevel, errorLevel),
			errorLevel,
			true,
		},
	}

	for _, tt := range cases {
		if got := hasAnalyzeError(tt.cmdOutput, tt.failSeverity); got != tt.want {
			t.Errorf("%s: got %t want %t", tt.title, got, tt.want)
		}
	}
}

func createFailMessage(failSeverity, protectedWord string) string {
	dummyFileName := "some_file"
	if protectedWord != "" {
		dummyFileName = protectedWord
	}

	return fmt.Sprintf("%s • lib/%s.dart:1:1", failSeverity, dummyFileName)
}

func TestHasOtherError(t *testing.T) {
	cases := []struct {
		title     string
		cmdOutput string
		want      bool
	}{
		{
			"contains other error",
			`step failed
			failed
			`,
			true,
		},
		{
			"contains other error containing the protected word 'info'",
			`big bad info occured`,
			true,
		},
		{
			"contains other error containing the protected word 'warning'",
			`big bad warning occured`,
			true,
		},
		{
			"contains other error containing the protected word 'error'",
			`big bad error occured`,
			true,
		},
		{
			"contains error and info level violation",
			`info • Unused import: 'dart:math' • lib/package.dart:3:8 • unused_import
			error • Expected to find ';' • lib/package.dart:3:8 • expected_token
			`,
			false,
		},
	}

	for _, tt := range cases {
		if got := hasOtherError(tt.cmdOutput); got != tt.want {
			t.Errorf("%s: got %t want %t", tt.title, got, tt.want)
		}
	}
}
