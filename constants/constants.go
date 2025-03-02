package constants

const (
	BaseConfigPathUATPtrms = "resources/uat"
)

const BaseConfigPathKey string = "base-config-path"
const BaseConfigPathDefaultValue string = "resources"
const BaseConfigPathUsage string = "path to folder that stores your configurations"

const (
	MySQL_DB_Config = "mysql_db"

	//Margin DB
	MySQL_DB_ConfigServerConfigKey             = "margindb.server"
	MySQL_DB_ConfigPortConfigKey               = "margindb.port"
	MySQL_DB_ConfigDBNameConfigKey             = "margindb.name"
	MySQL_DB_ConfigUsernameConfigKey           = "margindb.username"
	MySQL_DB_ConfigPasswordConfigKey           = "margindb.password"
	MySQL_DB_ConfigTableNameConfigKey          = "margindb.table"
	MySQL_DB_ConfigMaxOpenConnectionsKey       = "margindb.maxOpenConnections"
	MySQL_DB_ConfigMaxIdleConnectionsKey       = "margindb.maxIdleConnections"
	MySQL_DB_ConfigTransactionTimeoutInSeconds = "margindb.txnTimeOutInSeconds"
	MySQL_DB_ConfigLifetimeInSecondsKey        = "margindb.connectionMaxLifetimeInSeconds"
	MySQL_DB_ConfignMaxIdleTimeInSecondsKey    = "margindb.connectionMaxIdleTimeInSeconds"
	MySQL_DB_ConfigJobStatusTableNameConfigKey = "margindb.jobStatusTable"
	//DB Retry
	MySQL_DB_ConfigMaxRetryCount = "margindb.maxRetryCount"
)
