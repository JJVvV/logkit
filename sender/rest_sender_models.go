package sender

import "github.com/qiniu/logkit/utils"

// ModeUsages 用途说明
var ModeUsages = []utils.KeyValue{
	{TypePandora, "发送到 Pandora"},
	{TypeFile, "发送到本地文件"},
	{TypeMongodbAccumulate, "发送到 mongodb"},
	{TypeInfluxdb, "发送到 influxdb"},
	{TypeDiscard, "消费数据但不发送"},
	{TypeElastic, "发送到Elasticsearch"},
	{TypeKafka, "发送到Kafka"},
}

var (
	OptionSaveLogPath = utils.Option{
		KeyName:      KeyFtSaveLogPath,
		ChooseOnly:   false,
		Default:      "",
		DefaultNoUse: false,
		Description:  "管道本地盘数据保存路径(ft_save_log_path)",
	}
	OptionFtWriteLimit = utils.Option{
		KeyName:      KeyFtWriteLimit,
		ChooseOnly:   false,
		Default:      "",
		DefaultNoUse: false,
		Description:  "磁盘写入限速(MB/s)(ft_write_limit)",
		CheckRegex:   "\\d+",
	}
	OptionFtSyncEvery = utils.Option{
		KeyName:      KeyFtSyncEvery,
		ChooseOnly:   false,
		Default:      "",
		DefaultNoUse: false,
		Description:  "同步meta的间隔(ft_sync_every)",
		CheckRegex:   "\\d+",
	}
	OptionFtStrategy = utils.Option{
		KeyName:       KeyFtStrategy,
		ChooseOnly:    true,
		ChooseOptions: []string{KeyFtStrategyBackupOnly, KeyFtStrategyAlwaysSave, KeyFtStrategyConcurrent},
		Default:       KeyFtStrategyBackupOnly,
		DefaultNoUse:  false,
		Description:   "磁盘管道容错策略(仅备份错误|全部数据走管道)(ft_strategy)",
	}
	OptionFtProcs = utils.Option{
		KeyName:      KeyFtProcs,
		ChooseOnly:   false,
		Default:      "",
		DefaultNoUse: false,
		Description:  "发送并发数量(磁盘管道或内存管道 always_save 或 concurrent 模式生效)(ft_procs)",
		CheckRegex:   "\\d+",
	}
	OptionFtMemoryChannel = utils.Option{
		KeyName:       KeyFtMemoryChannel,
		ChooseOnly:    true,
		ChooseOptions: []string{"false", "true"},
		Default:       "false",
		DefaultNoUse:  false,
		Description:   "使用内存替换磁盘管道(加速)(ft_memory_channel)",
	}
	OptionFtMemoryChannelSize = utils.Option{
		KeyName:      KeyFtMemoryChannelSize,
		ChooseOnly:   false,
		Default:      "",
		DefaultNoUse: false,
		Description:  "内存管道长度(ft_memory_channel_size)",
		CheckRegex:   "\\d+",
	}
)
var ModeKeyOptions = map[string][]utils.Option{
	TypeFile: {
		{
			KeyName:      KeyFileSenderPath,
			ChooseOnly:   false,
			Default:      "/home/john/mylogs/my.log",
			DefaultNoUse: true,
			Description:  "发送到的目的文件路径(file_send_path)",
		},
	},
	TypePandora: {
		{
			KeyName:      KeyPandoraRepoName,
			ChooseOnly:   false,
			Default:      "my_work",
			DefaultNoUse: true,
			Description:  "Pandora 数据源名称(pandora_repo_name)",
			CheckRegex:   "^[a-zA-Z][a-zA-Z0-9_]{0,127}$",
		},
		{
			KeyName:      KeyPandoraAk,
			ChooseOnly:   false,
			Default:      "在此填写您七牛账号ak(access_key)",
			DefaultNoUse: true,
			Description:  "七牛的公钥(access_key)",
		},
		{
			KeyName:      KeyPandoraSk,
			ChooseOnly:   false,
			Default:      "在此填写您七牛账号的secret_key",
			DefaultNoUse: true,
			Description:  "七牛的私钥(secret_key)",
		},
		OptionSaveLogPath,
		{
			KeyName:       KeyLogkitSendTime,
			ChooseOnly:    true,
			ChooseOptions: []string{"true", "false"},
			Default:       "true",
			DefaultNoUse:  false,
			Description:   "是否在发送数据时自动添加发送时间(logkit_send_time)",
		},
		{
			KeyName:      KeyPandoraHost,
			ChooseOnly:   false,
			Default:      "https://pipeline.qiniu.com",
			DefaultNoUse: false,
			Description:  "Host地址(pandora_host)",
		},
		{
			KeyName:       KeyPandoraRegion,
			ChooseOnly:    true,
			ChooseOptions: []string{"nb"},
			Default:       "nb",
			DefaultNoUse:  false,
			Description:   "创建的资源所在区域(pandora_region)",
		},
		{
			KeyName:       KeyPandoraSchemaFree,
			ChooseOnly:    true,
			ChooseOptions: []string{"true", "false"},
			Default:       "true",
			DefaultNoUse:  false,
			Description:   "是否根据数据自动创建与增加字段(pandora_schema_free)",
		},
		{
			KeyName:      KeyPandoraAutoCreate,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "以DSL语法自动创建repo(pandora_auto_create)",
		},
		{
			KeyName:      KeyPandoraSchema,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "仅选择部分字段(重命名)发送(pandora_schema)",
		},
		{
			KeyName:       KeyPandoraEnableLogDB,
			ChooseOnly:    true,
			ChooseOptions: []string{"true", "false"},
			Default:       "true",
			DefaultNoUse:  false,
			Description:   "是否自动创建并导出到Pandora LogDB(pandora_enable_logdb)",
		},
		{
			KeyName:      KeyPandoraLogDBName,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "导出的 LogDB 仓库名称(pandora_logdb_name)",
		},
		{
			KeyName:      KeyPandoraLogDBHost,
			ChooseOnly:   false,
			Default:      "https://logdb.qiniu.com",
			DefaultNoUse: false,
			Description:  "LogDB host 地址(pandora_logdb_host)",
		},

		{
			KeyName:       KeyPandoraEnableKodo,
			ChooseOnly:    true,
			ChooseOptions: []string{"false", "true"},
			Default:       "false",
			DefaultNoUse:  false,
			Description:   "是否自动导出到七牛云存储(pandora_enable_kodo)",
		},
		{
			KeyName:      KeyPandoraKodoBucketName,
			ChooseOnly:   false,
			Default:      "my_bucket_name",
			DefaultNoUse: true,
			Description:  "云存储 Bucket 仓库名称(启用自动导出到云存储时必填)(pandora_bucket_name)",
		},
		{
			KeyName:      KeyPandoraEmail,
			ChooseOnly:   false,
			Default:      "my@email.com",
			DefaultNoUse: true,
			Description:  "邮箱(启用自动导出到云存储时必填)(qiniu_email)",
		},

		{
			KeyName:       KeyPandoraGzip,
			ChooseOnly:    true,
			ChooseOptions: []string{"true", "false"},
			Default:       "true",
			DefaultNoUse:  false,
			Description:   "gzip压缩发送(pandora_gzip)",
		},
		{
			KeyName:      KeyFlowRateLimit,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "流量限制(KB/s)(flow_rate_limit)",
			CheckRegex:   "\\d+",
		},
		{
			KeyName:      KeyRequestRateLimit,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "请求限制(次/s)(request_rate_limit)",
			CheckRegex:   "\\d+",
		},
		{
			KeyName:       KeyPandoraUUID,
			ChooseOnly:    true,
			ChooseOptions: []string{"false", "true"},
			Default:       "false",
			DefaultNoUse:  false,
			Description:   "每条数据植入UUID(pandora_uuid)",
		},
		{
			KeyName:       KeyPandoraWithIP,
			ChooseOnly:    false,
			ChooseOptions: []string{"false", "true"},
			Default:       "false",
			DefaultNoUse:  false,
			Description:   "每条数据植入IP地址(pandora_withip)",
		},
		OptionFtWriteLimit,
		OptionFtSyncEvery,
		OptionFtStrategy,
		OptionFtProcs,
		OptionFtMemoryChannel,
		OptionFtMemoryChannelSize,
		{
			KeyName:       KeyForceMicrosecond,
			ChooseOnly:    true,
			ChooseOptions: []string{"false", "true"},
			Default:       "false",
			DefaultNoUse:  false,
			Description:   "对于数据的时间字段抖动(force_microsecond)",
		},
		{
			KeyName:       KeyForceDataConvert,
			ChooseOnly:    true,
			ChooseOptions: []string{"false", "true"},
			Default:       "false",
			DefaultNoUse:  false,
			Description:   "数据强制类型转换(pandora_force_convert)",
		},
		{
			KeyName:       KeyIgnoreInvalidField,
			ChooseOnly:    true,
			ChooseOptions: []string{"true", "false"},
			Default:       "true",
			DefaultNoUse:  false,
			Description:   "忽略格式错误的字段(ignore_invalid_field)",
		},
		{
			KeyName:       KeyPandoraAutoConvertDate,
			ChooseOnly:    true,
			ChooseOptions: []string{"true", "false"},
			Default:       "true",
			DefaultNoUse:  false,
			Description:   "时间类型自动转换(pandora_auto_convert_date)",
		},
	},
	TypeMongodbAccumulate: {
		{
			KeyName:      KeyMongodbHost,
			ChooseOnly:   false,
			Default:      "mongodb://[username:password@]host1[:port1][,host2[:port2],...[,hostN[:portN]]][/[database][?options]]",
			DefaultNoUse: true,
			Description:  "数据库地址(mongodb_host)",
		},
		{
			KeyName:      KeyMongodbDB,
			ChooseOnly:   false,
			Default:      "app123",
			DefaultNoUse: true,
			Description:  "数据库名称(mongodb_db)",
		},
		{
			KeyName:      KeyMongodbCollection,
			ChooseOnly:   false,
			Default:      "collection1",
			DefaultNoUse: true,
			Description:  "数据表名称(mongodb_collection)",
		},
		{
			KeyName:      KeyMongodbUpdateKey,
			ChooseOnly:   false,
			Default:      "domain,uid",
			DefaultNoUse: true,
			Description:  "聚合条件列(mongodb_acc_updkey)",
		},
		{
			KeyName:      KeyMongodbAccKey,
			ChooseOnly:   false,
			Default:      "low,hit",
			DefaultNoUse: true,
			Description:  "聚合列(mongodb_acc_acckey)",
		},
		OptionSaveLogPath,
		OptionFtWriteLimit,
		OptionFtSyncEvery,
		OptionFtStrategy,
		OptionFtProcs,
		OptionFtMemoryChannel,
		OptionFtMemoryChannelSize,
	},
	TypeInfluxdb: {
		{
			KeyName:      KeyInfluxdbHost,
			ChooseOnly:   false,
			Default:      "127.0.0.1:8086",
			DefaultNoUse: true,
			Description:  "数据库地址(influxdb_host)",
		},
		{
			KeyName:      KeyInfluxdbDB,
			ChooseOnly:   false,
			Default:      "testdb",
			DefaultNoUse: true,
			Description:  "数据库名称(influxdb_db)",
		},
		{
			KeyName:      KeyInfluxdbMeasurement,
			ChooseOnly:   false,
			Default:      "test_table",
			DefaultNoUse: true,
			Description:  "measurement名称(influxdb_measurement)",
		},
		{
			KeyName:      KeyInfluxdbRetetion,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "retention名称(influxdb_retention)",
		},
		{
			KeyName:      KeyInfluxdbTags,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "标签列数据(influxdb_tags)",
		},
		{
			KeyName:      KeyInfluxdbFields,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "普通列数据(influxdb_fields)",
		},
		{
			KeyName:      KeyInfluxdbTimestamp,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "时间戳列(influxdb_timestamp)",
		},
		{
			KeyName:      KeyInfluxdbTimestampPrecision,
			ChooseOnly:   false,
			Default:      "100",
			DefaultNoUse: false,
			Description:  "时间戳列精度调整(influxdb_timestamp_precision)",
		},
		OptionSaveLogPath,
		OptionFtWriteLimit,
		OptionFtSyncEvery,
		OptionFtStrategy,
		OptionFtProcs,
		OptionFtMemoryChannel,
		OptionFtMemoryChannelSize,
	},
	TypeDiscard: {},
	TypeElastic: {
		{
			KeyName:      KeyElasticHost,
			ChooseOnly:   false,
			Default:      "192.168.31.203:9200",
			DefaultNoUse: false,
			Description:  "host地址(elastic_host)",
		},
		{
			KeyName:      KeyElasticIndex,
			ChooseOnly:   false,
			Default:      "app-repo-123",
			DefaultNoUse: true,
			Description:  "索引名称(elastic_index)",
		},
		{
			KeyName:       KeyElasticIndexStrategy,
			ChooseOnly:    true,
			ChooseOptions: []string{KeyDefaultIndexStrategy, KeyYearIndexStrategy, KeyMonthIndexStrategy, KeyDayIndexStrategy},
			Default:       KeyFtStrategyBackupOnly,
			DefaultNoUse:  false,
			Description:   "自动索引模式(默认索引|按年索引|按月索引|按日索引)(index_strategy)",
		},
		{
			KeyName:      KeyElasticType,
			ChooseOnly:   false,
			Default:      "app",
			DefaultNoUse: true,
			Description:  "索引类型名称(elastic_type)",
		},
		OptionSaveLogPath,
		OptionFtWriteLimit,
		OptionFtSyncEvery,
		OptionFtStrategy,
		OptionFtProcs,
		OptionFtMemoryChannel,
		OptionFtMemoryChannelSize,
	},
	TypeKafka: {
		{
			KeyName:      KeyKafkaHost,
			ChooseOnly:   false,
			Default:      "192.168.31.201:9092",
			DefaultNoUse: true,
			Description:  "broker的host地址(kafka_host)",
		},
		{
			KeyName:      KeyKafkaTopic,
			ChooseOnly:   false,
			Default:      "my_topic",
			DefaultNoUse: true,
			Description:  "打点的topic名称(kafka_topic)",
		},
		{
			KeyName:       KeyKafkaCompression,
			ChooseOnly:    true,
			ChooseOptions: []string{KeyKafkaCompressionNone, KeyKafkaCompressionGzip, KeyKafkaCompressionSnappy},
			Default:       KeyKafkaCompressionNone,
			DefaultNoUse:  false,
			Description:   "压缩模式(none不压缩|gzip压缩|snappy压缩)(kafka_compression)",
		},
		{
			KeyName:      KeyKafkaClientId,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "kafka客户端标识ID(kafka_client_id)",
		},
		{
			KeyName:      KeyKafkaRetryMax,
			ChooseOnly:   false,
			Default:      "3",
			DefaultNoUse: false,
			Description:  "kafka最大错误重试次数(kafka_retry_max)",
		},
		{
			KeyName:      KeyKafkaTimeout,
			ChooseOnly:   false,
			Default:      "30s",
			DefaultNoUse: false,
			Description:  "kafka连接超时时间(kafka_timeout)",
		},
		{
			KeyName:      KeyKafkaKeepAlive,
			ChooseOnly:   false,
			Default:      "0",
			DefaultNoUse: false,
			Description:  "kafka的keepalive时间(kafka_keep_alive)",
		},
		OptionSaveLogPath,
		OptionFtWriteLimit,
		OptionFtSyncEvery,
		OptionFtStrategy,
		OptionFtProcs,
		OptionFtMemoryChannel,
		OptionFtMemoryChannelSize,
	},
}
