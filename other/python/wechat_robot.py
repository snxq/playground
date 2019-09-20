"""
超简易企业微信机器人

精确到分钟级别的超简易定时提醒机器人，只需要填下参数就可以直接运行。
"""

#!/usr/bin/python3
import requests
import json
import time

# 填写 URL
URL = ""
# 触发的小时 24进制
HOURS = [10, 11, 14, 15, 16, 17, 18, 19]
# 触发的分钟 0-60
MINS = [5, 10, 15]
# 触发的星期* 0-7 星期一到星期天
WDAYS = [0, 1, 2]

DATA = {
    "msgtype": "text",
    "text": {
        # 填写发送的文本
        "content": "test"
    },
    # @的列表
    "mentioned_list": []
}

while True:
    time.sleep(50)
    # 此处依赖系统时区，设置+8
    t = time.localtime()
    if all([
        t.tm_hour in HOURS,
        t.tm_min in MINS,
        t.tm_wday in WDAYS
    ]):
        requests.post(
            url=URL,
            data=json.dumps(DATA),
            headers={"Content-Type": "application/json"}
        )