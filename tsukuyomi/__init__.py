from os import environ
from logging import getLogger
from flask import Flask, jsonify
from discord import Intents
from tsukuyomi.v1.discord import Discord

logger = getLogger(__name__)


def create_app():
    app = Flask(__name__)

    from tsukuyomi.v1 import v1_bp

    app.register_blueprint(v1_bp, url_prefix="/v1")

    @app.route('/version')
    def version():
        return jsonify({'version': environ.get('VERSION')}), 200

    intents = Intents.default()
    intents.message_content = True
    client = Discord(intents=intents)
    client.run(environ.get('DISCORD_TOKEN'))

    return app
