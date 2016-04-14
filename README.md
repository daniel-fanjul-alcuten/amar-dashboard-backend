# Compilation

    go get github.com/daniel-fanjul-alcuten/amar-dashboard-backend/amar-fetch
    go get github.com/daniel-fanjul-alcuten/amar-dashboard-backend/amar-save

# Execution

    amar-fetch -u "uid cookie" -p "pid cookie" > fetch.json
    amar-save -d mysql -h user:password@tcp(host:port)/dbname < fetch.json

or simply

    amar-fetch -u "uid cookie" -p "pid cookie" | amar-save -d mysql -h user:password@tcp(host:port)/dbname
