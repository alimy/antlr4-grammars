// Code generated from PROV_N.g4 by ANTLR 4.7.2. DO NOT EDIT.

package prov_n

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 54, 699,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29,
	3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35,
	6, 35, 418, 10, 35, 13, 35, 14, 35, 419, 3, 35, 3, 35, 3, 36, 3, 36, 3,
	36, 3, 36, 7, 36, 428, 10, 36, 12, 36, 14, 36, 431, 11, 36, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 7, 37, 442, 10, 37, 12,
	37, 14, 37, 445, 11, 37, 3, 37, 3, 37, 3, 38, 3, 38, 7, 38, 451, 10, 38,
	12, 38, 14, 38, 454, 11, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40,
	3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 7, 43, 469, 10, 43, 12,
	43, 14, 43, 472, 11, 43, 3, 43, 5, 43, 475, 10, 43, 3, 44, 3, 44, 3, 45,
	3, 45, 5, 45, 481, 10, 45, 3, 46, 3, 46, 3, 46, 3, 46, 5, 46, 487, 10,
	46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 5, 49, 496, 10, 49,
	3, 49, 3, 49, 3, 49, 3, 49, 5, 49, 502, 10, 49, 3, 50, 3, 50, 3, 50, 5,
	50, 507, 10, 50, 3, 50, 3, 50, 3, 50, 7, 50, 512, 10, 50, 12, 50, 14, 50,
	515, 11, 50, 3, 50, 3, 50, 5, 50, 519, 10, 50, 5, 50, 521, 10, 50, 3, 51,
	3, 51, 3, 51, 5, 51, 526, 10, 51, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3,
	53, 3, 53, 3, 54, 3, 54, 5, 54, 537, 10, 54, 3, 55, 3, 55, 5, 55, 541,
	10, 55, 3, 56, 5, 56, 544, 10, 56, 3, 56, 6, 56, 547, 10, 56, 13, 56, 14,
	56, 548, 3, 57, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58, 3, 59, 3, 59,
	3, 59, 7, 59, 561, 10, 59, 12, 59, 14, 59, 564, 11, 59, 3, 59, 3, 59, 3,
	60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 5, 60, 575, 10, 60, 3, 60,
	3, 60, 5, 60, 579, 10, 60, 7, 60, 581, 10, 60, 12, 60, 14, 60, 584, 11,
	60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 61, 5, 61, 591, 10, 61, 3, 61, 3, 61,
	3, 61, 3, 61, 6, 61, 597, 10, 61, 13, 61, 14, 61, 598, 3, 61, 3, 61, 3,
	61, 3, 61, 5, 61, 605, 10, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 5, 61,
	612, 10, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 5, 61, 621,
	10, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 5, 61, 628, 10, 61, 3, 61, 3,
	61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 6, 61, 638, 10, 61, 13, 61,
	14, 61, 639, 5, 61, 642, 10, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3,
	61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 6, 61, 655, 10, 61, 13, 61, 14,
	61, 656, 5, 61, 659, 10, 61, 5, 61, 661, 10, 61, 3, 61, 3, 61, 3, 61, 3,
	61, 3, 61, 3, 61, 5, 61, 669, 10, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61,
	3, 61, 3, 61, 3, 61, 5, 61, 679, 10, 61, 5, 61, 681, 10, 61, 3, 62, 3,
	62, 6, 62, 685, 10, 62, 13, 62, 14, 62, 686, 3, 62, 3, 62, 6, 62, 691,
	10, 62, 13, 62, 14, 62, 692, 7, 62, 695, 10, 62, 12, 62, 14, 62, 698, 11,
	62, 3, 429, 2, 63, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10,
	19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19,
	37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28,
	55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37,
	73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 2, 87, 2, 89, 2, 91,
	2, 93, 2, 95, 44, 97, 45, 99, 2, 101, 2, 103, 2, 105, 2, 107, 46, 109,
	47, 111, 48, 113, 49, 115, 50, 117, 51, 119, 52, 121, 53, 123, 54, 3, 2,
	23, 5, 2, 11, 12, 14, 15, 34, 34, 4, 2, 12, 12, 15, 15, 10, 2, 2, 34, 36,
	36, 62, 62, 64, 64, 94, 94, 96, 96, 98, 98, 125, 127, 15, 2, 67, 92, 99,
	124, 194, 216, 218, 248, 250, 769, 882, 895, 897, 8193, 8206, 8207, 8306,
	8593, 11266, 12273, 12291, 55297, 63746, 64977, 65010, 65535, 5, 2, 185,
	185, 770, 881, 8257, 8258, 9, 2, 35, 35, 37, 38, 40, 40, 44, 45, 49, 49,
	65, 66, 128, 128, 8, 2, 41, 43, 46, 48, 60, 61, 63, 63, 93, 93, 95, 95,
	4, 2, 67, 72, 99, 104, 10, 2, 36, 36, 41, 41, 94, 94, 100, 100, 104, 104,
	112, 112, 116, 116, 118, 118, 6, 2, 12, 12, 15, 15, 36, 36, 94, 94, 4,
	2, 36, 36, 94, 94, 3, 2, 51, 59, 3, 2, 50, 59, 3, 2, 50, 52, 3, 2, 51,
	52, 3, 2, 50, 51, 3, 2, 50, 53, 3, 2, 50, 55, 4, 2, 45, 45, 47, 47, 4,
	2, 67, 92, 99, 124, 5, 2, 50, 59, 67, 92, 99, 124, 2, 740, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3,
	2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27,
	3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2,
	35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2,
	2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2,
	2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2,
	2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3,
	2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73,
	3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2,
	81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2,
	2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3,
	2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2,
	121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 3, 125, 3, 2, 2, 2, 5, 133, 3, 2,
	2, 2, 7, 140, 3, 2, 2, 2, 9, 147, 3, 2, 2, 2, 11, 149, 3, 2, 2, 2, 13,
	151, 3, 2, 2, 2, 15, 153, 3, 2, 2, 2, 17, 155, 3, 2, 2, 2, 19, 157, 3,
	2, 2, 2, 21, 159, 3, 2, 2, 2, 23, 162, 3, 2, 2, 2, 25, 171, 3, 2, 2, 2,
	27, 186, 3, 2, 2, 2, 29, 188, 3, 2, 2, 2, 31, 193, 3, 2, 2, 2, 33, 206,
	3, 2, 2, 2, 35, 217, 3, 2, 2, 2, 37, 234, 3, 2, 2, 2, 39, 248, 3, 2, 2,
	2, 41, 254, 3, 2, 2, 2, 43, 272, 3, 2, 2, 2, 45, 288, 3, 2, 2, 2, 47, 304,
	3, 2, 2, 2, 49, 319, 3, 2, 2, 2, 51, 335, 3, 2, 2, 2, 53, 347, 3, 2, 2,
	2, 55, 364, 3, 2, 2, 2, 57, 374, 3, 2, 2, 2, 59, 376, 3, 2, 2, 2, 61, 378,
	3, 2, 2, 2, 63, 387, 3, 2, 2, 2, 65, 399, 3, 2, 2, 2, 67, 406, 3, 2, 2,
	2, 69, 417, 3, 2, 2, 2, 71, 423, 3, 2, 2, 2, 73, 437, 3, 2, 2, 2, 75, 448,
	3, 2, 2, 2, 77, 457, 3, 2, 2, 2, 79, 459, 3, 2, 2, 2, 81, 461, 3, 2, 2,
	2, 83, 463, 3, 2, 2, 2, 85, 465, 3, 2, 2, 2, 87, 476, 3, 2, 2, 2, 89, 480,
	3, 2, 2, 2, 91, 486, 3, 2, 2, 2, 93, 488, 3, 2, 2, 2, 95, 490, 3, 2, 2,
	2, 97, 501, 3, 2, 2, 2, 99, 506, 3, 2, 2, 2, 101, 525, 3, 2, 2, 2, 103,
	527, 3, 2, 2, 2, 105, 530, 3, 2, 2, 2, 107, 536, 3, 2, 2, 2, 109, 540,
	3, 2, 2, 2, 111, 543, 3, 2, 2, 2, 113, 550, 3, 2, 2, 2, 115, 554, 3, 2,
	2, 2, 117, 557, 3, 2, 2, 2, 119, 567, 3, 2, 2, 2, 121, 590, 3, 2, 2, 2,
	123, 682, 3, 2, 2, 2, 125, 126, 7, 102, 2, 2, 126, 127, 7, 103, 2, 2, 127,
	128, 7, 104, 2, 2, 128, 129, 7, 99, 2, 2, 129, 130, 7, 119, 2, 2, 130,
	131, 7, 110, 2, 2, 131, 132, 7, 118, 2, 2, 132, 4, 3, 2, 2, 2, 133, 134,
	7, 114, 2, 2, 134, 135, 7, 116, 2, 2, 135, 136, 7, 103, 2, 2, 136, 137,
	7, 104, 2, 2, 137, 138, 7, 107, 2, 2, 138, 139, 7, 122, 2, 2, 139, 6, 3,
	2, 2, 2, 140, 141, 7, 103, 2, 2, 141, 142, 7, 112, 2, 2, 142, 143, 7, 118,
	2, 2, 143, 144, 7, 107, 2, 2, 144, 145, 7, 118, 2, 2, 145, 146, 7, 123,
	2, 2, 146, 8, 3, 2, 2, 2, 147, 148, 7, 42, 2, 2, 148, 10, 3, 2, 2, 2, 149,
	150, 7, 43, 2, 2, 150, 12, 3, 2, 2, 2, 151, 152, 7, 46, 2, 2, 152, 14,
	3, 2, 2, 2, 153, 154, 7, 93, 2, 2, 154, 16, 3, 2, 2, 2, 155, 156, 7, 95,
	2, 2, 156, 18, 3, 2, 2, 2, 157, 158, 7, 63, 2, 2, 158, 20, 3, 2, 2, 2,
	159, 160, 7, 39, 2, 2, 160, 161, 7, 39, 2, 2, 161, 22, 3, 2, 2, 2, 162,
	163, 7, 99, 2, 2, 163, 164, 7, 101, 2, 2, 164, 165, 7, 118, 2, 2, 165,
	166, 7, 107, 2, 2, 166, 167, 7, 120, 2, 2, 167, 168, 7, 107, 2, 2, 168,
	169, 7, 118, 2, 2, 169, 170, 7, 123, 2, 2, 170, 24, 3, 2, 2, 2, 171, 172,
	7, 121, 2, 2, 172, 173, 7, 99, 2, 2, 173, 174, 7, 117, 2, 2, 174, 175,
	7, 73, 2, 2, 175, 176, 7, 103, 2, 2, 176, 177, 7, 112, 2, 2, 177, 178,
	7, 103, 2, 2, 178, 179, 7, 116, 2, 2, 179, 180, 7, 99, 2, 2, 180, 181,
	7, 118, 2, 2, 181, 182, 7, 103, 2, 2, 182, 183, 7, 102, 2, 2, 183, 184,
	7, 68, 2, 2, 184, 185, 7, 123, 2, 2, 185, 26, 3, 2, 2, 2, 186, 187, 7,
	61, 2, 2, 187, 28, 3, 2, 2, 2, 188, 189, 7, 119, 2, 2, 189, 190, 7, 117,
	2, 2, 190, 191, 7, 103, 2, 2, 191, 192, 7, 102, 2, 2, 192, 30, 3, 2, 2,
	2, 193, 194, 7, 121, 2, 2, 194, 195, 7, 99, 2, 2, 195, 196, 7, 117, 2,
	2, 196, 197, 7, 85, 2, 2, 197, 198, 7, 118, 2, 2, 198, 199, 7, 99, 2, 2,
	199, 200, 7, 116, 2, 2, 200, 201, 7, 118, 2, 2, 201, 202, 7, 103, 2, 2,
	202, 203, 7, 102, 2, 2, 203, 204, 7, 68, 2, 2, 204, 205, 7, 123, 2, 2,
	205, 32, 3, 2, 2, 2, 206, 207, 7, 121, 2, 2, 207, 208, 7, 99, 2, 2, 208,
	209, 7, 117, 2, 2, 209, 210, 7, 71, 2, 2, 210, 211, 7, 112, 2, 2, 211,
	212, 7, 102, 2, 2, 212, 213, 7, 103, 2, 2, 213, 214, 7, 102, 2, 2, 214,
	215, 7, 68, 2, 2, 215, 216, 7, 123, 2, 2, 216, 34, 3, 2, 2, 2, 217, 218,
	7, 121, 2, 2, 218, 219, 7, 99, 2, 2, 219, 220, 7, 117, 2, 2, 220, 221,
	7, 75, 2, 2, 221, 222, 7, 112, 2, 2, 222, 223, 7, 120, 2, 2, 223, 224,
	7, 99, 2, 2, 224, 225, 7, 110, 2, 2, 225, 226, 7, 107, 2, 2, 226, 227,
	7, 102, 2, 2, 227, 228, 7, 99, 2, 2, 228, 229, 7, 118, 2, 2, 229, 230,
	7, 103, 2, 2, 230, 231, 7, 102, 2, 2, 231, 232, 7, 68, 2, 2, 232, 233,
	7, 123, 2, 2, 233, 36, 3, 2, 2, 2, 234, 235, 7, 121, 2, 2, 235, 236, 7,
	99, 2, 2, 236, 237, 7, 117, 2, 2, 237, 238, 7, 75, 2, 2, 238, 239, 7, 112,
	2, 2, 239, 240, 7, 104, 2, 2, 240, 241, 7, 113, 2, 2, 241, 242, 7, 116,
	2, 2, 242, 243, 7, 111, 2, 2, 243, 244, 7, 103, 2, 2, 244, 245, 7, 102,
	2, 2, 245, 246, 7, 68, 2, 2, 246, 247, 7, 123, 2, 2, 247, 38, 3, 2, 2,
	2, 248, 249, 7, 99, 2, 2, 249, 250, 7, 105, 2, 2, 250, 251, 7, 103, 2,
	2, 251, 252, 7, 112, 2, 2, 252, 253, 7, 118, 2, 2, 253, 40, 3, 2, 2, 2,
	254, 255, 7, 121, 2, 2, 255, 256, 7, 99, 2, 2, 256, 257, 7, 117, 2, 2,
	257, 258, 7, 67, 2, 2, 258, 259, 7, 117, 2, 2, 259, 260, 7, 117, 2, 2,
	260, 261, 7, 113, 2, 2, 261, 262, 7, 101, 2, 2, 262, 263, 7, 107, 2, 2,
	263, 264, 7, 99, 2, 2, 264, 265, 7, 118, 2, 2, 265, 266, 7, 103, 2, 2,
	266, 267, 7, 102, 2, 2, 267, 268, 7, 89, 2, 2, 268, 269, 7, 107, 2, 2,
	269, 270, 7, 118, 2, 2, 270, 271, 7, 106, 2, 2, 271, 42, 3, 2, 2, 2, 272,
	273, 7, 121, 2, 2, 273, 274, 7, 99, 2, 2, 274, 275, 7, 117, 2, 2, 275,
	276, 7, 67, 2, 2, 276, 277, 7, 118, 2, 2, 277, 278, 7, 118, 2, 2, 278,
	279, 7, 116, 2, 2, 279, 280, 7, 107, 2, 2, 280, 281, 7, 100, 2, 2, 281,
	282, 7, 119, 2, 2, 282, 283, 7, 118, 2, 2, 283, 284, 7, 103, 2, 2, 284,
	285, 7, 102, 2, 2, 285, 286, 7, 86, 2, 2, 286, 287, 7, 113, 2, 2, 287,
	44, 3, 2, 2, 2, 288, 289, 7, 99, 2, 2, 289, 290, 7, 101, 2, 2, 290, 291,
	7, 118, 2, 2, 291, 292, 7, 103, 2, 2, 292, 293, 7, 102, 2, 2, 293, 294,
	7, 81, 2, 2, 294, 295, 7, 112, 2, 2, 295, 296, 7, 68, 2, 2, 296, 297, 7,
	103, 2, 2, 297, 298, 7, 106, 2, 2, 298, 299, 7, 99, 2, 2, 299, 300, 7,
	110, 2, 2, 300, 301, 7, 104, 2, 2, 301, 302, 7, 81, 2, 2, 302, 303, 7,
	104, 2, 2, 303, 46, 3, 2, 2, 2, 304, 305, 7, 121, 2, 2, 305, 306, 7, 99,
	2, 2, 306, 307, 7, 117, 2, 2, 307, 308, 7, 70, 2, 2, 308, 309, 7, 103,
	2, 2, 309, 310, 7, 116, 2, 2, 310, 311, 7, 107, 2, 2, 311, 312, 7, 120,
	2, 2, 312, 313, 7, 103, 2, 2, 313, 314, 7, 102, 2, 2, 314, 315, 7, 72,
	2, 2, 315, 316, 7, 116, 2, 2, 316, 317, 7, 113, 2, 2, 317, 318, 7, 111,
	2, 2, 318, 48, 3, 2, 2, 2, 319, 320, 7, 121, 2, 2, 320, 321, 7, 99, 2,
	2, 321, 322, 7, 117, 2, 2, 322, 323, 7, 75, 2, 2, 323, 324, 7, 112, 2,
	2, 324, 325, 7, 104, 2, 2, 325, 326, 7, 110, 2, 2, 326, 327, 7, 119, 2,
	2, 327, 328, 7, 103, 2, 2, 328, 329, 7, 112, 2, 2, 329, 330, 7, 101, 2,
	2, 330, 331, 7, 103, 2, 2, 331, 332, 7, 102, 2, 2, 332, 333, 7, 68, 2,
	2, 333, 334, 7, 123, 2, 2, 334, 50, 3, 2, 2, 2, 335, 336, 7, 99, 2, 2,
	336, 337, 7, 110, 2, 2, 337, 338, 7, 118, 2, 2, 338, 339, 7, 103, 2, 2,
	339, 340, 7, 116, 2, 2, 340, 341, 7, 112, 2, 2, 341, 342, 7, 99, 2, 2,
	342, 343, 7, 118, 2, 2, 343, 344, 7, 103, 2, 2, 344, 345, 7, 81, 2, 2,
	345, 346, 7, 104, 2, 2, 346, 52, 3, 2, 2, 2, 347, 348, 7, 117, 2, 2, 348,
	349, 7, 114, 2, 2, 349, 350, 7, 103, 2, 2, 350, 351, 7, 101, 2, 2, 351,
	352, 7, 107, 2, 2, 352, 353, 7, 99, 2, 2, 353, 354, 7, 110, 2, 2, 354,
	355, 7, 107, 2, 2, 355, 356, 7, 124, 2, 2, 356, 357, 7, 99, 2, 2, 357,
	358, 7, 118, 2, 2, 358, 359, 7, 107, 2, 2, 359, 360, 7, 113, 2, 2, 360,
	361, 7, 112, 2, 2, 361, 362, 7, 81, 2, 2, 362, 363, 7, 104, 2, 2, 363,
	54, 3, 2, 2, 2, 364, 365, 7, 106, 2, 2, 365, 366, 7, 99, 2, 2, 366, 367,
	7, 102, 2, 2, 367, 368, 7, 79, 2, 2, 368, 369, 7, 103, 2, 2, 369, 370,
	7, 111, 2, 2, 370, 371, 7, 100, 2, 2, 371, 372, 7, 103, 2, 2, 372, 373,
	7, 116, 2, 2, 373, 56, 3, 2, 2, 2, 374, 375, 7, 125, 2, 2, 375, 58, 3,
	2, 2, 2, 376, 377, 7, 127, 2, 2, 377, 60, 3, 2, 2, 2, 378, 379, 7, 102,
	2, 2, 379, 380, 7, 113, 2, 2, 380, 381, 7, 101, 2, 2, 381, 382, 7, 119,
	2, 2, 382, 383, 7, 111, 2, 2, 383, 384, 7, 103, 2, 2, 384, 385, 7, 112,
	2, 2, 385, 386, 7, 118, 2, 2, 386, 62, 3, 2, 2, 2, 387, 388, 7, 103, 2,
	2, 388, 389, 7, 112, 2, 2, 389, 390, 7, 102, 2, 2, 390, 391, 7, 70, 2,
	2, 391, 392, 7, 113, 2, 2, 392, 393, 7, 101, 2, 2, 393, 394, 7, 119, 2,
	2, 394, 395, 7, 111, 2, 2, 395, 396, 7, 103, 2, 2, 396, 397, 7, 112, 2,
	2, 397, 398, 7, 118, 2, 2, 398, 64, 3, 2, 2, 2, 399, 400, 7, 100, 2, 2,
	400, 401, 7, 119, 2, 2, 401, 402, 7, 112, 2, 2, 402, 403, 7, 102, 2, 2,
	403, 404, 7, 110, 2, 2, 404, 405, 7, 103, 2, 2, 405, 66, 3, 2, 2, 2, 406,
	407, 7, 103, 2, 2, 407, 408, 7, 112, 2, 2, 408, 409, 7, 102, 2, 2, 409,
	410, 7, 68, 2, 2, 410, 411, 7, 119, 2, 2, 411, 412, 7, 112, 2, 2, 412,
	413, 7, 102, 2, 2, 413, 414, 7, 110, 2, 2, 414, 415, 7, 103, 2, 2, 415,
	68, 3, 2, 2, 2, 416, 418, 9, 2, 2, 2, 417, 416, 3, 2, 2, 2, 418, 419, 3,
	2, 2, 2, 419, 417, 3, 2, 2, 2, 419, 420, 3, 2, 2, 2, 420, 421, 3, 2, 2,
	2, 421, 422, 8, 35, 2, 2, 422, 70, 3, 2, 2, 2, 423, 424, 7, 49, 2, 2, 424,
	425, 7, 44, 2, 2, 425, 429, 3, 2, 2, 2, 426, 428, 11, 2, 2, 2, 427, 426,
	3, 2, 2, 2, 428, 431, 3, 2, 2, 2, 429, 430, 3, 2, 2, 2, 429, 427, 3, 2,
	2, 2, 430, 432, 3, 2, 2, 2, 431, 429, 3, 2, 2, 2, 432, 433, 7, 44, 2, 2,
	433, 434, 7, 49, 2, 2, 434, 435, 3, 2, 2, 2, 435, 436, 8, 36, 2, 2, 436,
	72, 3, 2, 2, 2, 437, 438, 7, 49, 2, 2, 438, 439, 7, 49, 2, 2, 439, 443,
	3, 2, 2, 2, 440, 442, 10, 3, 2, 2, 441, 440, 3, 2, 2, 2, 442, 445, 3, 2,
	2, 2, 443, 441, 3, 2, 2, 2, 443, 444, 3, 2, 2, 2, 444, 446, 3, 2, 2, 2,
	445, 443, 3, 2, 2, 2, 446, 447, 8, 37, 2, 2, 447, 74, 3, 2, 2, 2, 448,
	452, 5, 77, 39, 2, 449, 451, 10, 4, 2, 2, 450, 449, 3, 2, 2, 2, 451, 454,
	3, 2, 2, 2, 452, 450, 3, 2, 2, 2, 452, 453, 3, 2, 2, 2, 453, 455, 3, 2,
	2, 2, 454, 452, 3, 2, 2, 2, 455, 456, 5, 79, 40, 2, 456, 76, 3, 2, 2, 2,
	457, 458, 7, 62, 2, 2, 458, 78, 3, 2, 2, 2, 459, 460, 7, 64, 2, 2, 460,
	80, 3, 2, 2, 2, 461, 462, 7, 48, 2, 2, 462, 82, 3, 2, 2, 2, 463, 464, 7,
	47, 2, 2, 464, 84, 3, 2, 2, 2, 465, 474, 5, 87, 44, 2, 466, 469, 5, 91,
	46, 2, 467, 469, 5, 81, 41, 2, 468, 466, 3, 2, 2, 2, 468, 467, 3, 2, 2,
	2, 469, 472, 3, 2, 2, 2, 470, 468, 3, 2, 2, 2, 470, 471, 3, 2, 2, 2, 471,
	473, 3, 2, 2, 2, 472, 470, 3, 2, 2, 2, 473, 475, 5, 91, 46, 2, 474, 470,
	3, 2, 2, 2, 474, 475, 3, 2, 2, 2, 475, 86, 3, 2, 2, 2, 476, 477, 9, 5,
	2, 2, 477, 88, 3, 2, 2, 2, 478, 481, 5, 87, 44, 2, 479, 481, 7, 97, 2,
	2, 480, 478, 3, 2, 2, 2, 480, 479, 3, 2, 2, 2, 481, 90, 3, 2, 2, 2, 482,
	487, 5, 89, 45, 2, 483, 487, 5, 83, 42, 2, 484, 487, 5, 93, 47, 2, 485,
	487, 9, 6, 2, 2, 486, 482, 3, 2, 2, 2, 486, 483, 3, 2, 2, 2, 486, 484,
	3, 2, 2, 2, 486, 485, 3, 2, 2, 2, 487, 92, 3, 2, 2, 2, 488, 489, 4, 50,
	59, 2, 489, 94, 3, 2, 2, 2, 490, 491, 5, 85, 43, 2, 491, 96, 3, 2, 2, 2,
	492, 493, 5, 85, 43, 2, 493, 494, 7, 60, 2, 2, 494, 496, 3, 2, 2, 2, 495,
	492, 3, 2, 2, 2, 495, 496, 3, 2, 2, 2, 496, 497, 3, 2, 2, 2, 497, 502,
	5, 99, 50, 2, 498, 499, 5, 85, 43, 2, 499, 500, 7, 60, 2, 2, 500, 502,
	3, 2, 2, 2, 501, 495, 3, 2, 2, 2, 501, 498, 3, 2, 2, 2, 502, 98, 3, 2,
	2, 2, 503, 507, 5, 89, 45, 2, 504, 507, 5, 93, 47, 2, 505, 507, 5, 101,
	51, 2, 506, 503, 3, 2, 2, 2, 506, 504, 3, 2, 2, 2, 506, 505, 3, 2, 2, 2,
	507, 520, 3, 2, 2, 2, 508, 512, 5, 91, 46, 2, 509, 512, 5, 81, 41, 2, 510,
	512, 5, 101, 51, 2, 511, 508, 3, 2, 2, 2, 511, 509, 3, 2, 2, 2, 511, 510,
	3, 2, 2, 2, 512, 515, 3, 2, 2, 2, 513, 511, 3, 2, 2, 2, 513, 514, 3, 2,
	2, 2, 514, 518, 3, 2, 2, 2, 515, 513, 3, 2, 2, 2, 516, 519, 5, 91, 46,
	2, 517, 519, 5, 101, 51, 2, 518, 516, 3, 2, 2, 2, 518, 517, 3, 2, 2, 2,
	519, 521, 3, 2, 2, 2, 520, 513, 3, 2, 2, 2, 520, 521, 3, 2, 2, 2, 521,
	100, 3, 2, 2, 2, 522, 526, 9, 7, 2, 2, 523, 526, 5, 105, 53, 2, 524, 526,
	5, 103, 52, 2, 525, 522, 3, 2, 2, 2, 525, 523, 3, 2, 2, 2, 525, 524, 3,
	2, 2, 2, 526, 102, 3, 2, 2, 2, 527, 528, 7, 94, 2, 2, 528, 529, 9, 8, 2,
	2, 529, 104, 3, 2, 2, 2, 530, 531, 7, 39, 2, 2, 531, 532, 5, 107, 54, 2,
	532, 533, 5, 107, 54, 2, 533, 106, 3, 2, 2, 2, 534, 537, 5, 93, 47, 2,
	535, 537, 9, 9, 2, 2, 536, 534, 3, 2, 2, 2, 536, 535, 3, 2, 2, 2, 537,
	108, 3, 2, 2, 2, 538, 541, 5, 117, 59, 2, 539, 541, 5, 119, 60, 2, 540,
	538, 3, 2, 2, 2, 540, 539, 3, 2, 2, 2, 541, 110, 3, 2, 2, 2, 542, 544,
	7, 47, 2, 2, 543, 542, 3, 2, 2, 2, 543, 544, 3, 2, 2, 2, 544, 546, 3, 2,
	2, 2, 545, 547, 5, 93, 47, 2, 546, 545, 3, 2, 2, 2, 547, 548, 3, 2, 2,
	2, 548, 546, 3, 2, 2, 2, 548, 549, 3, 2, 2, 2, 549, 112, 3, 2, 2, 2, 550,
	551, 7, 41, 2, 2, 551, 552, 5, 97, 49, 2, 552, 553, 7, 41, 2, 2, 553, 114,
	3, 2, 2, 2, 554, 555, 7, 94, 2, 2, 555, 556, 9, 10, 2, 2, 556, 116, 3,
	2, 2, 2, 557, 562, 7, 36, 2, 2, 558, 561, 10, 11, 2, 2, 559, 561, 5, 115,
	58, 2, 560, 558, 3, 2, 2, 2, 560, 559, 3, 2, 2, 2, 561, 564, 3, 2, 2, 2,
	562, 560, 3, 2, 2, 2, 562, 563, 3, 2, 2, 2, 563, 565, 3, 2, 2, 2, 564,
	562, 3, 2, 2, 2, 565, 566, 7, 36, 2, 2, 566, 118, 3, 2, 2, 2, 567, 568,
	7, 36, 2, 2, 568, 569, 7, 36, 2, 2, 569, 570, 7, 36, 2, 2, 570, 582, 3,
	2, 2, 2, 571, 575, 7, 36, 2, 2, 572, 573, 7, 36, 2, 2, 573, 575, 7, 36,
	2, 2, 574, 571, 3, 2, 2, 2, 574, 572, 3, 2, 2, 2, 574, 575, 3, 2, 2, 2,
	575, 578, 3, 2, 2, 2, 576, 579, 10, 12, 2, 2, 577, 579, 5, 115, 58, 2,
	578, 576, 3, 2, 2, 2, 578, 577, 3, 2, 2, 2, 579, 581, 3, 2, 2, 2, 580,
	574, 3, 2, 2, 2, 581, 584, 3, 2, 2, 2, 582, 580, 3, 2, 2, 2, 582, 583,
	3, 2, 2, 2, 583, 585, 3, 2, 2, 2, 584, 582, 3, 2, 2, 2, 585, 586, 7, 36,
	2, 2, 586, 587, 7, 36, 2, 2, 587, 588, 7, 36, 2, 2, 588, 120, 3, 2, 2,
	2, 589, 591, 7, 47, 2, 2, 590, 589, 3, 2, 2, 2, 590, 591, 3, 2, 2, 2, 591,
	604, 3, 2, 2, 2, 592, 593, 9, 13, 2, 2, 593, 594, 9, 14, 2, 2, 594, 596,
	9, 14, 2, 2, 595, 597, 9, 14, 2, 2, 596, 595, 3, 2, 2, 2, 597, 598, 3,
	2, 2, 2, 598, 596, 3, 2, 2, 2, 598, 599, 3, 2, 2, 2, 599, 605, 3, 2, 2,
	2, 600, 601, 7, 50, 2, 2, 601, 602, 9, 14, 2, 2, 602, 603, 9, 14, 2, 2,
	603, 605, 9, 14, 2, 2, 604, 592, 3, 2, 2, 2, 604, 600, 3, 2, 2, 2, 605,
	606, 3, 2, 2, 2, 606, 611, 7, 47, 2, 2, 607, 608, 7, 50, 2, 2, 608, 612,
	9, 13, 2, 2, 609, 610, 7, 51, 2, 2, 610, 612, 9, 15, 2, 2, 611, 607, 3,
	2, 2, 2, 611, 609, 3, 2, 2, 2, 612, 613, 3, 2, 2, 2, 613, 620, 7, 47, 2,
	2, 614, 615, 7, 50, 2, 2, 615, 621, 9, 13, 2, 2, 616, 617, 9, 16, 2, 2,
	617, 621, 9, 14, 2, 2, 618, 619, 7, 53, 2, 2, 619, 621, 9, 17, 2, 2, 620,
	614, 3, 2, 2, 2, 620, 616, 3, 2, 2, 2, 620, 618, 3, 2, 2, 2, 621, 622,
	3, 2, 2, 2, 622, 660, 7, 86, 2, 2, 623, 624, 9, 17, 2, 2, 624, 628, 9,
	14, 2, 2, 625, 626, 7, 52, 2, 2, 626, 628, 9, 18, 2, 2, 627, 623, 3, 2,
	2, 2, 627, 625, 3, 2, 2, 2, 628, 629, 3, 2, 2, 2, 629, 630, 7, 60, 2, 2,
	630, 631, 9, 19, 2, 2, 631, 632, 9, 14, 2, 2, 632, 633, 7, 60, 2, 2, 633,
	634, 9, 19, 2, 2, 634, 641, 9, 14, 2, 2, 635, 637, 7, 48, 2, 2, 636, 638,
	9, 14, 2, 2, 637, 636, 3, 2, 2, 2, 638, 639, 3, 2, 2, 2, 639, 637, 3, 2,
	2, 2, 639, 640, 3, 2, 2, 2, 640, 642, 3, 2, 2, 2, 641, 635, 3, 2, 2, 2,
	641, 642, 3, 2, 2, 2, 642, 661, 3, 2, 2, 2, 643, 644, 7, 52, 2, 2, 644,
	645, 7, 54, 2, 2, 645, 646, 7, 60, 2, 2, 646, 647, 7, 50, 2, 2, 647, 648,
	7, 50, 2, 2, 648, 649, 7, 60, 2, 2, 649, 650, 7, 50, 2, 2, 650, 651, 7,
	50, 2, 2, 651, 658, 3, 2, 2, 2, 652, 654, 7, 48, 2, 2, 653, 655, 7, 50,
	2, 2, 654, 653, 3, 2, 2, 2, 655, 656, 3, 2, 2, 2, 656, 654, 3, 2, 2, 2,
	656, 657, 3, 2, 2, 2, 657, 659, 3, 2, 2, 2, 658, 652, 3, 2, 2, 2, 658,
	659, 3, 2, 2, 2, 659, 661, 3, 2, 2, 2, 660, 627, 3, 2, 2, 2, 660, 643,
	3, 2, 2, 2, 661, 680, 3, 2, 2, 2, 662, 681, 7, 92, 2, 2, 663, 678, 9, 20,
	2, 2, 664, 665, 7, 50, 2, 2, 665, 669, 9, 14, 2, 2, 666, 667, 7, 51, 2,
	2, 667, 669, 9, 18, 2, 2, 668, 664, 3, 2, 2, 2, 668, 666, 3, 2, 2, 2, 669,
	670, 3, 2, 2, 2, 670, 671, 7, 60, 2, 2, 671, 672, 9, 19, 2, 2, 672, 679,
	9, 14, 2, 2, 673, 674, 7, 51, 2, 2, 674, 675, 7, 54, 2, 2, 675, 676, 7,
	60, 2, 2, 676, 677, 7, 50, 2, 2, 677, 679, 7, 50, 2, 2, 678, 668, 3, 2,
	2, 2, 678, 673, 3, 2, 2, 2, 679, 681, 3, 2, 2, 2, 680, 662, 3, 2, 2, 2,
	680, 663, 3, 2, 2, 2, 680, 681, 3, 2, 2, 2, 681, 122, 3, 2, 2, 2, 682,
	684, 7, 66, 2, 2, 683, 685, 9, 21, 2, 2, 684, 683, 3, 2, 2, 2, 685, 686,
	3, 2, 2, 2, 686, 684, 3, 2, 2, 2, 686, 687, 3, 2, 2, 2, 687, 696, 3, 2,
	2, 2, 688, 690, 7, 47, 2, 2, 689, 691, 9, 22, 2, 2, 690, 689, 3, 2, 2,
	2, 691, 692, 3, 2, 2, 2, 692, 690, 3, 2, 2, 2, 692, 693, 3, 2, 2, 2, 693,
	695, 3, 2, 2, 2, 694, 688, 3, 2, 2, 2, 695, 698, 3, 2, 2, 2, 696, 694,
	3, 2, 2, 2, 696, 697, 3, 2, 2, 2, 697, 124, 3, 2, 2, 2, 698, 696, 3, 2,
	2, 2, 46, 2, 419, 429, 443, 452, 468, 470, 474, 480, 486, 495, 501, 506,
	511, 513, 518, 520, 525, 536, 540, 543, 548, 560, 562, 574, 578, 582, 590,
	598, 604, 611, 620, 627, 639, 641, 656, 658, 660, 668, 678, 680, 686, 692,
	696, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'default'", "'prefix'", "'entity'", "'('", "')'", "','", "'['", "']'",
	"'='", "'%%'", "'activity'", "'wasGeneratedBy'", "';'", "'used'", "'wasStartedBy'",
	"'wasEndedBy'", "'wasInvalidatedBy'", "'wasInformedBy'", "'agent'", "'wasAssociatedWith'",
	"'wasAttributedTo'", "'actedOnBehalfOf'", "'wasDerivedFrom'", "'wasInfluencedBy'",
	"'alternateOf'", "'specializationOf'", "'hadMember'", "'{'", "'}'", "'document'",
	"'endDocument'", "'bundle'", "'endBundle'", "", "", "", "", "'<'", "'>'",
	"'.'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "DOCUMENT", "ENDDOCUMENT",
	"BUNDLE", "ENDBUNDLE", "WS", "COMMENT", "LINE_COMMENT", "IRI_REF", "LESS",
	"GREATER", "DOT", "MINUS", "PREFX", "QUALIFIED_NAME", "HEX", "STRING_LITERAL",
	"INT_LITERAL", "QUALIFIED_NAME_LITERAL", "ECHAR", "STRING_LITERAL2", "STRING_LITERAL_LONG2",
	"DATETIME", "LANGTAG",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "DOCUMENT", "ENDDOCUMENT", "BUNDLE",
	"ENDBUNDLE", "WS", "COMMENT", "LINE_COMMENT", "IRI_REF", "LESS", "GREATER",
	"DOT", "MINUS", "PN_PREFIX", "PN_CHARS_BASE", "PN_CHARS_U", "PN_CHARS",
	"DIGIT", "PREFX", "QUALIFIED_NAME", "PN_LOCAL", "PN_CHARS_OTHERS", "PN_CHARS_ESC",
	"PERCENT", "HEX", "STRING_LITERAL", "INT_LITERAL", "QUALIFIED_NAME_LITERAL",
	"ECHAR", "STRING_LITERAL2", "STRING_LITERAL_LONG2", "DATETIME", "LANGTAG",
}

type PROV_NLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewPROV_NLexer(input antlr.CharStream) *PROV_NLexer {

	l := new(PROV_NLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "PROV_N.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PROV_NLexer tokens.
const (
	PROV_NLexerT__0                   = 1
	PROV_NLexerT__1                   = 2
	PROV_NLexerT__2                   = 3
	PROV_NLexerT__3                   = 4
	PROV_NLexerT__4                   = 5
	PROV_NLexerT__5                   = 6
	PROV_NLexerT__6                   = 7
	PROV_NLexerT__7                   = 8
	PROV_NLexerT__8                   = 9
	PROV_NLexerT__9                   = 10
	PROV_NLexerT__10                  = 11
	PROV_NLexerT__11                  = 12
	PROV_NLexerT__12                  = 13
	PROV_NLexerT__13                  = 14
	PROV_NLexerT__14                  = 15
	PROV_NLexerT__15                  = 16
	PROV_NLexerT__16                  = 17
	PROV_NLexerT__17                  = 18
	PROV_NLexerT__18                  = 19
	PROV_NLexerT__19                  = 20
	PROV_NLexerT__20                  = 21
	PROV_NLexerT__21                  = 22
	PROV_NLexerT__22                  = 23
	PROV_NLexerT__23                  = 24
	PROV_NLexerT__24                  = 25
	PROV_NLexerT__25                  = 26
	PROV_NLexerT__26                  = 27
	PROV_NLexerT__27                  = 28
	PROV_NLexerT__28                  = 29
	PROV_NLexerDOCUMENT               = 30
	PROV_NLexerENDDOCUMENT            = 31
	PROV_NLexerBUNDLE                 = 32
	PROV_NLexerENDBUNDLE              = 33
	PROV_NLexerWS                     = 34
	PROV_NLexerCOMMENT                = 35
	PROV_NLexerLINE_COMMENT           = 36
	PROV_NLexerIRI_REF                = 37
	PROV_NLexerLESS                   = 38
	PROV_NLexerGREATER                = 39
	PROV_NLexerDOT                    = 40
	PROV_NLexerMINUS                  = 41
	PROV_NLexerPREFX                  = 42
	PROV_NLexerQUALIFIED_NAME         = 43
	PROV_NLexerHEX                    = 44
	PROV_NLexerSTRING_LITERAL         = 45
	PROV_NLexerINT_LITERAL            = 46
	PROV_NLexerQUALIFIED_NAME_LITERAL = 47
	PROV_NLexerECHAR                  = 48
	PROV_NLexerSTRING_LITERAL2        = 49
	PROV_NLexerSTRING_LITERAL_LONG2   = 50
	PROV_NLexerDATETIME               = 51
	PROV_NLexerLANGTAG                = 52
)