from os import environ
from logging import config
from multiprocessing import Process
from waitress import serve
from yaml import safe_load
from tsukuyomi import create_api, create_discord_client
from config import config_data

if __name__ == '__main__':
    config.dictConfig(safe_load(open(config_data).read()))
    api = create_api()
    discord_client = create_discord_client()

    api_process = Process(target=serve, args=(api,), kwargs={
        'host': '0.0.0.0',
        'port': environ['PORT'],
        'threads': environ.get('THREAD', 4)
    })
    discord_process = Process(target=discord_client.run, args=(environ['DISCORD_TOKEN'],))
    api_process.start()
    discord_process.start()

    api_process.join()
    discord_process.terminate()
