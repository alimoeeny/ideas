package ideas

var UNIT_UNITLESS = &Unit{ID: "1", Name: "Unitless", Short1: ""}

var UNIT_RATIO = &Unit{ID: "10", Name: "Ratio", Short1: "ratio"}
var UNIT_PERCENTAGE = &Unit{ID: "20", Name: "Percentage", Short1: "%"}

var UNIT_IDENTIFIER = &Unit{ID: "110", Name: "identifier", Short1: ""}

var UNIT_COUNT = &Unit{ID: "12300", Name: "count", Short1: ""}
var UNIT_CATEGORICAL = &Unit{ID: "12350", Name: "categorical", Short1: "cat"}
var UNIT_POSITIVE_NEGATIVE_MISSING = &Unit{ID: "12360", Name: "Positive Negative Missing", Short1: "PNM"}

var UNIT_COUNT_PER_LITER = &Unit{ID: "12400", Name: "counts per liter", Short1: "/l"}
var UNIT_COUNT_MILLION_PER_LITER = &Unit{ID: "12500", Name: "counts million per liter", Short1: "/l"}   // 10^6 / l
var UNIT_COUNT_BILLION_PER_LITER = &Unit{ID: "12501", Name: "counts billion per liter", Short1: "/l"}   // 10^9 / l
var UNIT_COUNT_TRILLION_PER_LITER = &Unit{ID: "12502", Name: "counts trillion per liter", Short1: "/l"} // 10^12 / l

var UNIT_COUNTS_PER_MILLILITER = &Unit{ID: "12600", Name: "counts per milliliter", Short1: "/ml"}

var UNIT_COUNT_PER_UL = &Unit{ID: "125", Name: "count per microliter", Short1: "/µl"}
var UNIT_COUNT_MILLION_PER_UL = &Unit{ID: "125", Name: "count per microliter", Short1: "/µl"}

var UNIT_MILES = &Unit{ID: "1013", Name: "miles", Short1: "miles"}

var UNIT_MILLIMOLES = &Unit{ID: "11014", Name: "millimoles", Short1: "mmol"}

var UNIT_GRAMS_PER_LITER = &Unit{ID: "12015", Name: "grams per liter", Short1: "g/l"}

var UNIT_LITER = &Unit{ID: "13016", Name: "liters", Short1: "l"}
var UNIT_DECILITER = &Unit{ID: "13017", Name: "deciliter", Short1: "dl"}
var UNIT_CENTILITER = &Unit{ID: "13018", Name: "centiliter", Short1: "cl"}
var UNIT_MILLILITER = &Unit{ID: "13019", Name: "milliliter", Short1: "ml"}
var UNIT_MICROLITER = &Unit{ID: "13020", Name: "microliter", Short1: "µl"}
var UNIT_NANOLITER = &Unit{ID: "13021", Name: "nanoliter", Short1: "nl"}
var UNIT_PICOLITER = &Unit{ID: "13022", Name: "picoliter", Short1: "pl"}
var UNIT_FEMTOLITER = &Unit{ID: "13023", Name: "femtoliter", Short1: "fl"} // 1e-15 liter

var UNIT_GRAM = &Unit{ID: "14025", Name: "gram", Short1: "g"}
var UNIT_KILOGRAM = &Unit{ID: "14026", Name: "kilogram", Short1: "kg"}
var UNIT_MILLIGRAM = &Unit{ID: "14027", Name: "milligram", Short1: "mg"}
var UNIT_MICROGRAM = &Unit{ID: "14028", Name: "microgram", Short1: "µg"}
var UNIT_NANOGRAM = &Unit{ID: "14029", Name: "nanogram", Short1: "ng"}
var UNIT_PICOGRAM = &Unit{ID: "14030", Name: "picogram", Short1: "pg"}   // 1e-12 gram
var UNIT_FEMTOGRAM = &Unit{ID: "14031", Name: "femtogram", Short1: "fg"} // 1e-15 gram

var UNIT_TIMESTAMP = &Unit{ID: "1651075578382886759", Name: "timestamp", Short1: "timestamp"}
