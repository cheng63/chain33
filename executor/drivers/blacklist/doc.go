package blacklist

/*
blacklist执行器作为一个demo,主要实现黑名单功能，用于在链上记录个人的信用情况。

问题记录：
	1.chain33底层数据库不支持多键值查询，以及模糊查询，影响业务功能的拓展性
       暂时采用多次存储和选择性规避。                                   完成
    2.chain33的查询，直接调用queryChain接口，当调用查询接口需要更改数据时，还要需要构建交易，这点不太友好 （系统架构问题）
    3.信用积分情况：
      机构每上传一次记录，就会获得10点积分，每查询一条记录需要支付2点积分，
      自己查询自己不需要支付积分。                    ps:暂时查询记录时积分扣除的业务逻辑没有实现
    4.每次查询需要查询几条记录呢？是否可以让用户选择,有待商榷       待完成
    5.新增登陆用户结构体，用户校验登陆用户是否合法。                  完成





*/
