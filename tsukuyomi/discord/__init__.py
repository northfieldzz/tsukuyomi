from discord import Client, ForumChannel
from random import choice


class Discord(Client):
    async def on_ready(self):
        print("Logged on as")

    async def on_message(self, message):
        if message.author == self.user:
            return

        if 'pong' in message.content:
            next_game_topic = None
            for channel in list(self.get_all_channels()):
                if isinstance(channel, ForumChannel):
                    next_game_topic = channel
                    break

            # print(next_game_topic.threads)

            thread = choice(next_game_topic.threads)
            await message.channel.send(thread.jump_url)
        print(message.content)
        # if isinstance(channel, ForumChannel):SS
        #     print(channel)
        #     print(type(channel))
        # await message.channel.send('pong')




