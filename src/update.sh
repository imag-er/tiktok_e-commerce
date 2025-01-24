#!/bin/bash
# 根据IDL更新代码

variable_list=("auth" "cart" "checkout" "order" "payment" "product" "user")
# 遍历列表
for VARIABLE in "${variable_list[@]}"; do
    echo "========================"
    echo "Processing ${VARIABLE}"
    
    kitex -module src -type protobuf -I idl/ idl/${VARIABLE}.proto
    mkdir -p rpc/${VARIABLE}
    cd rpc/${VARIABLE}
    kitex -module src -type protobuf -service src.${VARIABLE} -use src/kitex_gen -I ../../idl ../../idl/${VARIABLE}.proto
    # 执行完命令后，返回上级目录，以便下一次循环
    cd ../..
    echo "========================"
done
