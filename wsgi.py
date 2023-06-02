from os import environ
from logging import getLogger, INFO
from multiprocessing import Process
from waitress import serve
from tsukuyomi import create_api, create_discord_client

root_logger = getLogger()
root_logger.setLevel(INFO)

logger = getLogger('waitress')
logger.setLevel(INFO)

if __name__ == '__main__':
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
