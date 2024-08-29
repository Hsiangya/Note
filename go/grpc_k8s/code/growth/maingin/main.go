package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"growth/conf"
	"growth/dbhelper"
	"growth/pb"
	"growth/ugserver"
	"log"
	"net/http"
	"time"
)

func initDb() {
	// default UTC time location
	time.Local = time.UTC
	// Load global config
	conf.LoadConfigs()
	// Initialize db
	dbhelper.InitDb()
}

var AllowOrigin = map[string]bool{
	"http://a.site.com": true,
	"http://b.site.com": true,
	"http://web.com":    true,
}

func mainGin() {
	// 连接到grpc服务的客户端
	//conn, err := grpc.Dial("localhost:7789", grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:7789", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	clientCoin := pb.NewUserCoinClient(conn)
	clientGrade := pb.NewUserGradeClient(conn)

	router := gin.New()
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello")
	})
	// 用户积分服务的方法
	v1Group := router.Group("/v1", func(ctx *gin.Context) {

		// 支持跨域
		origin := ctx.GetHeader("Origin")
		if AllowOrigin[origin] {
			ctx.Header("Access-Control-Allow-Origin", origin)
			ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTION")
			ctx.Header("Access-Control-Allow-Headers", "*")
			ctx.Header("Access-Control-Allow-Credentials", "true")
		}
		ctx.Next()
	})
	gUserCoin := v1Group.Group("/UserGrowth.UserCoin")
	gUserCoin.GET("/ListTasks", func(ctx *gin.Context) {
		out, err := clientCoin.ListTasks(ctx, &pb.ListTasksRequest{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    2,
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, out)
		}
	})
	gUserCoin.POST("/UserCoinChange", func(ctx *gin.Context) {
		body := &pb.UserCoinChangeRequest{}
		err := ctx.BindJSON(body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    2,
				"message": err.Error(),
			})
		} else if out, err := clientCoin.UserCoinChange(ctx, body); err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    2,
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, out)
		}
		ctx.JSON(http.StatusOK, nil)
	})

	// 用户等级服务的方法
	gUserGrade := v1Group.Group("/UserGrowth.UserGrade")
	gUserGrade.GET("/ListGrades", func(ctx *gin.Context) {
		out, err := clientGrade.ListGrades(ctx, &pb.ListGradesRequest{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    2,
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, out)
		}
	})

	// 为http/2配置参数
	h2Handler := h2c.NewHandler(router, &http2.Server{})
	// 配置http服务
	server := &http.Server{
		Addr:    ":8080",
		Handler: h2Handler,
	}
	// 启动http服务
	server.ListenAndServe()
}

func mainGateWay() {
	initDb()
	s := grpc.NewServer()
	// 注册GRPC服务
	pb.RegisterUserCoinServer(s, &ugserver.UgCoinServer{})
	pb.RegisterUserGradeServer(s, &ugserver.UgGradeServer{})

	// 注册gateway服务
	mux := runtime.NewServeMux()
	ctx := context.Background()
	if err := pb.RegisterUserCoinHandlerServer(ctx, mux, &ugserver.UgCoinServer{}); err != nil {
		log.Printf("Fail to RegisterUserCoinHandlerServer error=%v", err)
	}
	if err := pb.RegisterUserGradeHandlerServer(ctx, mux, &ugserver.UgGradeServer{}); err != nil {
		log.Printf("Fail to RegisterUserGradeHandlerServer error=%v", err)
	}

	// 定义gateway对象
	httpMux := http.NewServeMux()
	httpMux.Handle("/v1/UserGrowth", mux) // 地址需要与pb文件中一致
	server := &http.Server{
		Addr: ":8081",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("http.HandleFunc url=%s", r.URL)
			mux.ServeHTTP(w, r)
		}),
	}

	// 启动gateway服务
	log.Printf("Server.ListenAndServe(%s)", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("ListenAndServe error=%v", err)
	}
}

func main() {
	go mainGateWay()
	mainGin()
}
