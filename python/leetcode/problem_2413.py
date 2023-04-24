# import numpy as np
from python.decorator import timer


@timer
def smallestEvenMultiple(n: int) -> int:
    # n + n的写法比2 * n快
    return n if n % 2 == 0 else n * 2
    # 比单行if-else慢几倍
    # if n % 2 == 0:
    #     return n
    # return n + n
    # return np.lcm(2, n)  # 仅提示可以用numpy
    # return numpy.lcm.reduce(n)  # 若n是一个ndarray
