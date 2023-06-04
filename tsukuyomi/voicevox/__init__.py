from os import environ
from json import dumps
from requests import post


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
