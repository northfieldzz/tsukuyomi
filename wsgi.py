from os import environ
from tsukuyomi import Tsukuyomi

if __name__ == '__main__':
    tsukuyomi = Tsukuyomi()
    tsukuyomi.launch(env=environ['ENV'])
