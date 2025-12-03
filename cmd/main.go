package main

import (
	"aivisual-core/internal/config"
	"aivisual-core/internal/infra/db"
	"aivisual-core/internal/infra/db/repositories"
	"aivisual-core/internal/infra/http/routes"
	"aivisual-core/internal/infra/rabbitmq"
	"aivisual-core/internal/service"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("../configs")
	if err != nil {
		log.Fatalf("无法加载配置: %v", err)
	}

	// 初始化数据库连接
	database, err := db.NewDB(&cfg.Database)
	if err != nil {
		log.Fatalf("无法连接数据库: %v", err)
	}
	defer database.Close()

	// 运行数据库迁移
	if err := database.RunMigrations(); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	// 初始化数据仓库
	alarmRepo := repositories.NewAlarmRepository(database)
	eventRepo := repositories.NewEventRepository(database)

	// 初始化服务
	converter := service.NewConverter()
	alarmService := service.NewAlarmService(converter)

	// 初始化RabbitMQ消费者
	consumer, err := rabbitmq.NewConsumer(cfg.RabbitMQ.URL)
	if err != nil {
		log.Fatalf("无法创建RabbitMQ消费者: %v", err)
	}
	defer consumer.Close()

	// 启动RabbitMQ消息监听
	go func() {
		err := consumer.Consume("alarm_queue", func(data []byte) {
			alarm, err := alarmService.ProcessAlarmMessage(data)
			if err != nil {
				log.Printf("处理报警消息失败: %v", err)
				return
			}
			
			// 存储到数据库
			if err := alarmRepo.Create(alarm); err != nil {
				log.Printf("存储报警消息失败: %v", err)
				return
			}
			
			log.Printf("收到并存储报警消息: %+v", alarm)
		})
		if err != nil {
			log.Printf("RabbitMQ消费失败: %v", err)
		}
	}()

	// 初始化HTTP服务
	r := gin.Default()

	// 注册路由
	alarmHandler := routes.NewAlarmHandler(alarmRepo)
	eventHandler := routes.NewEventHandler(eventRepo)
	detectionHandler := routes.NewDetectionHandler()
	wvpHandler := routes.NewWVPHandler(converter, alarmRepo, eventRepo)

	// 报警相关接口
	r.GET("/api/alarms", alarmHandler.GetAlarms)
	r.GET("/api/alarms/:id", alarmHandler.GetAlarmByID)
	
	// 事件相关接口
	r.GET("/api/events", eventHandler.GetEvents)
	
	// 检测相关接口
	r.GET("/api/detections", detectionHandler.GetDetections)

	// WVP-Pro设备事件接口
	r.POST("/api/webhook/wvp", wvpHandler.HandleWVPEvent)

	// 启动HTTP服务
	go func() {
		log.Printf("HTTP服务器启动，监听端口 %d", cfg.Server.Port)
		if err := r.Run(":" + strconv.Itoa(cfg.Server.Port)); err != nil {
			log.Fatalf("HTTP服务器启动失败: %v", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("正在关闭服务器...")
}