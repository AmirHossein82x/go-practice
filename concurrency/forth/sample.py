import requests
import time


results = []



def getUrl(id):
    response = requests.get("https://jsonplaceholder.typicode.com/todos/"+str(id))
    results.append(response.json())



if __name__ == "__main__":
    start = time.perf_counter()
    for i in range(1, 20):
        getUrl(i)

    print(results)
    end = time.perf_counter()
    print(end - start)
