package data

import "flag"

var dataSourceInstance *string

func GetDataSourceInstance() *string {
    if dataSourceInstance == nil {
        dataSourceInstance = flag.String("datasource", "postgres", 
                                "Data source: postgres (default) or mock")
        flag.Parse()
    }
    return dataSourceInstance
}

