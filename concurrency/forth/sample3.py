import threading
import time


def summ():
    a = 0
    for i in range(0, 20000000):
        a+= i





start = time.perf_counter()
t1 = threading.Thread(target=summ)
t2 = threading.Thread(target=summ)
t3 = threading.Thread(target=summ)
t1.start()
t2.start()
t3.start()


t1.join()
t2.join()
t3.join()
end = time.perf_counter()
print(f"end: {end - start}")