package service

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"dbm-services/common/go-pubpkg/errno"
	"dbm-services/mysql/priv-service/util"
)

const mysql string = "mysql" // 包含tendbha和tendbsingle
const tendbha string = "tendbha"
const tendbsingle string = "tendbsingle"
const tendbcluster string = "tendbcluster"
const machineTypeBackend string = "backend"
const machineTypeSingle string = "single"
const machineTypeRemote string = "remote"
const machineTypeProxy string = "proxy"
const machineTypeSpider string = "spider"
const backendSlave string = "backend_slave"
const masterEntry = "master_entry"
const slaveEntry = "slave_entry"
const running string = "running"
const tdbctl string = "tdbctl"
const sqlserver string = "sqlserver"
const sqlserverHA string = "sqlserver_ha"
const sqlserverSingle string = "sqlserver_single"
const backendMaster string = "backend_master"
const orphan string = "orphan"
const sqlserverSysDB string = "Monitor"
const mongodb string = "mongodb"

// GetAllClustersInfo 获取业务下所有集群信息
/*
	[{
		  "db_module_id": 126,
		  "bk_biz_id": "3",
		  "cluster_type": "tendbsingle",
		  "proxies": [],
		  "storages": [
		    {
		      "ip": "1.1.1.1.",
		      "instance_role": "orphan",
		      "port": 30000
		    }
		  ],
		  "immute_domain": "singledb.1.hayley.db"
		},
		{
		  "db_module_id": 500,
		  "bk_biz_id": "3",
		  "cluster_type": "tendbha",
		  "proxies": [
		    {
		      "ip": "1.1.1.1",
		      "admin_port": 41000,
		      "port": 40000
		    },
		    {
		      "ip": "2.2.2.2",
		      "admin_port": 41000,
		      "port": 40000
		    }
		  ],
		  "storages": [
		    {
		      "ip": "127.0.0.3",
		      "instance_role": "backend_slave",
		      "port": 30000
		    },
		    {
		      "ip": "127.0.0.4",
		      "instance_role": "backend_master",
		      "port": 40000
		    }
		  ],
		  "immute_domain": "gamedb.2.hayley.db"
		}]
*/
func GetAllClustersInfo(id BkBizIdPara) ([]Cluster, error) {
	var resp []Cluster
	url := "/apis/proxypass/dbmeta/priv_manager/biz_clusters/"
	result, err := util.DbmetaClient.Do(http.MethodPost, url, id)
	if err != nil {
		slog.Error("msg", url, err)
		return resp, err
	}
	if err = json.Unmarshal(result.Data, &resp); err != nil {
		slog.Error("msg", url, err)
		return resp, err
	}
	return resp, nil
}

// GetCluster 根据域名获取集群信息
func GetCluster(ClusterType string, dns Domain) (Instance, error) {
	var resp Instance
	var url string
	if ClusterType == sqlserverHA || ClusterType == sqlserverSingle || ClusterType == sqlserver {
		// 走sqlserver授权逻辑
		url = fmt.Sprintf("/apis/proxypass/dbmeta/priv_manager/sqlserver/%s/cluster_instances/", ClusterType)
	} else {
		url = fmt.Sprintf("/apis/proxypass/dbmeta/priv_manager/mysql/%s/cluster_instances/", ClusterType)
	}

	result, err := util.DbmetaClient.Do(http.MethodPost, url, dns)
	if err != nil {
		slog.Error("msg", url, err)
		return resp, errno.DomainNotExists.Add(fmt.Sprintf(" %s: %s", dns.EntryName, err.Error()))
	}
	if err = json.Unmarshal(result.Data, &resp); err != nil {
		slog.Error("msg", url, err)
		return resp, err
	}
	return resp, nil
}
