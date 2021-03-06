package main

import (
  "fmt"
	"github.com/chengchenginc/go-autocode/autocode"
)


func main() {
	sql := "CREATE TABLE `company` ( `id` int(11) NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `weight` tinyint(4) NOT NULL DEFAULT '0' COMMENT '权重', `intro` varchar(255) DEFAULT NULL COMMENT '简介', `level` tinyint(1) NOT NULL DEFAULT '0' COMMENT '阶段', `industry_type` tinyint(4) DEFAULT NULL COMMENT '行业类型', `address` varchar(255) DEFAULT NULL, `zip_code` char(6) DEFAULT NULL, `lng` varchar(15) DEFAULT NULL COMMENT '经度', `lat` varchar(15) DEFAULT NULL COMMENT '纬度', `create_time` datetime DEFAULT NULL, `site` varchar(255) DEFAULT NULL COMMENT '网址', `slogan` varchar(255) DEFAULT NULL COMMENT '口号', `stock_code` varchar(10) DEFAULT NULL, PRIMARY KEY (`id`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT"
	template,_ := autocode.Gen(sql, "./template/beego-orm-domain.tpl")
	fmt.Printf(template)
}
