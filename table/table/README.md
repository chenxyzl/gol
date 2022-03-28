## 表格代码生成器
---

### 文件说明：

* 执行脚本 gen_table.sh
* 表格代码生成器 gen_table.go
* 表格数据列表文件 table_list.txt
* 表格模板文件 table_template.txt

### 参数说明：

* -template 表格代码模板文件路径
* -tablelist 表格类型列表文件路径
* -out 输出目录

### 新增加表格的步骤：

1. 在logic/table/table_type.go中增加新的表格类型
2. 在table_list.txt中另起一行添加1中新增的类型
3. 执行gen_table.sh脚本
4. 会在logic/table目录生成如xxx_table.go文件
5. 在logic/table/table.go中添加新的表

### 多类型数据的自定义解析（参考table_test的TestSetting）

1. 使用 *json.RawMessage 来定义字段类型 如 value *json.RawMessage
2. 自定义想要的数据结构 res := make([][]int32, 0)
3. 调用json.Unmarshal(*value, &res) 