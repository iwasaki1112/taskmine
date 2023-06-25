# Taskmine

## Endpoint
POST http://localhost:8080/task HTTP/1.1
content-type: application/json

{
    "title": "sample",
    "description": "Wed, 21 Oct 2015 18:27:50 GMT"
}

PATCH http://localhost:8080/task HTTP/1.1
content-type: application/json

{
    "id": "2",
		"title": "tomijriou",
    "description": "ojisan"
}

## Directories
/Taskmine
	/cmd
		/main
			- main.go
	/domain		
		/repository
			- task_repository.go
		/entity
			- task.go
    	- task_status.go
	/application		
		- task_interactor.go (エンティティの生成、更新、保存等)
		- task_input.go (タイトルと説明を受け取るためのデータ構造を定義)
		- task_output.go (作成されたタスクの詳細を提供するためのデータ構造を定義)
	/config
		- config.go
	/interface (外部との通信やデータの入出力)
		/http (httpリクエストとレスポンスをハンドリング)			
			- task_handler.go	
	/infrastructure	
		- mysql_task_repository.go (mysqlへのアクセス)				