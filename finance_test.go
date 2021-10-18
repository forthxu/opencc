package opencc

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestFinance(t *testing.T) {
	raw := `1、 联合健康 Q3 净营收 723.37 亿美元，去年同期为 651.15 亿美元。净利润为 41.91 亿美元，去年同期为 32.56 亿美元。
2、人民网发文称富途、老虎等跨境互联网券商在用户信息安全以及合法化、合规化方面存在风险。
3、【腾讯看点即将更名为腾讯 “信息平台与服务线” 部分团队将提升至事业部编制】本周五，腾讯控股 PCG 旗下的腾讯看点将正式官宣更名，新名称为 “信息平台与服务线”，这一组织架构调整的主要原因是合并部分搜狗团队，同时进行战略转型。据瞭解，在新的 “信息平台与服务线 “下，部分中心将与搜狗团队深度整合，升级为事业部（BU）编制。腾讯内部多位知晓调整情况的人士表示，新业务线的 “服务 “定位未来将主要由搜索引擎业务承担。（界面）
4、据美国证券交易委员会文件：特殊目的收购公司 (SPAC) 数位医疗 Digital Health 递交 IPO 申请，寻求在纳斯达克上市。
5、台积电三季度营收 148.78 亿美元，较去年三季度的 121.38 亿美元增加 27.4 亿美元，同比增长 22.6%。净利润方面，财报显示为 56.14 亿美元，去年三季度为 46.78 亿美元，增加 9.36 亿美元，同比增长 20%。台积电今年第四季度的销售额为 154-157 亿美元；期内毛利率 51%-53%。全年毛利率预计将高于 50%，收入增速约为 24%。台积电最快可在 2025 年拥有 2 纳米芯片技术。
6、【下半年以来超 250 亿资金借道 ETF 抄底中概互联网板块】7 月以来，借道中概互联、恒生互联等相关 ETF 抄底互联网的资金超过了 250 亿。多位业内人士认为，中概互联网巨头公司的安全边际正逐步显现，股价接近底部。从回调力度和估值角度看，现在配置中概股或是性价比较高的选择。（证券时报）
7、富国银行第三季度营收 188.3 亿美元，预估 184 亿美元；每股收益 1.17 美元。
8、美国银行 Q3 净营收 228 亿美元，去年同期为 203 亿美元；净利润为 77 亿美元，去年同期为 49 亿美元；财富与投资管理总收入为 53.1 亿美元，估计为 52.1 亿美元；投行收入 21.7 亿美元；每股收益 85 美分。
9、抖音否认进入外卖行业：相关招商、代理信息不实。（界面）
盘前异动
美股盘前，道指期货涨 0.6%，纳指期货涨 0.77%，标普 500 指数期货涨 0.68%。
美大型科技股普涨，推特、谷歌、facebook 涨 1%，微软、特斯拉涨约 0.8%，奈飞涨 0.6%。
美芯片股普涨，台积电涨近 4%，安森美半导体、英伟达、AMD 涨 1.6%，美光科技涨超 1%，英特尔涨 0.8%。
热门中概科网股普跌，阿里巴巴跌 0.4%，拼多多跌近 1%，京东跌 0.3%，腾讯音乐跌 1.3%，贝壳跌 2.7%，滴滴跌 0.7%，百度跌 0.4%，网易跌 0.5%。
阿斯麦涨超 3%，此前业内预计阿斯麦市值明年将达 5000 亿美元。
台积电涨 3.55%，第三季度营收同比增长 22.6%，净利润再超 50 亿美元。预计四季度营收超 150 亿美元。
富途控股、老虎证券跌超 10%，人民网发文称富途、老虎等跨境互联网券商在用户信息安全以及合法化、合规化方面存在风险。
苹果涨 0.74%，韦德布什分析师认为 iPhone 超级周期不受减产影响，苹果股价将继续跑赢大盘。
机构观点
1、美联储紧缩脚步临近 美债收益率曲线释放危险衰退信号
美债收益率曲线趋平释放信号，市场对美联储遏制通胀的努力会破坏美国的复苏愈发担心。美国消费者价格再度超出预期之际，30 年期和 5 年期美债收益率之差本周跌向 17 个月低点。目前该收益率差在 100 个基点左右，自 6 月份美联储官员暗示在 2023 年底前将加息两次以来，美债收益率曲线的趋平态势就开始加速。Asset Management One Co.基金经理 Akira Takei 称，这反映出市场对增长放缓及高通胀并存的预期。供应瓶颈驱动的高通胀将减少需求并损害经济。
2、摩根士丹利：予苹果跑赢大市评级 目标价 168 美元
摩根士丹利分析师 Katy Huberty 认为，在全球供应困境中，投资者还是应该 “逢低买入” 苹果股票。Huberty 承认，更广泛的供应困境是市场上的一个 “真正的问题”。该分析师在周二的一份致客户报告中写道，苹果今年迄今的营收表现超出预期，这符合该行的看法，即苹果 “在供应紧张时期将获得最优级待遇”。她还认为，对苹果来说，“需求是不会消失的”，苹果拥有强大的客户忠诚度，这可能有助于将 iPhone 需求推至后面几个季度。Huberty 指出，如果媒体有关苹果减产的报道是准确的，预计当前季度的 iPhone 出货量可能会发生变化，但这对 2022 财年的出货量预期影响不大。她仍预计苹果 2022 财年的 iPhone 出货量将达到 2.385 亿部。Huberty 对苹果的评级为跑赢大市，目标价 168 美元。
3、苹果 iPhone 减产或影响关键假日购物季营收
分析人士表示，苹果如果调低今年 iPhone 13 系列智能手机的产量目标，可能会影响苹果公司在关键假日购物季的营收。美国银行分析师 Wamsi Mohan 在一份报告中表示，周二有报道称苹果 iPhone 13 系列智能手机可能减产，这或将导致今年第四季度苹果新款手机销售额低于预期。尽管美国银行预计苹果第三季度销售额将超过预期，但 Mohan 表示，“限制因素可能对第四季度销售产生的影响更大，我们预计第四季度苹果销售额将低于市场平均预期。” 莫汉对苹果股票给出了中性评级。苹果可能的减产目标也影响了思佳讯、意法半导体和日本显示器等苹果设备供应商的股价。
美股策略
1、股价低、增长前景好，这五大矿业美股值得拥有
2、阿里股价反弹到位了吗？摩根大通：还能再涨 50%！
3、真超过 150 亿美元 台积电预计 Q4 营收 154-157 亿美元
4、大摩谈苹果减产：买入良机来了 供应困境中竞争对手 “日子更难过”
5、迷雾剧场能够浇灭爱奇艺的盈利焦虑吗？
6、特斯拉财报前瞻：Q3 交付量大超预期 营收料再创新高
7、台积电 Q3 净利同比增长 13.8%，7nm 制程营收占比增至 34%
8、供应链困难持续，第二代特斯拉 Roadster 生产再次推迟
9、抛弃特斯拉，华尔街多空势力嗜血中国新造车
10、特斯拉联合创始人：别盲目制定电动化时间表，要看供应链能否跟得上`

	expected := `1、 聯合健康 Q3 淨營收 723.37 億美元，去年同期為 651.15 億美元。淨利潤為 41.91 億美元，去年同期為 32.56 億美元。
2、人民網發文稱富途、老虎等跨境互聯網券商在用户信息安全以及合法化、合規化方面存在風險。
3、【騰訊看點即將更名為騰訊 “信息平台與服務線” 部分團隊將提升至事業部編制】本週五，騰訊控股 PCG 旗下的騰訊看點將正式官宣更名，新名稱為 “信息平台與服務線”，這一組織架構調整的主要原因是合併部分搜狗團隊，同時進行戰略轉型。據瞭解，在新的 “信息平台與服務線 “下，部分中心將與搜狗團隊深度整合，升級為事業部（BU）編制。騰訊內部多位知曉調整情況的人士表示，新業務線的 “服務 “定位未來將主要由搜索引擎業務承擔。（界面）
4、據美國證券交易委員會文件：特殊目的收購公司 (SPAC) 數位醫療 Digital Health 遞交 IPO 申請，尋求在納斯達克上市。
5、台積電三季度營收 148.78 億美元，較去年三季度的 121.38 億美元增加 27.4 億美元，同比增長 22.6%。淨利潤方面，財報顯示為 56.14 億美元，去年三季度為 46.78 億美元，增加 9.36 億美元，同比增長 20%。台積電今年第四季度的銷售額為 154-157 億美元；期內毛利率 51%-53%。全年毛利率預計將高於 50%，收入增速約為 24%。台積電最快可在 2025 年擁有 2 納米芯片技術。
6、【下半年以來超 250 億資金借道 ETF 抄底中概互聯網板塊】7 月以來，借道中概互聯、恆生互聯等相關 ETF 抄底互聯網的資金超過了 250 億。多位業內人士認為，中概互聯網巨頭公司的安全邊際正逐步顯現，股價接近底部。從回調力度和估值角度看，現在配置中概股或是性價比較高的選擇。（證券時報）
7、富國銀行第三季度營收 188.3 億美元，預估 184 億美元；每股收益 1.17 美元。
8、美國銀行 Q3 淨營收 228 億美元，去年同期為 203 億美元；淨利潤為 77 億美元，去年同期為 49 億美元；財富與投資管理總收入為 53.1 億美元，估計為 52.1 億美元；投行收入 21.7 億美元；每股收益 85 美分。
9、抖音否認進入外賣行業：相關招商、代理信息不實。（界面）
盤前異動
美股盤前，道指期貨漲 0.6%，納指期貨漲 0.77%，標普 500 指數期貨漲 0.68%。
美大型科技股普漲，推特、谷歌、facebook 漲 1%，微軟、特斯拉漲約 0.8%，奈飛漲 0.6%。
美芯片股普漲，台積電漲近 4%，安森美半導體、英偉達、AMD 漲 1.6%，美光科技漲超 1%，英特爾漲 0.8%。
熱門中概科網股普跌，阿里巴巴跌 0.4%，拼多多跌近 1%，京東跌 0.3%，騰訊音樂跌 1.3%，貝殼跌 2.7%，滴滴跌 0.7%，百度跌 0.4%，網易跌 0.5%。
阿斯麥漲超 3%，此前業內預計阿斯麥市值明年將達 5000 億美元。
台積電漲 3.55%，第三季度營收同比增長 22.6%，淨利潤再超 50 億美元。預計四季度營收超 150 億美元。
富途控股、老虎證券跌超 10%，人民網發文稱富途、老虎等跨境互聯網券商在用户信息安全以及合法化、合規化方面存在風險。
蘋果漲 0.74%，韋德布什分析師認為 iPhone 超級週期不受減產影響，蘋果股價將繼續跑贏大盤。
機構觀點
1、美聯儲緊縮腳步臨近 美債收益率曲線釋放危險衰退信號
美債收益率曲線趨平釋放信號，市場對美聯儲遏制通脹的努力會破壞美國的復甦愈發擔心。美國消費者價格再度超出預期之際，30 年期和 5 年期美債收益率之差本週跌向 17 個月低點。目前該收益率差在 100 個點子左右，自 6 月份美聯儲官員暗示在 2023 年底前將加息兩次以來，美債收益率曲線的趨平態勢就開始加速。Asset Management One Co.基金經理 Akira Takei 稱，這反映出市場對增長放緩及高通脹並存的預期。供應瓶頸驅動的高通脹將減少需求並損害經濟。
2、摩根士丹利：予蘋果跑贏大市評級 目標價 168 美元
摩根士丹利分析師 Katy Huberty 認為，在全球供應困境中，投資者還是應該 “逢低買入” 蘋果股票。Huberty 承認，更廣泛的供應困境是市場上的一個 “真正的問題”。該分析師在週二的一份致客户報告中寫道，蘋果今年迄今的營收表現超出預期，這符合該行的看法，即蘋果 “在供應緊張時期將獲得最優級待遇”。她還認為，對蘋果來説，“需求是不會消失的”，蘋果擁有強大的客户忠誠度，這可能有助於將 iPhone 需求推至後面幾個季度。Huberty 指出，如果媒體有關蘋果減產的報道是準確的，預計當前季度的 iPhone 出貨量可能會發生變化，但這對 2022 財年的出貨量預期影響不大。她仍預計蘋果 2022 財年的 iPhone 出貨量將達到 2.385 億部。Huberty 對蘋果的評級為跑贏大市，目標價 168 美元。
3、蘋果 iPhone 減產或影響關鍵假日購物季營收
分析人士表示，蘋果如果調低今年 iPhone 13 系列智能手機的產量目標，可能會影響蘋果公司在關鍵假日購物季的營收。美國銀行分析師 Wamsi Mohan 在一份報告中表示，週二有報道稱蘋果 iPhone 13 系列智能手機可能減產，這或將導致今年第四季度蘋果新款手機銷售額低於預期。儘管美國銀行預計蘋果第三季度銷售額將超過預期，但 Mohan 表示，“限制因素可能對第四季度銷售產生的影響更大，我們預計第四季度蘋果銷售額將低於市場平均預期。” 莫漢對蘋果股票給出了中性評級。蘋果可能的減產目標也影響了思佳訊、意法半導體和日本顯示器等蘋果設備供應商的股價。
美股策略
1、股價低、增長前景好，這五大礦業美股值得擁有
2、阿里股價反彈到位了嗎？摩根大通：還能再漲 50%！
3、真超過 150 億美元 台積電預計 Q4 營收 154-157 億美元
4、大摩談蘋果減產：買入良機來了 供應困境中競爭對手 “日子更難過”
5、迷霧劇場能夠澆滅愛奇藝的盈利焦慮嗎？
6、特斯拉財報前瞻：Q3 交付量大超預期 營收料再創新高
7、台積電 Q3 淨利同比增長 13.8%，7nm 製程營收佔比增至 34%
8、供應鏈困難持續，第二代特斯拉 Roadster 生產再次推遲
9、拋棄特斯拉，華爾街多空勢力嗜血中國新造車
10、特斯拉聯合創始人：別盲目制定電動化時間表，要看供應鏈能否跟得上`

	cc, _ := New("s2hk-finance")
	out, _ := cc.Convert(raw)

	fmt.Println(out)

	if strings.TrimSpace(expected) != strings.TrimSpace(out) {
		dmp := diffmatchpatch.New()

		diffs := dmp.DiffMain(expected, out, true)

		t.Errorf(dmp.DiffPrettyText(diffs))
	}

}