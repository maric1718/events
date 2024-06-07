1. Run "docker compose up -d" to start my_app, kafka, zookeeper and kafka-ui

2. Open http://localhost:29093/ui for kafka UI

3. App will get initial data from internal/adapter/storage/file_system/json_data/inital and store it in application state

4. Data can be updated by sending event and market json files to their kafka topics

5. Client can access this data calling v1/events GET and v1/markets GET routes 
    v1/events filters data by receiving "to" and "from" query parameters

6. Old data will be periodically removed using cronjob