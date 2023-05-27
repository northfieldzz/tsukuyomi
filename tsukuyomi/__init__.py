from os import environ
from logging import getLogger
from threading import Thread
from flask import Flask, jsonify
from discord import Intents
from tsukuyomi.discord import Discord

logger = getLogger(__name__)


class Tsukuyomi:

    def create_api(self):
        app = Flask(__name__)

        from tsukuyomi.v1 import v1_bp

        app.register_blueprint(v1_bp, url_prefix="/api/v1")

        @app.route('/version')
        def version():
            return jsonify({'version': environ.get('VERSION')}), 200

        return app

    def create_discord_client(self):
        intents = Intents.default()
        intents.message_content = True
        return Discord(intents=intents)

    def launch(self, env=None):
        api = self.create_api()
        if env == 'production':
            from waitress import serve
            thread = Thread(target=serve, args=(api,), kwargs={
                'host': '0.0.0.0',
                'port': environ['PORT'],
                'debug': False
            })
        else:
            thread = Thread(target=api.run, kwargs={
                'host': '0.0.0.0',
                'port': environ['PORT'],
                'debug': True
            })
        thread.start()
        client = self.create_discord_client()
        client.run(environ.get('DISCORD_TOKEN'))
