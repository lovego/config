module github.com/lovego/config

go 1.16

//replace github.com/lovego/fs => ../fs
require (
	github.com/garyburd/redigo v1.6.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gocql/gocql v0.0.0-20210515062232-b7ef815b4556
	github.com/lib/pq v1.10.2
	github.com/lovego/alarm v0.0.6
	github.com/lovego/bsql v0.0.3
	github.com/lovego/config_sdk v0.0.0-20211229081009-5e3148dda5bd // indirect
	github.com/lovego/config_sdk/go_config_sdk v0.0.0-20211229081009-5e3148dda5bd
	github.com/lovego/duration v0.0.0-20200802140436-42c773e4fb38
	github.com/lovego/email v0.0.5
	github.com/lovego/fs v0.0.4
	github.com/lovego/logger v0.0.3
	github.com/lovego/strmap v0.0.0-20211228093343-09d28c495476
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b

)
