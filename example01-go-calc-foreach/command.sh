# 编译 antlr
antlr -Dlanguage=Go -o parser Calc.g4 

# run 代码
go run main.go


# example1: grun -gui
grun Calc start -gui
1 + 2
^D
