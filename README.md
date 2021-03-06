# datafoundry_plan

```
datafoundry套餐微服务
```

##数据库设计

```
CREATE TABLE IF NOT EXISTS DF_PLAN_REGION
(
    ID                TINYINT NOT NULL AUTO_INCREMENT,
    REGION            VARCHAR(30) NOT NULL,
    REGION_DESCRIBE   VARCHAR(128) NOT NULL,
    IDENTIFICATION    VARCHAR(128) NOT NULL,
    PRIMARY KEY (ID)
) DEFAULT CHARSET=UTF8;

CREATE TABLE IF NOT EXISTS DF_PLAN
(
    ID                INT NOT NULL AUTO_INCREMENT,
    PLAN_ID           VARCHAR(64) NOT NULL,
    PLAN_NAME         VARCHAR(128) NOT NULL,
    PLAN_TYPE         VARCHAR(128) NOT NULL,
    BELONG            VARCHAR(128) NOT NULL,
    PLAN_LEVEL        TINYINT NOT NULL,
    SPECIFICATION1    VARCHAR(128) NOT NULL,
    SPECIFICATION2    VARCHAR(128) NOT NULL,
    DESCRIPTION       TEXT,
    PRICE             DOUBLE(10,2)  NOT NULL,
    CYCLE             VARCHAR(2) NOT NULL,
    REGION            VARCHAR(30) NOT NULL,
    CREATE_TIME       DATETIME,
    STATUS            VARCHAR(2) NOT NULL,
    PRIMARY KEY (ID)
)  DEFAULT CHARSET=UTF8;
```

## API设计  

### POST /charge/v1/plans

创建一个套餐计划。

Body Parameters:
```
plan_name: 套餐名称
plan_type: 套餐类型
plan_level: 套餐等级
specification1: 套餐规格1
specification2: 套餐规格2
description： 套餐描述
price: 套餐价格
cycle: 计价周期
region_id: 套餐所在区域id
```

Return Result (json):
```
code: 返回码
msg: 返回信息
data.plan_id: 套餐id
```

### DELETE /charge/v1/plans/{planId}

删除一个套餐，并不是把套餐从表中删除，而是把状态从激活状态'Y'置为未激活状态'N'。

Path Parameters:
```
id: 套餐id
```

Return Result (json):

```
code: 返回码
msg: 返回信息
```

### PUT /charge/v1/plans/{planId}

更新一个套餐，新添加一个新的套餐计划再把原来的套餐计划置为未激活'N'。

Path Parameters:
```
planId: 套餐ID
```

Body Parameters:
注: 参数需要全部传入
```
plan_name: 套餐名称
plan_type: 套餐类型
plan_level: 套餐等级
specification1: 套餐规格1
specification2: 套餐规格2
description： 套餐描述
price: 套餐价格
cycle: 计价周期
region_id: 套餐所在区域id

```

Return Result (json):
```
code: 返回码
msg: 返回信息
```

### GET /charge/v1/plans/{planId}

查询一个套餐计划。

Path Parameters:
```
planId: 套餐ID
```

Return Result (json):
```
code: 返回码
msg: 返回信息
data.plan_id
data.plan_name
data.plan_type
data.belong
data.plan_level
data.specification1
data.specification2
data.description
data.price
data.cycle
data.region
data.create_time
data.status
```

### GET /charge/v1/plans

查询套餐列表

Query Parameters:
```
page: 第几页。可选。最小值为1。默认为1。
size: 每页最多返回多少条数据。可选。最小为1，最大为100。默认为30。
region: (可选)按套餐所属区域查询。
type: (可选)按套餐类型查询。
```

Return Result (json):
```
code: 返回码
msg: 返回信息
data.total
data.results
data.results[0].plan_id
data.results[0].plan_name
data.results[0].plan_type
data.results[0].belong
data.results[0].plan_level
data.results[0].specification1
data.results[0].specification2
data.results[0].description
data.results[0].price
data.results[0].cycle
data.results[0].region
data.results[0].create_time
data.results[0].status
...
```

### GET /charge/v1/query/plans/region

查询所有套餐所在区域

Return Result (json):
```
code: 返回码
msg: 返回信息
data.total
data.results
data..results[0].region
...
```


