from os import path
from random import choice
from tempfile import TemporaryDirectory
from time import sleep
from logging import getLogger
from discord import Client, ForumChannel, FFmpegPCMAudio
from tsukuyomi.voicevox import Voicevox

logger = getLogger(__name__)


class Discord(Client):
    async def on_ready(self):
        logger.info('Logged on')

    async def on_message(self, message):
        if message.author == self.user:
            #
            return

        if 'pong' in message.content:
            #
            next_game_topic = None
            for channel in list(self.get_all_channels()):
                if isinstance(channel, ForumChannel):
                    next_game_topic = channel
                    break
            thread = choice(next_game_topic.threads)
            await message.channel.send(thread.jump_url)
        logger.info(message.content)

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
            sleep(1)
