host="http://52.67.218.214"

# Health check
curl "$host/health-check" -v

# New game
curl -X POST "$host/game" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"cols\": 5, \"rows\": 5, \"mines\": 8, \"username\": \"myuser\"}" -v

gameid="58e9745e-4874-4449-a39b-a49d2e61bc84"

# Flag
curl -X PUT "$host/game/$gameid/flag" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"col\": 0, \"row\": 0, \"flag\": \"red_flag\", \"username\": \"myuser\"}" -v

# Explore
curl -X POST "$host/game/$gameid/explore" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"col\": 0, \"row\": 0, \"username\": \"myuser\"}" -v

# Resume game
curl -X POST "$host/game/$gameid/resume" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"username\": \"myuser\"}" -v
