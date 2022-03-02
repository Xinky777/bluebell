package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局变量,用来保存程序的所有配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Post         int    `mapstructure:"post"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	//方法1：直接指定配置文件路径（相对路径或者绝对路径）
	//指定具体配置文件用该函数 既包含文件名 还包含配置名
	//相对路径：相对执行的可执行文件的相对路径
	//viper.SetConfigFile("./config.yaml")
	//绝对路径：系统中实际的文件路径
	//viper.SetConfigFile("/Users/xinky/go/src/web_app/config.yaml")

	//方法2：指定配置文件名和配置文件的位置，viper自行查找可用的配置文件
	//配置文件名不需要带后缀
	//配置文件的位置可配置多个
	viper.SetConfigName("config") //指定配置文件名称（不带后缀）
	viper.AddConfigPath(".")      //指定查找配置文件的路径（此处使用相对路径）

	//配合远程配置中心使用 告诉viper当前的数据使用什么格式去解析
	//	viper.SetConfigType("yaml")   //指定配置文件类型 如果是JSON格式配置文件，只需要将"yaml"改为"json"即可

	if err = viper.ReadInConfig(); err != nil { //读取配置信息
		//读取配置文件失败
		fmt.Printf("ReadInConfig failed，err：%v\n", err)
		return
	}
	//把读取到的配置信息反序列化到 Conf 变量中
	if err1 := viper.Unmarshal(Conf); err1 != nil {
		fmt.Printf("viper.Unmarshal failed,err:%v\n", err1)
	}
	//监控文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已修改")
		if err1 := viper.Unmarshal(Conf); err1 != nil {
			fmt.Printf("viper.Unmarshal failed,err:%v\n", err1)
		}
	})
	return
}
