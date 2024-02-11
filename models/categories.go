package models

type ExpenseCategory uint

const (
	UNKNOWN_EXPENSE ExpenseCategory = iota
	RENT
	HOME
	GROCERIES
	FOOD
	SPLIT
	TRAVEL
	LUXURY
	TELECOM
	OUTING
)

var ExpenseNameToCategoryMap = map[string][]ExpenseCategory{
	"TEA":                     {FOOD},
	"PARLE":                   {LUXURY},
	"COOK":                    {SPLIT, HOME},
	"PETROL":                  {TRAVEL},
	"SUPRA":                   {TRAVEL},
	"FUEL":                    {TRAVEL},
	"BIKE":                    {TRAVEL},
	"MSHRUTHI":                {LUXURY},
	"PURE O NATURAL":          {GROCERIES},
	"BBINSTANT":               {GROCERIES, SPLIT},
	"EATSURE":                 {FOOD},
	"GPAYBILLPAY.RCHRG":       {TELECOM},
	"LUNCH":                   {FOOD},
	"FOOD":                    {FOOD},
	"DINNER":                  {FOOD},
	"TAARA KITCHEN":           {FOOD},
	"AUTO":                    {TRAVEL},
	"SWIGGY":                  {FOOD},
	"INSTAMART":               {GROCERIES, SPLIT},
	"WIFI":                    {HOME, SPLIT},
	"ELEC":                    {HOME, SPLIT},
	"VIJETHA":                 {GROCERIES, SPLIT},
	"RATNADEEP":               {GROCERIES, SPLIT},
	"YOUTUBE":                 {LUXURY, SPLIT},
	"RAPIDO":                  {TRAVEL},
	"GOOGLE PLAY SER CYBS S ": {LUXURY},
	"X2":                      {SPLIT},
	"X3":                      {SPLIT},
	"X4":                      {SPLIT},
	"SPLIT":                   {SPLIT},
	"POS":                     {OUTING},
	"OUTING":                  {OUTING},
	"IBIBO":                   {TRAVEL},
	"BUS":                     {TRAVEL},
	"PUSHPAK":                 {TRAVEL},
	"TO HOME":                 {TRAVEL},
	"BACK HOME":               {TRAVEL},
}
