// Copyright (c) 2020 Sam Blenny
// SPDX-License-Identifier: Apache-2.0 OR MIT
//
package font

// Holds description of sprite sheet and character map for generating a font
type FontSpec struct {
	Name    string // Name of font
	Sprites string // Which file holds the sprite sheet image with the grid of glyphs?
	Size    int    // How many pixels on a side is each glyph (precondition: square glyphs)
	Cols    int    // How many glyphs wide is the grid?
	Gutter  int    // How many px between glyphs?
	Border  int    // How many px wide are top and left borders?
	Legal   string // What credits or license notices need to be included in font file comments?
	RustOut string // Where should the generated source code go?
}

// Holds mappings from extended grapheme clusters to sprite sheet glyph grid coordinates
type CharSpec struct {
	HexCluster string
	Row        int
	Col        int
}

// Holds specification for a Unicode block (codepoint bounds and identifying string)
type UBlock struct {
	Low  uint32
	High uint32
	Name string
}

// Holds an alias for a hex-codepoint grapheme cluster in the primary index
type GCAlias struct {
	CanonHex string // Cannonical form in the index (has a CharSpec)
	AliasHex string // This one should map to same glyph as CanonHex
}

// Holds a matrix of pixel values
type Matrix [][]int

// Holds one row of pixel values from a matrix
type MatrixRow []int

// Holds packed XOR mask values of a blit pattern for character's glyph.
//
// Header: .Words[0] = ((width of blit pattern in px uint8) << 16)
//                     | (height of blit pattern in px as uint8) << 8)
//                     | (y-offset down from top of pattern as uint8)
// Pixel order: row-major order traversal of px matrix; top-left pixel goes in
//              least significant bit of of .Words[1]
// Mask values: bit=1 means foreground, bit=0 means background
//
// Patterns that need padding because their size is not a multiple of 32 bits
// (width*height % 32 != 0) get padded with zeros in the least significant bits
// of the last word.
type BlitPattern struct {
	Words []uint32
	CS    CharSpec
}
