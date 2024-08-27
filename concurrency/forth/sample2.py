import requests
import time
import threading


results = []



def getUrl(id):
    response = requests.get("https://jsonplaceholder.typicode.com/todos/"+str(id))
    results.append(response.json())


threads = []
if __name__ == "__main__":
    start = time.perf_counter()
    for i in range(1, 50):
        t = threading.Thread(target=getUrl, args=(i,))
        t.start()
        threads.append(t)

    for item in threads:
        item.join()

    print(results)
    end = time.perf_counter()
    print(end - start)
