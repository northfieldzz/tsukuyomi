from os import environ
from flask import Flask
from discord import Intents
from tsukuyomi.v1.discord import Discord


def create_app():
    app = Flask(__name__)

    from tsukuyomi.v1 import v1_bp

    app.register_blueprint(v1_bp, url_prefix="/v1")

    @app.route('/version')
    def version():
        return environ.get('VERSION')

    intents = Intents.default()
    intents.message_content = True
    client = Discord(intents=intents)
    client.run(environ.get('DISCORD_TOKEN'))

    return app