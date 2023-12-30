package models

import (
	"fmt"
)

type (
	// Name is the name of the color.
	ColorName string

	// Color is the model for the color.
	Color struct {
		// Hex is the hex code of the color
		Hex string `json:"hex"`
	}
)

// String returns the string representation of the name.
func (c ColorName) String() string {
	return string(c)
}

const (
	ColorDefault           = ColorGray
	ColorRed     ColorName = "red"
	ColorOrange  ColorName = "orange"
	ColorYellow  ColorName = "yellow"
	ColorGreen   ColorName = "green"
	ColorBlue    ColorName = "blue"
	ColorIndigo  ColorName = "indigo"
	ColorViolet  ColorName = "violet"
	ColorGray    ColorName = "gray"
	ColorBrown   ColorName = "brown"
	ColorPink    ColorName = "pink"
	ColorPurple  ColorName = "purple"
)

var (
	Colors = map[ColorName]Color{
		ColorRed: {
			Hex: "#EF4444",
		},
		ColorOrange: {
			Hex: "#F97316",
		},
		ColorYellow: {
			Hex: "#FCD34D",
		},
		ColorGreen: {
			Hex: "#10B981",
		},
		ColorBlue: {
			Hex: "#3B82F6",
		},
		ColorIndigo: {
			Hex: "#6366F1",
		},
		ColorViolet: {
			Hex: "#8B5CF6",
		},
		ColorGray: {
			Hex: "#6B7280",
		},
		ColorBrown: {
			Hex: "#B45309",
		},
		ColorPink: {
			Hex: "#EC4899",
		},
		ColorPurple: {
			Hex: "#9333EA",
		},
	}
	ColorsName = func() (cN []string) {
		for colorName := range Colors {
			cN = append(cN, colorName.String())
		}
		return cN
	}
)

// GetColor returns the color based on the name.
// return ErrNotFound if the color name is not found.
func GetColor(name ColorName) (Color, error) {
	if ok, err := isValidColor(name); !ok {
		return Color{}, fmt.Errorf("GetColor: the color %s is %w (AllowedValues:%v)", name, err, ColorsName())
	}

	return Colors[name], nil
}

// GetDefaultColor returns the default color.
func GetDefaultColor() Color {
	return Colors[ColorDefault]
}

// IsValidColor checks if the color name is valid.
func IsValidColor(name ColorName) (isValid bool, err error) {
	return isValidColor(name)
}

// isValidColor checks if the color name is valid.
// return ErrNotFound if the color name is not found.
func isValidColor(name ColorName) (isValid bool, err error) {
	if _, ok := Colors[name]; !ok {
		return false, fmt.Errorf("%w color name %s", ErrNotFound, name.String())
	}

	return true, nil
}
