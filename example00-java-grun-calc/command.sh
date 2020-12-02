# 配置classpath
# 找到antlr的路径, 修改.bashrc中classpath
# export CLASSPATH="/usr/local/Cellar/antlr/4.8_1/antlr-4.8-complete.jar:$CLASSPATH"

# 生成antlr
antlr -o parser Calc.g4

# 编译
cd parser/
javac ./*.java

# example1: grun -gui
grun Calc start -gui
1 + 2 * 3 + 4 * 6
^D

# example2: grun -tokens
grun Calc start -tokens
1 + 2 * 3 + 4 * 6
^D
