from os import environ
from logging import getLogger
from flask import Flask, jsonify
from discord import Intents
from tsukuyomi.discord import Discord

logger = getLogger(__name__)


def create_api():
    app = Flask(__name__)

    from tsukuyomi.v1 import v1_bp

    app.register_blueprint(v1_bp, url_prefix="/api/v1")

    @app.route('/version')
    def version():
        logger.info('sample')
        return jsonify({'version': environ.get('VERSION')}), 200

    return app


def create_discord_client():
    intents = Intents.default()
    intents.message_content = True
    return Discord(intents=intents)
