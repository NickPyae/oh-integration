import json

def updateTimeSeries(x):
    rids = execute('ZRANGE', x['key'], '0', '-1')
    for rid in rids:
        reading = json.loads(execute('GET', rid))
        ts = 'ts:%s:%s' % (reading['device'], reading['name'])
        execute('TS.ADD', ts, reading['created'], reading['value'], 'LABELS', 'id', reading['id'])

gb = GearsBuilder()
gb.foreach(updateTimeSeries)
gb.register(prefix='event:readings:*')
