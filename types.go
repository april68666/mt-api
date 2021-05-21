package meituanapi

// OrderListParams 订单列表查询参数
type OrderListParams struct {
	//0 团购订单
	//2 酒店订单
	//4 外卖订单
	//5 话费订单
	//6 闪购订单
	Type string `json:"type"`
	//查询起始时间10位时间戳，以下单时间为准
	StartTime string `json:"startTime"`
	//查询截止时间10位时间戳，以下单时间为准
	EndTime string `json:"endTime"`
	//分页参数，起始值从1开始
	Page string `json:"page"`
	//每页显示数据条数，最大值为100
	Limit string `json:"limit"`
	//查询时间类型，枚举值,默认1
	//1 按订单支付时间查询
	//2 按订单发生修改时间查询
	QueryTimeType string `json:"queryTimeType"`
}

// OrderListResp 订单列表返回数据
type OrderListResp struct {
	Msg      string `json:"msg"`    // 如果被限流会返回信息
	Des      string `json:"des"`    //异常描述信息
	Status   int    `json:"status"` //状态值，0为成功，非0为异常
	DataList []struct {
		AppKey       string `json:"appkey"`       //订单对应的appkey，外卖、话费、闪购订单会返回该字段
		OrderID      string `json:"orderid"`      //订单id
		PayPrice     string `json:"payprice"`     //订单用户实际支付金额
		PayTime      string `json:"paytime"`      //订单支付时间，10位时间戳
		Profit       string `json:"profit"`       //订单预估返佣金额
		Sid          string `json:"sid"`          //订单对应的推广位sid
		SmsTitle     string `json:"smstitle"`     //订单标题
		RefundPrice  string `json:"refundprice"`  //订单实际退款金额，外卖、话费、闪购订单若发生退款会返回该字段
		RefundTime   string `json:"refundtime"`   //订单退款时间，10位时间戳，外卖、话费、闪购订单若发生退款会返回该字段
		RefundProfit string `json:"refundprofit"` //订单需要扣除的返佣金额，外卖、话费、闪购订单若发生退款会返回该字段
		//订单状态，外卖、话费、闪购订单会返回该字段
		//1 已付款
		//8 已完成
		//9 已退款或风控
		Status int `json:"status"`
		//订单的奖励类型
		//话费订单类型返回该字段
		//3 首购奖励
		//5 留存奖励
		TradeTypeList string `json:"tradeTypeList"`
		//奖励类型对应平台类型的映射
		//格式：{3:[3,5]}
		//value的枚举值：
		//1 外卖
		//2 分销酒店
		//3 平台
		//4 券类型酒店
		//5 团好货
		TradeTypeBusinessTypeMapStr string `json:"tradeTypeBusinessTypeMapStr"`
	} `json:"dataList"`
	Total int `json:"total"` //查询条件命中的总数据条数，用于计算分页参数
}

// OrderInfoParams 单个订单信息查询参数
type OrderInfoParams struct {
	//0 团购订单
	//2 酒店订单
	//4 外卖订单
	//5 话费订单
	//6 闪购订单
	Type string `json:"type"`
	OID  string `json:"oid"`  //订单id
	Full string `json:"full"` //是否返回完整订单信息(即是否包含返佣、退款信息),需要完整信息full=1
}

// OrderInfoResp 单个订单信息查询返回数据
type OrderInfoResp struct {
	Msg    string `json:"msg"`    // 如果被限流会返回信息
	Des    string `json:"des"`    //异常描述信息
	Status int    `json:"status"` //状态值，0为成功，非0为异常
	Order  struct {
		OrderID  string `json:"orderid"`  //订单id
		UID      string `json:"uid"`      //联盟媒体id
		Sid      string `json:"sid"`      //推广位sid
		Total    string `json:"total"`    //订单总金额
		Direct   string `json:"direct"`   //订单用户实付金额
		Quantity string `json:"quantity"` //订单包含的团购券数量(团购订单该字段有意义，其它订单不用考虑该字段)
		DealID   string `json:"dealid"`   //订单所含的商品id
		SmsTitle string `json:"smstitle"` //订单标题
		PayTime  string `json:"paytime"`  //订单付款时间，10位时间戳
		ModTime  string `json:"modtime"`  //订单记录更新时间，10位时间戳
		Sign     string `json:"sign"`     //订单主体信息签名
		//订单状态，外卖、话费、闪购订单会返回该字段
		//1 已付款
		//8 已完成
		//9 已退款或风控
		Status int `json:"status"`
	} `json:"order"`
	Coupon []struct {
		OrderID  string `json:"orderid"`  //订单id
		Sequence string `json:"sequence"` //核销序列号
		UseTime  string `json:"usetime"`  //核销时间
		Price    string `json:"price"`    //核销实际支付金额
		Profit   string `json:"profit"`   //核销实际返佣金额
	} `json:"coupon"` //返佣信息,根据full值决定是否回传
	Refund []struct {
		OrderID    string `json:"orderid"`    //订单id
		Quantity   string `json:"quantity"`   //退款笔数
		RefundTime string `json:"refundtime"` //退款时间，10位时间戳
		Money      string `json:"money"`      //实际退款金额
	} `json:"refund"` //退款信息,根据full值决定是否回传
}

// CouponListParams 优惠券列表查询参数
type CouponListParams struct {
	//0 团购订单
	//2 酒店订单
	//4 外卖订单
	//5 话费订单
	//6 闪购订单
	Type      string `json:"type"`
	StartTime string `json:"startTime"`
	//查询截止时间10位时间戳，以下单时间为准
	EndTime string `json:"endTime"`
	//分页参数，起始值从1开始
	Page string `json:"page"`
	//每页显示数据条数，最大值为100
	Limit string `json:"limit"`
	Sid   string `json:"sid"` //推广位sid
}

// CouponListResp 优惠券列表返回数据
type CouponListResp struct {
	Msg      string `json:"msg"`    // 如果被限流会返回信息
	Des      string `json:"des"`    //异常描述信息
	Status   int    `json:"status"` //状态值，0为成功，非0为异常
	Total    int    `json:"total"`
	DataList []struct {
		AppKey      string `json:"appKey"`      //媒体appkey
		Sid         string `json:"sid"`         //推广位sid
		CouponTime  string `json:"couponTime"`  //领券日期yyyy-MM-dd HH:mm:ss
		Money       string `json:"money"`       //券优惠金额
		MinUseMoney string `json:"minUseMoney"` //用券门槛金额
		CouponName  string `json:"couponName"`  //券名称
		CouponType  string `json:"couponType"`  //券类型
		BeginTime   int    `json:"beginTime"`   //券生效起始时间，10位时间戳
		EndTime     int    `json:"endTime"`     //券生效截止时间，10位时间戳
		CouponCode  string `json:"couponCode"`  //券唯一标识
	} `json:"dataList"`
}

// GenerateLinkParams 自助取链请求参数
type GenerateLinkParams struct {
	ActID int64  `json:"actId"` //活动id，可以在联盟活动列表中查看获取
	Sid   string `json:"sid"`   //推广位sid，支持通过接口自定义创建，不受平台200个上限限制，长度不能超过64个字符，支持小写字母和数字，历史已创建的推广位不受这个约束
	//链接类型，枚举值：
	//1 h5链接
	//2 deeplink(唤起)链接
	//3 中间页唤起链接
	//4 微信小程序唤起路径
	LinkType int `json:"linkType"`
}

// GenerateLinkResp 自助取链返回数据
type GenerateLinkResp struct {
	Msg    string `json:"msg"`    // 如果被限流会返回信息
	Des    string `json:"des"`    //异常描述信息
	Status int    `json:"status"` //状态值，0为成功，非0为异常
	Data   string `json:"data"`   //最终的推广链接
}
