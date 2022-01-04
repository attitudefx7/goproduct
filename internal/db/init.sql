CREATE TABLE `pms_brand` (
 `id` BIGINT ( 20 ) NOT NULL AUTO_INCREMENT,
 `name` VARCHAR ( 64 ) DEFAULT NULL COMMENT '名称',
 `first_letter` VARCHAR ( 8 ) DEFAULT NULL COMMENT '首字母',
 `sort` INT ( 11 ) DEFAULT NULL COMMENT '排序',
 `factory_status` INT ( 1 ) DEFAULT NULL COMMENT '是否为品牌制造商：0->不是；1->是',
 `show_status` INT ( 1 ) DEFAULT NULL COMMENT '是否显示',
 `product_count` INT ( 11 ) DEFAULT NULL COMMENT '产品数量',
 `product_comment_count` INT ( 11 ) DEFAULT NULL COMMENT '产品评论数量',
 `logo` VARCHAR ( 255 ) DEFAULT NULL COMMENT '品牌logo',
 `big_pic` VARCHAR ( 255 ) DEFAULT NULL COMMENT '专区大图',
 `brand_story` text COMMENT '品牌故事',
 PRIMARY KEY ( `id` )
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `pms_product` (
 `id` BIGINT ( 20 ) NOT NULL AUTO_INCREMENT,
 `brand_id` BIGINT ( 20 ) DEFAULT NULL COMMENT '品牌id',
 `product_category_id` BIGINT ( 20 ) DEFAULT NULL COMMENT '品牌分类id',
 `feight_template_id` BIGINT ( 20 ) DEFAULT '0' COMMENT '运费模版id',
 `product_attribute_category_id` BIGINT ( 20 ) DEFAULT NULL COMMENT '品牌属性分类id',
 `name` VARCHAR ( 64 ) NOT NULL COMMENT '商品名称',
 `pic` VARCHAR ( 255 ) DEFAULT NULL COMMENT '图片',
 `product_sn` VARCHAR ( 64 ) NOT NULL COMMENT '货号',
 `delete_status` INT ( 1 ) DEFAULT NULL COMMENT '删除状态：0->未删除；1->已删除',
 `publish_status` INT ( 1 ) DEFAULT NULL COMMENT '上架状态：0->下架；1->上架',
 `new_status` INT ( 1 ) DEFAULT NULL COMMENT '新品状态:0->不是新品；1->新品',
 `recommand_status` INT ( 1 ) DEFAULT NULL COMMENT '推荐状态；0->不推荐；1->推荐',
 `sort` INT ( 11 ) DEFAULT NULL COMMENT '排序',
 `sale` INT ( 11 ) DEFAULT NULL COMMENT '销量',
 `price` DECIMAL ( 10, 2 ) DEFAULT NULL COMMENT '价格',
 `promotion_price` DECIMAL ( 10, 2 ) DEFAULT NULL COMMENT '促销价格',
 `sub_title` VARCHAR ( 255 ) DEFAULT NULL COMMENT '副标题',
 `description` text COMMENT '商品描述',
 `stock` INT ( 11 ) DEFAULT NULL COMMENT '库存',
 `low_stock` INT ( 11 ) DEFAULT NULL COMMENT '库存预警值',
 `keywords` VARCHAR ( 255 ) DEFAULT NULL COMMENT '关键字',
 `note` VARCHAR ( 255 ) DEFAULT NULL COMMENT '备注',
 `album_pics` VARCHAR ( 255 ) DEFAULT NULL COMMENT '画册图片，连产品图片限制为5张，以逗号分割',
 `detail_title` VARCHAR ( 255 ) DEFAULT NULL COMMENT '详情标题',
 `detail_desc` text COMMENT '详情描述',
 `product_category_name` VARCHAR ( 255 ) DEFAULT NULL COMMENT '产品分类名称',
 `brand_name` VARCHAR ( 255 ) DEFAULT NULL COMMENT '品牌名称',
 PRIMARY KEY ( `id` )
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `pms_product_attribute` ( `id` BIGINT ( 20 ) NOT NULL AUTO_INCREMENT, `name` VARCHAR ( 64 ) DEFAULT NULL COMMENT '名称', PRIMARY KEY ( `id` ) ) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `pms_product_attribute_value` (
 `id` BIGINT ( 20 ) NOT NULL AUTO_INCREMENT,
 `product_id` BIGINT ( 20 ) DEFAULT NULL COMMENT '商品id',
 `attribute_id` BIGINT ( 20 ) DEFAULT NULL COMMENT '商品属性id',
 `value` VARCHAR ( 64 ) DEFAULT NULL COMMENT '参数值',
 PRIMARY KEY ( `id` )
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `pms_product_category` (
 `id` BIGINT ( 20 ) NOT NULL AUTO_INCREMENT,
 `parent_id` BIGINT ( 20 ) DEFAULT NULL COMMENT '上级分类的编号：0表示一级分类',
 `name` VARCHAR ( 64 ) DEFAULT NULL COMMENT '名称',
 `level` INT ( 1 ) DEFAULT NULL COMMENT '分类级别：0->1级；1->2级',
 `product_count` INT ( 11 ) DEFAULT NULL COMMENT '商品数量',
 `product_unit` VARCHAR ( 64 ) DEFAULT NULL COMMENT '商品单位',
 `nav_status` INT ( 1 ) DEFAULT NULL COMMENT '是否显示在导航栏：0->不显示；1->显示',
 `show_status` INT ( 1 ) DEFAULT NULL COMMENT '显示状态：0->不显示；1->显示',
 `sort` INT ( 11 ) DEFAULT NULL COMMENT '排序',
 `icon` VARCHAR ( 255 ) DEFAULT NULL COMMENT '图标',
 `keywords` VARCHAR ( 255 ) DEFAULT NULL COMMENT '关键字',
 `description` text COMMENT '描述',
 PRIMARY KEY ( `id` )
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `pms_sku_stock` (
 `id` BIGINT ( 20 ) NOT NULL AUTO_INCREMENT,
 `product_id` BIGINT ( 20 ) DEFAULT NULL COMMENT '商品id',
 `sku_code` VARCHAR ( 64 ) NOT NULL COMMENT 'sku编码',
 `price` DECIMAL ( 10, 2 ) DEFAULT NULL COMMENT '价格',
 `stock` INT ( 11 ) DEFAULT '0' COMMENT '库存',
 `low_stock` INT ( 11 ) DEFAULT NULL COMMENT '预警库存',
 `pic` VARCHAR ( 255 ) DEFAULT NULL COMMENT '展示图片',
 `sale` INT ( 11 ) DEFAULT NULL COMMENT '销量',
 `promotion_price` DECIMAL ( 10, 2 ) DEFAULT NULL COMMENT '单品促销价格',
 `lock_stock` INT ( 11 ) DEFAULT '0' COMMENT '锁定库存',
 `sp_data` VARCHAR ( 500 ) DEFAULT '' COMMENT '商品销售属性，json格式',
 PRIMARY KEY ( `id` )
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;