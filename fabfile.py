# coding:utf-8

import os.path
import time
from fabric.api import run, env, roles, cd, put

env.parallel = False
env.use_ssh_config = True
env.hosts = ["node3", "node4"]

def deploy():
    run("supervisorctl stop spiderservice")
    time.sleep(1)
    put("SpiderService", "/services/spider/")
    run("supervisorctl start spiderservice")
    time.sleep(1)

