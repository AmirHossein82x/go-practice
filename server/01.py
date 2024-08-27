import requests
import json

user = {"id": 2, "full_name": "alireza", 'age': 21}
user_json = json.dumps(user)

# response = requests.get("http://localhost:8080/get-users?id=2")

# response = requests.post("http://localhost:8080/get-users", json=user, headers={"x": "00"})
response = requests.get("http://localhost:8080/users")
print(response)
print(response.content)
print(response.headers)