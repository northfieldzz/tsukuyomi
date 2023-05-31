from os import path, environ
from random import choice
from json import dumps
from tempfile import TemporaryDirectory
from time import sleep
from logging import getLogger
from discord import Client, ForumChannel, FFmpegPCMAudio
from requests import post

logger = getLogger(__name__)


class Discord(Client):
    async def on_ready(self):
        print('Logged on as')

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

    async def on_voice_state_update(self, member, before, after):

        if member.bot:
            # 参加したメンバーがbot
            return

        # Bot Memberインスタンスの取得
        bot = member.guild.get_member(self.user.id)
        with TemporaryDirectory() as temp_dir:
            if before.channel is None:
                # ユーザが参加
                if len([m for m in member.voice.channel.members if bot.id == m.id]) == 0:
                    # TsukuyomiちゃんがVoiceChannelにまだいない
                    await member.voice.channel.connect()
                    voice = Voicevox()
                    wave_file = voice.fetch_voice(f'{member.name}, こんにちは', path.join(temp_dir, 'voice.wav'))
                    channel = self.get_channel(after.channel.id)
                    channel.guild.voice_client.play(FFmpegPCMAudio(wave_file))

            elif after.channel is None:
                # ユーザが退出
                channel = self.get_channel(before.channel.id)
                if len([member for member in channel.members if not member.bot]) == 0:
                    # bot以外のメンバーが全員退出
                    await channel.guild.voice_client.disconnect()
            elif before.channel != after.channel:
                # ユーザが移動
                channel = self.get_channel(before.channel.id)
                await channel.guild.voice_client.move_to(member.voice.channel)
                voice = Voicevox()
                wave_file = voice.fetch_voice('おい、どこいくのだ', path.join(temp_dir, 'voice.wav'))
                channel = self.get_channel(after.channel.id)
                channel.guild.voice_client.play(FFmpegPCMAudio(wave_file))
            sleep(0.1)


class Voicevox:
    scheme = environ['VOICEVOX_SCHEME']
    host = environ['VOICEVOX_HOST']
    port = environ['VOICEVOX_PORT']

    @property
    def url(self):
        return f'{self.scheme}://{self.host}:{self.port}'

    def fetch_voice(self, text, export_file, speaker=1):
        return self.fetch_wave(self.fetch_query(text, speaker), export_file)

    def fetch_query(self, text, speaker=1):
        r = post(f'{self.url}/audio_query', params={'text': text, 'speaker': speaker})
        if r.status_code == 200:
            query_data = r.json()
            return query_data
        return None

    def fetch_wave(self, query, export_file):
        res = post(f'{self.url}/synthesis', params={'speaker': 1}, data=dumps(query))
        if res.status_code == 200:
            with open(export_file, 'wb') as fp:
                fp.write(res.content)
        return export_file
