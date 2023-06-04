from os import environ
from logging import config
from multiprocessing import Process
from yaml import safe_load
from tsukuyomi import create_api, create_discord_client
from config import config_data

if __name__ == '__main__':
    config.dictConfig(safe_load(open(config_data).read()))
    api = create_api()
    discord_client = create_discord_client()

    api_process = Process(target=api.run, kwargs={
        'host': '0.0.0.0',
        'port': environ['PORT'],
        'debug': True
    })
    discord_process = Process(target=discord_client.run, args=(environ['DISCORD_TOKEN'],))
    api_process.start()
    discord_process.start()

    api_process.join()
    discord_process.terminate()
