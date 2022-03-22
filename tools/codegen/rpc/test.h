
//进入竞技场 (涉及到离线数据，走异步)
CL_ASYNC cl.SC_EnterArena RequestEnterArena(cl.PlaceHolder in) = 21700 ResponeEnterArena

//开始战斗 (涉及到离线数据，走异步)
CL_ASYNC cl.SC_BeginArenaBattle  RequestBeginArenaBattle(cl.CS_BeginArenaBattle in) = 21702 ResponeBeginArenaBattle

//战斗结果->胜利验证->同时回传需要翻牌的数据
CL_ASYNC cl.SC_ArenaBattleResult RequestArenaBattleResult(cl.CS_ArenaBattleResult in) = 21703 ResponeArenaBattleResult

//扫荡
CL cl.SC_SweepArena Sweep(cl.CS_SweepArena in) = 21704

//兑换
CL cl.SC_ExchangeIdx ArenaExchange(cl.CS_ExchangeIdx in) = 21705

//排行
CL_ASYNC cl.SC_ArenaRank RequestGetArenaRank(cl.CS_ArenaRank in) = 21706 ResponeGetArenaRank

//战报
CL cl.SC_ArenaBattleLog ArenaBattleLog(cl.PlaceHolder in) = 21707

//购买竞技场挑战次数
CL cl.SC_BuyArenaChallengeTimes BuyArenaChallengeTimes(cl.PlaceHolder in) = 21708

//被攻击排名发生变化通知
LC void PushArenaRankChange(cl.PlaceHolder out) = 21709

//战斗次数奖励
CL cl.SC_CommonRet GetBattleTimesReward(cl.CS_ExchangeIdx in) = 21910

//防守 走阵型保存
//商店 走通用商店
//查看玩家信息 走通用查看
