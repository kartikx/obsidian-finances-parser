## Description
This program attempts to convert HDFC bank statements into a format that my Obsidian budget planning workflow
can understand.

### HDFC Bank Statement Format

HDFC Bank Statements are in PDF format. Here's a sample:

```
 01/12/23      UPI-SWIGGY-SWIGGYSTORES@ICICI-ICIC0DC009     0000333506502743    01/12/23    222.00  [BALANCE]

               9-333506502743-GROCERIES

 01/12/23      UPI-SWIGGY-SWIGGY.STORES@AXISBANK-UTIB00     0000333510222540    01/12/23    473.00  [BALANCE]

               00100-333510222540-SWIGGY ORDER ID 16

 03/12/23      UPI-AMAZON INDIA-AMAZON@YAPL-YESB0APLUPI     0000333748773349    03/12/23    119.00  [BALANCE]

               -333748773349-YOU ARE PAYING FOR
```

### Obsidian Workflow Format
The same expenses would be represented in my Obsidian Budget Workflow as:

```
* #expense (name:Groceries) (amount:222) (date:2023-01-12) (category:Groceries)
* #expense (name:Food) (amount:473) (date:2023-01-12) (category:Food)
* #expense (name:Amazon) (amount:119) (date:2023-03-12) (category:Unknown)
```

The amount and date follow a one-to-one mapping with the values in the Bank Statement. Figuring out the name and the
category is slightly more involved. LLMs can help here.