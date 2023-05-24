from os import environ
from discord import Client


class Discord(Client):
    async def on_ready(self):
        print("Logged on as")

    async def on_message(self, message):
        if message.author ==  self.user:
            return

        if message.content == 'ping':
            await message.channel.send('pong')


