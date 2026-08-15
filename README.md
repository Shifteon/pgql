## Testing grammar
Use ANTLR to build the java files. Output to the `java` folder
```bash
antlr4 -o java pgql.g4    
```
Go to the java directory and build the java
```bash
cd java && javac pgql*.java 
```
Input the query to test in `sample.txt`. Run using `grun`
```bash
grun pgql statement -gui sample.txt
```

## Compiling Go Code
Easy. Just run
```bash
antlr4 -Dlanguage=Go -o parser -visitor -no-listener pgql.g4
```