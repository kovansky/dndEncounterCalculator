/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

package enum

//ErrorType specifies possible error types
type ErrorType int

const (
	ErrorEasy   ErrorType = 1 // Easy (user error)
	ErrorMedium ErrorType = 2 // Medium (file, code error)
	ErrorHard   ErrorType = 3 // Critical (code/system/runtime error)
)
