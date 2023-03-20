package type1c

import "github.com/benoitkugler/textlayout/fonts/simpleencodings"

// the Standard encoding the same as in PDF
// but expertEncoding is not the same as MacExpert
var expertEncoding = simpleencodings.Encoding{
	32:  "space",
	33:  "exclamsmall",
	34:  "Hungarumlautsmall",
	36:  "dollaroldstyle",
	37:  "dollarsuperior",
	38:  "ampersandsmall",
	39:  "Acutesmall",
	40:  "parenleftsuperior",
	41:  "parenrightsuperior",
	42:  "twodotenleader",
	43:  "onedotenleader",
	44:  "comma",
	45:  "hyphen",
	46:  "period",
	47:  "fraction",
	48:  "zerooldstyle",
	49:  "oneoldstyle",
	50:  "twooldstyle",
	51:  "threeoldstyle",
	52:  "fouroldstyle",
	53:  "fiveoldstyle",
	54:  "sixoldstyle",
	55:  "sevenoldstyle",
	56:  "eightoldstyle",
	57:  "nineoldstyle",
	58:  "colon",
	59:  "semicolon",
	60:  "commasuperior",
	61:  "threequartersemdash",
	62:  "periodsuperior",
	63:  "questionsmall",
	65:  "asuperior",
	66:  "bsuperior",
	67:  "centsuperior",
	68:  "dsuperior",
	69:  "esuperior",
	73:  "isuperior",
	76:  "lsuperior",
	77:  "msuperior",
	78:  "nsuperior",
	79:  "osuperior",
	82:  "rsuperior",
	83:  "ssuperior",
	84:  "tsuperior",
	86:  "ff",
	87:  "fi",
	88:  "fl",
	89:  "ffi",
	90:  "ffl",
	91:  "parenleftinferior",
	93:  "parenrightinferior",
	94:  "Circumflexsmall",
	95:  "hyphensuperior",
	96:  "Gravesmall",
	97:  "Asmall",
	98:  "Bsmall",
	99:  "Csmall",
	100: "Dsmall",
	101: "Esmall",
	102: "Fsmall",
	103: "Gsmall",
	104: "Hsmall",
	105: "Ismall",
	106: "Jsmall",
	107: "Ksmall",
	108: "Lsmall",
	109: "Msmall",
	110: "Nsmall",
	111: "Osmall",
	112: "Psmall",
	113: "Qsmall",
	114: "Rsmall",
	115: "Ssmall",
	116: "Tsmall",
	117: "Usmall",
	118: "Vsmall",
	119: "Wsmall",
	120: "Xsmall",
	121: "Ysmall",
	122: "Zsmall",
	123: "colonmonetary",
	124: "onefitted",
	125: "rupiah",
	126: "Tildesmall",
	161: "exclamdownsmall",
	162: "centoldstyle",
	163: "Lslashsmall",
	166: "Scaronsmall",
	167: "Zcaronsmall",
	168: "Dieresissmall",
	169: "Brevesmall",
	170: "Caronsmall",
	172: "Dotaccentsmall",
	175: "Macronsmall",
	178: "figuredash",
	179: "hypheninferior",
	182: "Ogoneksmall",
	183: "Ringsmall",
	184: "Cedillasmall",
	188: "onequarter",
	189: "onehalf",
	190: "threequarters",
	191: "questiondownsmall",
	192: "oneeighth",
	193: "threeeighths",
	194: "fiveeighths",
	195: "seveneighths",
	196: "onethird",
	197: "twothirds",
	200: "zerosuperior",
	201: "onesuperior",
	202: "twosuperior",
	203: "threesuperior",
	204: "foursuperior",
	205: "fivesuperior",
	206: "sixsuperior",
	207: "sevensuperior",
	208: "eightsuperior",
	209: "ninesuperior",
	210: "zeroinferior",
	211: "oneinferior",
	212: "twoinferior",
	213: "threeinferior",
	214: "fourinferior",
	215: "fiveinferior",
	216: "sixinferior",
	217: "seveninferior",
	218: "eightinferior",
	219: "nineinferior",
	220: "centinferior",
	221: "dollarinferior",
	222: "periodinferior",
	223: "commainferior",
	224: "Agravesmall",
	225: "Aacutesmall",
	226: "Acircumflexsmall",
	227: "Atildesmall",
	228: "Adieresissmall",
	229: "Aringsmall",
	230: "AEsmall",
	231: "Ccedillasmall",
	232: "Egravesmall",
	233: "Eacutesmall",
	234: "Ecircumflexsmall",
	235: "Edieresissmall",
	236: "Igravesmall",
	237: "Iacutesmall",
	238: "Icircumflexsmall",
	239: "Idieresissmall",
	240: "Ethsmall",
	241: "Ntildesmall",
	242: "Ogravesmall",
	243: "Oacutesmall",
	244: "Ocircumflexsmall",
	245: "Otildesmall",
	246: "Odieresissmall",
	247: "OEsmall",
	248: "Oslashsmall",
	249: "Ugravesmall",
	250: "Uacutesmall",
	251: "Ucircumflexsmall",
	252: "Udieresissmall",
	253: "Yacutesmall",
	254: "Thornsmall",
	255: "Ydieresissmall",
}

var stdStrings = [391]string{
	".notdef",
	"space",
	"exclam",
	"quotedbl",
	"numbersign",
	"dollar",
	"percent",
	"ampersand",
	"quoteright",
	"parenleft",
	"parenright",
	"asterisk",
	"plus",
	"comma",
	"hyphen",
	"period",
	"slash",
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"colon",
	"semicolon",
	"less",
	"equal",
	"greater",
	"question",
	"at",
	"A",
	"B",
	"C",
	"D",
	"E",
	"F",
	"G",
	"H",
	"I",
	"J",
	"K",
	"L",
	"M",
	"N",
	"O",
	"P",
	"Q",
	"R",
	"S",
	"T",
	"U",
	"V",
	"W",
	"X",
	"Y",
	"Z",
	"bracketleft",
	"backslash",
	"bracketright",
	"asciicircum",
	"underscore",
	"quoteleft",
	"a",
	"b",
	"c",
	"d",
	"e",
	"f",
	"g",
	"h",
	"i",
	"j",
	"k",
	"l",
	"m",
	"n",
	"o",
	"p",
	"q",
	"r",
	"s",
	"t",
	"u",
	"v",
	"w",
	"x",
	"y",
	"z",
	"braceleft",
	"bar",
	"braceright",
	"asciitilde",
	"exclamdown",
	"cent",
	"sterling",
	"fraction",
	"yen",
	"florin",
	"section",
	"currency",
	"quotesingle",
	"quotedblleft",
	"guillemotleft",
	"guilsinglleft",
	"guilsinglright",
	"fi",
	"fl",
	"endash",
	"dagger",
	"daggerdbl",
	"periodcentered",
	"paragraph",
	"bullet",
	"quotesinglbase",
	"quotedblbase",
	"quotedblright",
	"guillemotright",
	"ellipsis",
	"perthousand",
	"questiondown",
	"grave",
	"acute",
	"circumflex",
	"tilde",
	"macron",
	"breve",
	"dotaccent",
	"dieresis",
	"ring",
	"cedilla",
	"hungarumlaut",
	"ogonek",
	"caron",
	"emdash",
	"AE",
	"ordfeminine",
	"Lslash",
	"Oslash",
	"OE",
	"ordmasculine",
	"ae",
	"dotlessi",
	"lslash",
	"oslash",
	"oe",
	"germandbls",
	"onesuperior",
	"logicalnot",
	"mu",
	"trademark",
	"Eth",
	"onehalf",
	"plusminus",
	"Thorn",
	"onequarter",
	"divide",
	"brokenbar",
	"degree",
	"thorn",
	"threequarters",
	"twosuperior",
	"registered",
	"minus",
	"eth",
	"multiply",
	"threesuperior",
	"copyright",
	"Aacute",
	"Acircumflex",
	"Adieresis",
	"Agrave",
	"Aring",
	"Atilde",
	"Ccedilla",
	"Eacute",
	"Ecircumflex",
	"Edieresis",
	"Egrave",
	"Iacute",
	"Icircumflex",
	"Idieresis",
	"Igrave",
	"Ntilde",
	"Oacute",
	"Ocircumflex",
	"Odieresis",
	"Ograve",
	"Otilde",
	"Scaron",
	"Uacute",
	"Ucircumflex",
	"Udieresis",
	"Ugrave",
	"Yacute",
	"Ydieresis",
	"Zcaron",
	"aacute",
	"acircumflex",
	"adieresis",
	"agrave",
	"aring",
	"atilde",
	"ccedilla",
	"eacute",
	"ecircumflex",
	"edieresis",
	"egrave",
	"iacute",
	"icircumflex",
	"idieresis",
	"igrave",
	"ntilde",
	"oacute",
	"ocircumflex",
	"odieresis",
	"ograve",
	"otilde",
	"scaron",
	"uacute",
	"ucircumflex",
	"udieresis",
	"ugrave",
	"yacute",
	"ydieresis",
	"zcaron",
	"exclamsmall",
	"Hungarumlautsmall",
	"dollaroldstyle",
	"dollarsuperior",
	"ampersandsmall",
	"Acutesmall",
	"parenleftsuperior",
	"parenrightsuperior",
	"twodotenleader",
	"onedotenleader",
	"zerooldstyle",
	"oneoldstyle",
	"twooldstyle",
	"threeoldstyle",
	"fouroldstyle",
	"fiveoldstyle",
	"sixoldstyle",
	"sevenoldstyle",
	"eightoldstyle",
	"nineoldstyle",
	"commasuperior",
	"threequartersemdash",
	"periodsuperior",
	"questionsmall",
	"asuperior",
	"bsuperior",
	"centsuperior",
	"dsuperior",
	"esuperior",
	"isuperior",
	"lsuperior",
	"msuperior",
	"nsuperior",
	"osuperior",
	"rsuperior",
	"ssuperior",
	"tsuperior",
	"ff",
	"ffi",
	"ffl",
	"parenleftinferior",
	"parenrightinferior",
	"Circumflexsmall",
	"hyphensuperior",
	"Gravesmall",
	"Asmall",
	"Bsmall",
	"Csmall",
	"Dsmall",
	"Esmall",
	"Fsmall",
	"Gsmall",
	"Hsmall",
	"Ismall",
	"Jsmall",
	"Ksmall",
	"Lsmall",
	"Msmall",
	"Nsmall",
	"Osmall",
	"Psmall",
	"Qsmall",
	"Rsmall",
	"Ssmall",
	"Tsmall",
	"Usmall",
	"Vsmall",
	"Wsmall",
	"Xsmall",
	"Ysmall",
	"Zsmall",
	"colonmonetary",
	"onefitted",
	"rupiah",
	"Tildesmall",
	"exclamdownsmall",
	"centoldstyle",
	"Lslashsmall",
	"Scaronsmall",
	"Zcaronsmall",
	"Dieresissmall",
	"Brevesmall",
	"Caronsmall",
	"Dotaccentsmall",
	"Macronsmall",
	"figuredash",
	"hypheninferior",
	"Ogoneksmall",
	"Ringsmall",
	"Cedillasmall",
	"questiondownsmall",
	"oneeighth",
	"threeeighths",
	"fiveeighths",
	"seveneighths",
	"onethird",
	"twothirds",
	"zerosuperior",
	"foursuperior",
	"fivesuperior",
	"sixsuperior",
	"sevensuperior",
	"eightsuperior",
	"ninesuperior",
	"zeroinferior",
	"oneinferior",
	"twoinferior",
	"threeinferior",
	"fourinferior",
	"fiveinferior",
	"sixinferior",
	"seveninferior",
	"eightinferior",
	"nineinferior",
	"centinferior",
	"dollarinferior",
	"periodinferior",
	"commainferior",
	"Agravesmall",
	"Aacutesmall",
	"Acircumflexsmall",
	"Atildesmall",
	"Adieresissmall",
	"Aringsmall",
	"AEsmall",
	"Ccedillasmall",
	"Egravesmall",
	"Eacutesmall",
	"Ecircumflexsmall",
	"Edieresissmall",
	"Igravesmall",
	"Iacutesmall",
	"Icircumflexsmall",
	"Idieresissmall",
	"Ethsmall",
	"Ntildesmall",
	"Ogravesmall",
	"Oacutesmall",
	"Ocircumflexsmall",
	"Otildesmall",
	"Odieresissmall",
	"OEsmall",
	"Oslashsmall",
	"Ugravesmall",
	"Uacutesmall",
	"Ucircumflexsmall",
	"Udieresissmall",
	"Yacutesmall",
	"Thornsmall",
	"Ydieresissmall",
	"001.000",
	"001.001",
	"001.002",
	"001.003",
	"Black",
	"Bold",
	"Book",
	"Light",
	"Medium",
	"Regular",
	"Roman",
	"Semibold",
}
