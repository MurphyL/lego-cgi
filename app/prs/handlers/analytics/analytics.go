package analytics

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"murphyl.com/app/prs/middleware"
)

/**
 * Analytics 数据分析
 */

// AnalyticsHandler 数据分析处理器
type AnalyticsHandler struct {
	db *gorm.DB
}

// NewAnalyticsHandler 创建数据分析处理器
func NewAnalyticsHandler(dao *gorm.DB) func(router fiber.Router) {
	return func(router fiber.Router) {
		h := &AnalyticsHandler{db: dao}
		h.RegisterRoutes(router)
	}
}

// RegisterRoutes 注册路由
func (h *AnalyticsHandler) RegisterRoutes(router fiber.Router) {
	// 房源分析
	router.Get("/analytics/property/vacancy-rate", middleware.AuthMiddleware("analytics:property_vacancy"), h.GetPropertyVacancyRate)
	router.Get("/analytics/property/rental-trend", middleware.AuthMiddleware("analytics:property_rental_trend"), h.GetRentalTrend)
	router.Get("/analytics/property/heat-rank", middleware.AuthMiddleware("analytics:property_heat_rank"), h.GetPropertyHeatRank)

	// 租户分析
	router.Get("/analytics/tenant/level-distribution", middleware.AuthMiddleware("analytics:tenant_level"), h.GetTenantLevelDistribution)
	router.Get("/analytics/tenant/demand-distribution", middleware.AuthMiddleware("analytics:tenant_demand"), h.GetTenantDemandDistribution)

	// 财务分析
	router.Get("/analytics/finance/income-expense-trend", middleware.AuthMiddleware("analytics:finance_income_expense"), h.GetIncomeExpenseTrend)
	router.Get("/analytics/finance/cash-flow", middleware.AuthMiddleware("analytics:finance_cash_flow"), h.GetCashFlow)

	// 运营分析
	router.Get("/analytics/operation/viewing-conversion", middleware.AuthMiddleware("analytics:operation_conversion"), h.GetViewingConversionRate)
	router.Get("/analytics/operation/performance", middleware.AuthMiddleware("analytics:operation_performance"), h.GetOperationPerformance)
}

// GetPropertyVacancyRate 获取房源空置率
func (h *AnalyticsHandler) GetPropertyVacancyRate(c fiber.Ctx) error {
	// 计算总房源数
	var totalProperties int64
	h.db.Model(&struct{}{}).Table("hrs_property").Count(&totalProperties)

	// 计算空置房源数
	var vacantProperties int64
	h.db.Model(&struct{}{}).Table("hrs_property").Where("status = ?", 0).Count(&vacantProperties)

	// 计算空置率
	vacancyRate := 0.0
	if totalProperties > 0 {
		vacancyRate = float64(vacantProperties) / float64(totalProperties) * 100
	}

	// 按区域统计空置率
	type AreaVacancy struct {
		Area        string  `json:"area"`
		VacancyRate float64 `json:"vacancy_rate"`
		TotalCount  int64   `json:"total_count"`
		VacantCount int64   `json:"vacant_count"`
	}

	var areaVacancies []AreaVacancy
	// 这里简化处理，实际应根据地址字段提取区域信息
	h.db.Raw(`
		SELECT 
			SUBSTRING_INDEX(address, '区', 1) as area, 
			COUNT(*) as total_count, 
			SUM(CASE WHEN status = 0 THEN 1 ELSE 0 END) as vacant_count,
			(SUM(CASE WHEN status = 0 THEN 1 ELSE 0 END) / COUNT(*)) * 100 as vacancy_rate
		FROM hrs_property 
		GROUP BY area
	`).Scan(&areaVacancies)

	return c.JSON(fiber.Map{
		"total_properties":  totalProperties,
		"vacant_properties": vacantProperties,
		"vacancy_rate":      vacancyRate,
		"area_vacancies":    areaVacancies,
	})
}

// GetRentalTrend 获取租金趋势
func (h *AnalyticsHandler) GetRentalTrend(c fiber.Ctx) error {
	// 获取最近6个月的租金趋势
	endDate := time.Now()
	startDate := endDate.AddDate(0, -6, 0)

	type MonthlyRental struct {
		Month   string  `json:"month"`
		AvgRent float64 `json:"avg_rent"`
		Count   int64   `json:"count"`
	}

	var monthlyRentals []MonthlyRental
	h.db.Raw(`
		SELECT 
			DATE_FORMAT(create_time, '%Y-%m') as month, 
			AVG(price) as avg_rent, 
			COUNT(*) as count
		FROM hrs_property 
		WHERE create_time BETWEEN ? AND ?
		GROUP BY month
		ORDER BY month
	`, startDate, endDate).Scan(&monthlyRentals)

	return c.JSON(monthlyRentals)
}

// GetPropertyHeatRank 获取房源热度排行
func (h *AnalyticsHandler) GetPropertyHeatRank(c fiber.Ctx) error {
	// 根据带看次数和关注程度计算热度
	type PropertyHeat struct {
		PropertyID    uint    `json:"property_id"`
		PropertyTitle string  `json:"property_title"`
		Address       string  `json:"address"`
		Price         float64 `json:"price"`
		ViewingCount  int64   `json:"viewing_count"`
		HeatScore     int64   `json:"heat_score"`
	}

	var propertyHeats []PropertyHeat
	h.db.Raw(`
		SELECT 
			p.id as property_id, 
			p.property_title, 
			p.address, 
			p.price, 
			COUNT(v.id) as viewing_count,
			COUNT(v.id) as heat_score
		FROM hrs_property p
		LEFT JOIN hrs_property_viewing v ON p.id = v.property_id
		GROUP BY p.id
		ORDER BY heat_score DESC
		LIMIT 10
	`).Scan(&propertyHeats)

	return c.JSON(propertyHeats)
}

// GetTenantLevelDistribution 获取租户等级分布
func (h *AnalyticsHandler) GetTenantLevelDistribution(c fiber.Ctx) error {
	type LevelDistribution struct {
		Level      uint8   `json:"level"`
		LevelName  string  `json:"level_name"`
		Count      int64   `json:"count"`
		Percentage float64 `json:"percentage"`
	}

	var distributions []LevelDistribution
	h.db.Raw(`
		SELECT 
			level, 
			CASE 
				WHEN level = 0 THEN '普通' 
				WHEN level = 1 THEN 'VIP' 
				ELSE '其他' 
			END as level_name, 
			COUNT(*) as count
		FROM hrs_tenant 
		GROUP BY level
	`).Scan(&distributions)

	// 计算总租户数
	var totalTenants int64
	h.db.Model(&struct{}{}).Table("hrs_tenant").Count(&totalTenants)

	// 计算百分比
	for i := range distributions {
		if totalTenants > 0 {
			distributions[i].Percentage = float64(distributions[i].Count) / float64(totalTenants) * 100
		}
	}

	return c.JSON(distributions)
}

// GetTenantDemandDistribution 获取租户需求分布
func (h *AnalyticsHandler) GetTenantDemandDistribution(c fiber.Ctx) error {
	// 根据带看记录中的反馈分析需求
	type DemandDistribution struct {
		DemandType string  `json:"demand_type"`
		Count      int64   `json:"count"`
		Percentage float64 `json:"percentage"`
	}

	var distributions []DemandDistribution
	h.db.Raw(`
		SELECT 
			CASE 
				WHEN feedback LIKE '%价格%' THEN '价格' 
				WHEN feedback LIKE '%户型%' THEN '户型' 
				WHEN feedback LIKE '%位置%' THEN '位置' 
				WHEN feedback LIKE '%装修%' THEN '装修' 
				ELSE '其他' 
			END as demand_type, 
			COUNT(*) as count
		FROM hrs_property_viewing 
		WHERE feedback IS NOT NULL
		GROUP BY demand_type
	`).Scan(&distributions)

	// 计算总记录数
	var totalRecords int64
	h.db.Model(&struct{}{}).Table("hrs_property_viewing").Where("feedback IS NOT NULL").Count(&totalRecords)

	// 计算百分比
	for i := range distributions {
		if totalRecords > 0 {
			distributions[i].Percentage = float64(distributions[i].Count) / float64(totalRecords) * 100
		}
	}

	return c.JSON(distributions)
}

// GetIncomeExpenseTrend 获取收支趋势
func (h *AnalyticsHandler) GetIncomeExpenseTrend(c fiber.Ctx) error {
	// 获取最近6个月的收支趋势
	endDate := time.Now()
	startDate := endDate.AddDate(0, -6, 0)

	type MonthlyFinance struct {
		Month     string  `json:"month"`
		Income    float64 `json:"income"`
		Expense   float64 `json:"expense"`
		NetIncome float64 `json:"net_income"`
	}

	var monthlyFinances []MonthlyFinance
	h.db.Raw(`
		SELECT 
			DATE_FORMAT(p.payment_time, '%Y-%m') as month, 
			SUM(p.amount) as income, 
			0 as expense,
			SUM(p.amount) as net_income
		FROM hrs_payment p
		WHERE p.payment_time BETWEEN ? AND ?
		GROUP BY month
		ORDER BY month
	`, startDate, endDate).Scan(&monthlyFinances)

	return c.JSON(monthlyFinances)
}

// GetCashFlow 获取现金流
func (h *AnalyticsHandler) GetCashFlow(c fiber.Ctx) error {
	// 获取最近30天的现金流
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -30)

	type DailyCashFlow struct {
		Date      string  `json:"date"`
		Income    float64 `json:"income"`
		Expense   float64 `json:"expense"`
		NetAmount float64 `json:"net_amount"`
	}

	var dailyCashFlows []DailyCashFlow
	h.db.Raw(`
		SELECT 
			DATE(p.payment_time) as date, 
			SUM(p.amount) as income, 
			0 as expense,
			SUM(p.amount) as net_amount
		FROM hrs_payment p
		WHERE p.payment_time BETWEEN ? AND ?
		GROUP BY date
		ORDER BY date
	`, startDate, endDate).Scan(&dailyCashFlows)

	return c.JSON(dailyCashFlows)
}

// GetViewingConversionRate 获取带看转化率
func (h *AnalyticsHandler) GetViewingConversionRate(c fiber.Ctx) error {
	// 计算带看转化率 = 带看后签约的数量 / 总带看数量
	type ConversionRate struct {
		Period        string  `json:"period"`
		TotalViewings int64   `json:"total_viewings"`
		Conversions   int64   `json:"conversions"`
		Rate          float64 `json:"rate"`
	}

	var rates []ConversionRate
	h.db.Raw(`
		SELECT 
			DATE_FORMAT(v.create_time, '%Y-%m') as period, 
			COUNT(v.id) as total_viewings, 
			COUNT(DISTINCT c.id) as conversions
		FROM hrs_property_viewing v
		LEFT JOIN hrs_contract c ON v.property_id = c.property_id AND v.tenant_id = c.tenant_id
		WHERE v.create_time BETWEEN DATE_SUB(NOW(), INTERVAL 6 MONTH) AND NOW()
		GROUP BY period
		ORDER BY period
	`).Scan(&rates)

	// 计算转化率
	for i := range rates {
		if rates[i].TotalViewings > 0 {
			rates[i].Rate = float64(rates[i].Conversions) / float64(rates[i].TotalViewings) * 100
		}
	}

	return c.JSON(rates)
}

// GetOperationPerformance 获取运营绩效
func (h *AnalyticsHandler) GetOperationPerformance(c fiber.Ctx) error {
	// 计算经纪人绩效
	type AgentPerformance struct {
		AgentID          uint    `json:"agent_id"`
		AgentName        string  `json:"agent_name"` // 实际应从用户表获取
		TotalViewings    int64   `json:"total_viewings"`
		TotalContracts   int64   `json:"total_contracts"`
		TotalIncome      float64 `json:"total_income"`
		PerformanceScore float64 `json:"performance_score"`
	}

	var performances []AgentPerformance
	h.db.Raw(`
		SELECT 
			v.agent_id, 
			'Agent ' || v.agent_id as agent_name, 
			COUNT(v.id) as total_viewings, 
			COUNT(DISTINCT c.id) as total_contracts,
			SUM(p.amount) as total_income
		FROM hrs_property_viewing v
		LEFT JOIN hrs_contract c ON v.property_id = c.property_id
		LEFT JOIN hrs_payment p ON c.id = p.bill_id
		GROUP BY v.agent_id
		ORDER BY total_income DESC
		LIMIT 10
	`).Scan(&performances)

	// 计算绩效分数
	for i := range performances {
		// 简单的绩效计算逻辑，实际应根据业务需求调整
		performances[i].PerformanceScore = float64(performances[i].TotalViewings)*0.3 +
			float64(performances[i].TotalContracts)*0.5 +
			performances[i].TotalIncome*0.2
	}

	return c.JSON(performances)
}
