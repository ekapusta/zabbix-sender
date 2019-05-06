# Zabbix sender

Утилита для отправки значения метрик в zabbix-сервер. Имя сервера берётся из конфигурации zabbix-агента zabbix_agentd.conf 

## Установка и запуск

```shell
go install github.com/ekapusta/zabbix-sender
zabbix-sender -p test_metric -v 1 -с /etc/zabbix/zabbix_agentd.conf
```

## Параметры запуска

```
Usage of ./zabbix-sender:
  -c string
    	Zabbix config file (default "/etc/zabbix/zabbix_agentd.conf")
  -p string
    	Parameter (default "test")
  -v string
    	Parameter value
```