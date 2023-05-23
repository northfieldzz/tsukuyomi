from os import environ
from discord import Client, Intents


class Discord(Client):
    async def on_ready(self):
        print("Logged on as")

    async def on_message(self, message):
        if message.author ==  self.user:
            return

        if message.content == 'ping':
            await message.channel.send('pong')


intents = Intents.default()
intents.message_content = True
client = Discord(intents=intents)
client.run(environ.get('DISCORD_TOKEN'))