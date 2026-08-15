# ANTLR 4 Syntax & Mechanics Reference

## 1. Lexer vs. Parser Divide

ANTLR enforces a strict naming convention to distinguish between tokenization (lexing) and structure building (parsing).

| Rule Type | Casing | Input Handled | Purpose | Example |
| :--- | :--- | :--- | :--- | :--- |
| **Lexer (Terminals / Tokens)** | `UPPERCASE` | Raw Characters | Groups raw text into atomic tokens | `INT : [0-9]+ ;` |
| **Parser (Non-Terminals)** | `lowercase` | Token Stream | Groups tokens into hierarchical structures | `expr : INT '+' INT ;` |

> **Note on Literals:** Inline single-quoted strings in parser rules (e.g., `'for'`, `','`) are treated directly as terminal tokens.

---

## 2. ANTLR EBNF Operators Cheat Sheet

Every rule uses a colon `:` to separate the rule name from its body and ends with a semicolon `;`.

### A. Quantifiers (Repetition & Optionality)
* `?` $\rightarrow$ **Optional:** Matches **0 or 1** time.
  * *Example:* `'by'?`
* `*` $\rightarrow$ **Kleene Star:** Matches **0 or more** times.
  * *Example:* `(',' DIMENSION)*`
* `+` $\rightarrow$ **Kleene Plus:** Matches **1 or more** times.
  * *Example:* `[0-9]+`

### B. Combinators & Combinations
* **Sequence (Concatenation):** Space-separated elements must appear in exact order.
  * *Example:* `forClause : 'for' DIMENSION STRING ;`
* **Alternation / Choice (`|`):** Matches any one of the pipe-separated branches.
  * *Example:* `type : 'int' | 'float' | 'string' ;`
* **Grouping `( ... )`:** Combines elements so quantifiers or choices apply to the entire block.
  * *Example:* `(',' DIMENSION)*`

### C. Character Matchers (Lexer Regex)
* `[a-z]` $\rightarrow$ Matches any lowercase character.
* `[0-9]` $\rightarrow$ Matches any digit.
* `~[ ... ]` $\rightarrow$ **Inverted Set:** Matches any character *except* those listed.
  * *Example:* `~["\r\n]*` (matches anything except double quotes or line breaks).

---

## 3. How They Come Together (Architecture)
1. **Lexing Phase:** The input character stream is scanned from left to right. The Lexer matches character patterns using **UPPERCASE** rules and emits a flat stream of tokens.
2. **Parsing Phase:** The Parser evaluates the token stream using **lowercase** rules (starting at the root rule) to verify syntax and form a hierarchical **Parse Tree**.

### Example Integration
```antlr
grammar Sample;

/* Parser Rules (Structural Hierarchy) */
query      : showClause forClause? ;
showClause : 'show' measure ;
forClause  : 'for' dimension STRING ;

dimension  : PLAYER | WIN ;
measure    : KILLS | WIN ;

/* Lexer Rules (Tokens) */
WIN    : 'win' ;
KILLS  : 'kills' ;
PLAYER : 'player' ;
STRING : '"' ~["\r\n]* '"' ;
WS     : [ \t\r\n]+ -> skip ; // Discard whitespace