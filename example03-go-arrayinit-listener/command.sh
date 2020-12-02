
# 生成antlr
antlr -Dlanguage=Go -o parser ArrayInit.g4

# 执行main.go, 将{1,2,3,4} -> 十六进制
go run main.go
