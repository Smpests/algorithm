from functools import wraps
import time

def timer(func):
    """
    函数执行时间计时器
    """
    @wraps(func)
    def wrapper(*args, **kwargs):
        start_time = time.monotonic()
        result = func(*args, **kwargs)
        took = time.monotonic() - start_time
        if took >= 1:
            print(f"{func.__name__}() took {took:.5f}s.")
        else:
            print(f"{func.__name__}() took {took * 1000:.5f}ms.")
        return result
    return wrapper
