# New game
curl -X POST "http://18.228.204.155/game" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"cols\": 5, \"rows\": 5, \"mines\": 8, \"username\": \"myuser\"}" -v

# 112783cd-66bd-4a85-9445-1ca293c41c25 is a valid id, and is used for the rest of the tests

# Flag
curl -X PUT "http://18.228.204.155/game/112783cd-66bd-4a85-9445-1ca293c41c25/flag" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"col\": 0, \"row\": 0, \"flag\": \"red_flag\", \"username\": \"myuser\"}" -v

# Explore
curl -X POST "http://18.228.204.155/game/112783cd-66bd-4a85-9445-1ca293c41c25/explore" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"col\": 0, \"row\": 0, \"username\": \"myuser\"}" -v

# Resume game
curl -X POST "http://18.228.204.155/game/112783cd-66bd-4a85-9445-1ca293c41c25/resume" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"username\": \"myuser\"}" -v
